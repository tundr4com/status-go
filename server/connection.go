package server

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

type ConnectionParamVersion int
type Mode int

const (
	Version1 ConnectionParamVersion = iota + 1
)

const (
	Receiving Mode = iota + 1
	Sending
)

type ConnectionParams struct {
	version    ConnectionParamVersion
	netIP      net.IP
	port       int
	privateKey *ecdsa.PrivateKey
	notBefore  time.Time
	serverMode Mode
}

func NewConnectionParams(netIP net.IP, port int, privateKey *ecdsa.PrivateKey, notBefore time.Time, mode Mode) *ConnectionParams {
	cp := new(ConnectionParams)
	cp.version = Version1
	cp.netIP = netIP
	cp.port = port
	cp.privateKey = privateKey
	cp.notBefore = notBefore
	cp.serverMode = mode
	return cp
}

// ToString generates a string required for generating a secure connection to another Status device.
//
// The returned string will look like below:
//   - "2:4FHRnp:H6G:6jpbvo2ucrtrnpXXF4DQYuysh697isH9ppd2aT8uSRDh:eQUriVtGtkWhPJFeLZjF:2"
//
// Format bytes encoded into a base58 string, delimited by ":"
//   - version
//   - net.IP
//   - port
//   - ecdsa.PrivateKey D field
//   - asn1.Marshal time.Time
//   - server mode
func (cp *ConnectionParams) ToString() (string, error) {
	v := base58.Encode(new(big.Int).SetInt64(int64(cp.version)).Bytes())
	ip := base58.Encode(cp.netIP)
	p := base58.Encode(new(big.Int).SetInt64(int64(cp.port)).Bytes())
	k := base58.Encode(cp.privateKey.D.Bytes())
	tb, err := asn1.Marshal(cp.notBefore.UTC())
	if err != nil {
		return "", err
	}
	t := base58.Encode(tb)
	m := base58.Encode(new(big.Int).SetInt64(int64(cp.serverMode)).Bytes())

	return fmt.Sprintf("%s:%s:%s:%s:%s:%s", v, ip, p, k, t, m), nil
}

// FromString parses a connection params string required for to securely connect to another Status device.
// This function parses a connection string generated by ToString
func (cp *ConnectionParams) FromString(s string) error {
	requiredParams := 6

	sData := strings.Split(s, ":")
	if len(sData) != requiredParams {
		return fmt.Errorf("expected data '%s' to have length of '%d', received '%d'", s, requiredParams, len(sData))
	}

	cp.version = ConnectionParamVersion(new(big.Int).SetBytes(base58.Decode(sData[0])).Int64())
	cp.netIP = base58.Decode(sData[1])
	cp.port = int(new(big.Int).SetBytes(base58.Decode(sData[2])).Int64())
	cp.privateKey = ToECDSA(base58.Decode(sData[3]))

	t := time.Time{}
	_, err := asn1.Unmarshal(base58.Decode(sData[4]), &t)
	if err != nil {
		return err
	}
	cp.notBefore = t
	cp.serverMode = Mode(new(big.Int).SetBytes(base58.Decode(sData[5])).Int64())

	return cp.validate()
}

func (cp *ConnectionParams) validate() error {
	err := cp.validateVersion()
	if err != nil {
		return err
	}

	err = cp.validateNetIP()
	if err != nil {
		return err
	}

	err = cp.validatePort()
	if err != nil {
		return err
	}

	err = cp.validatePrivateKey()
	if err != nil {
		return err
	}

	err = cp.validateNotBefore()
	if err != nil {
		return err
	}

	return cp.validateServerMode()
}

func (cp *ConnectionParams) validateVersion() error {
	switch cp.version {
	case Version1:
		return nil
	default:
		return fmt.Errorf("unsupported version '%d'", cp.version)
	}
}

func (cp *ConnectionParams) validateNetIP() error {
	if ok := net.ParseIP(cp.netIP.String()); ok == nil {
		return fmt.Errorf("invalid net ip '%s'", cp.netIP)
	}
	return nil
}

func (cp *ConnectionParams) validatePort() error {
	if cp.port > 0 && cp.port < 0x10000 {
		return nil
	}

	return fmt.Errorf("port '%d' outside of bounds of 1 - 65535", cp.port)
}

func (cp *ConnectionParams) validatePrivateKey() error {
	switch {
	case cp.privateKey.D == nil, cp.privateKey.D.Cmp(big.NewInt(0)) == 0:
		return fmt.Errorf("private key D not set")
	case cp.privateKey.PublicKey.X == nil, cp.privateKey.PublicKey.X.Cmp(big.NewInt(0)) == 0:
		return fmt.Errorf("public key X not set")
	case cp.privateKey.PublicKey.Y == nil, cp.privateKey.PublicKey.Y.Cmp(big.NewInt(0)) == 0:
		return fmt.Errorf("public key Y not set")
	default:
		return nil
	}
}

func (cp *ConnectionParams) validateNotBefore() error {
	if cp.notBefore.IsZero() {
		return fmt.Errorf("notBefore time is zero")
	}
	return nil
}

func (cp *ConnectionParams) validateServerMode() error {
	switch cp.serverMode {
	case Receiving, Sending:
		return nil
	default:
		return fmt.Errorf("invalid server mode '%d'", cp.serverMode)
	}
}

// Generate returns a *url.URL and encoded pem.Block generated from ConnectionParams set fields
func (cp *ConnectionParams) Generate() (*url.URL, []byte, error) {
	err := cp.validate()
	if err != nil {
		return nil, nil, err
	}

	u := &url.URL{
		Scheme: "https",
		Host:   fmt.Sprintf("%s:%d", cp.netIP, cp.port),
	}

	_, pem, err := GenerateCertFromKey(cp.privateKey, cp.notBefore, cp.netIP.String())
	if err != nil {
		return nil, nil, err
	}

	return u, pem, nil
}
