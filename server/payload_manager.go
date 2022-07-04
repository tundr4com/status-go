package server

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/proto"

	"github.com/status-im/status-go/account/generator"
	"github.com/status-im/status-go/eth-node/keystore"
	"github.com/status-im/status-go/images"
	"github.com/status-im/status-go/multiaccounts"
	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/protobuf"
)

// PayloadManager is the interface for PayloadManagers and wraps the basic functions for fulfilling payload management
type PayloadManager interface {
	Mount() error
	Receive(data []byte) error
	ToSend() []byte
	Received() []byte
}

// PairingPayloadManagerConfig represents the initialisation parameters required for a PairingPayloadManager
type PairingPayloadManagerConfig struct {
	DB                             *multiaccounts.Database
	KeystorePath, KeyUID, Password string
}

// PairingPayloadManager is responsible for the whole lifecycle of a PairingPayload
type PairingPayloadManager struct {
	pem *PayloadEncryptionManager
	ppm *PairingPayloadMarshaller
	ppr PayloadRepository
}

// NewPairingPayloadManager generates a new and initialised PairingPayloadManager
func NewPairingPayloadManager(pk *ecdsa.PrivateKey, config *PairingPayloadManagerConfig) (*PairingPayloadManager, error) {
	pem, err := NewPayloadEncryptionManager(pk)
	if err != nil {
		return nil, err
	}

	// A new SHARED PairingPayload
	p := new(PairingPayload)

	return &PairingPayloadManager{
		pem: pem,
		ppm: NewPairingPayloadMarshaller(p),
		ppr: NewPairingPayloadRepository(p, config),
	}, nil
}

// Mount loads and prepares the payload to be stored in the PairingPayloadManager's state ready for later access
func (ppm *PairingPayloadManager) Mount() error {
	err := ppm.ppr.LoadFromSource()
	if err != nil {
		return err
	}

	pb, err := ppm.ppm.MarshalToProtobuf()
	if err != nil {
		return err
	}

	return ppm.pem.Encrypt(pb)
}

// Receive takes a []byte representing raw data, parses and stores the data
func (ppm *PairingPayloadManager) Receive(data []byte) error {
	err := ppm.pem.Decrypt(data)
	if err != nil {
		return err
	}

	err = ppm.ppm.UnmarshalProtobuf(ppm.pem.Received())
	if err != nil {
		return err
	}

	return ppm.ppr.StoreToSource()
}

// ToSend returns the result of Mount
func (ppm *PairingPayloadManager) ToSend() []byte {
	return ppm.pem.ToSend()
}

// Received returns the decrypted input of Receive
func (ppm *PairingPayloadManager) Received() []byte {
	return ppm.pem.Received()
}

// EncryptionPayload represents the plain text and encrypted text of payload data
type EncryptionPayload struct {
	plain     []byte
	encrypted []byte
}

// PayloadEncryptionManager is responsible for encrypting and decrypting payload data
type PayloadEncryptionManager struct {
	aesKey   []byte
	toSend   *EncryptionPayload
	received *EncryptionPayload
}

func NewPayloadEncryptionManager(pk *ecdsa.PrivateKey) (*PayloadEncryptionManager, error) {
	ek, err := makeEncryptionKey(pk)
	if err != nil {
		return nil, err
	}

	return &PayloadEncryptionManager{ek, new(EncryptionPayload), new(EncryptionPayload)}, nil
}

func (pem *PayloadEncryptionManager) Encrypt(data []byte) error {
	ep, err := common.Encrypt(data, pem.aesKey, rand.Reader)
	if err != nil {
		return err
	}

	pem.toSend.plain = data
	pem.toSend.encrypted = ep
	return nil
}

func (pem *PayloadEncryptionManager) Decrypt(data []byte) error {
	pd, err := common.Decrypt(data, pem.aesKey)
	if err != nil {
		return err
	}

	pem.received.encrypted = data
	pem.received.plain = pd
	return nil
}

func (pem *PayloadEncryptionManager) ToSend() []byte {
	return pem.toSend.encrypted
}

func (pem *PayloadEncryptionManager) Received() []byte {
	return pem.received.plain
}

func (pem *PayloadEncryptionManager) ResetPayload() {
	pem.toSend = new(EncryptionPayload)
	pem.received = new(EncryptionPayload)
}

// PairingPayload represents the payload structure a PairingServer handles
type PairingPayload struct {
	keys         map[string][]byte
	multiaccount *multiaccounts.Account
	password     string
}

// PairingPayloadMarshaller is responsible for marshalling and unmarshalling PairingServer payload data
type PairingPayloadMarshaller struct {
	*PairingPayload
}

func NewPairingPayloadMarshaller(p *PairingPayload) *PairingPayloadMarshaller {
	return &PairingPayloadMarshaller{PairingPayload: p}
}

func (ppm *PairingPayloadMarshaller) MarshalToProtobuf() ([]byte, error) {
	return proto.Marshal(&protobuf.LocalPairingPayload{
		Keys:         ppm.accountKeysToProtobuf(),
		Multiaccount: ppm.multiaccountToProtobuf(),
		Password:     ppm.password,
	})
}

func (ppm *PairingPayloadMarshaller) accountKeysToProtobuf() []*protobuf.LocalPairingPayload_Key {
	var keys []*protobuf.LocalPairingPayload_Key
	for name, data := range ppm.keys {
		keys = append(keys, &protobuf.LocalPairingPayload_Key{Name: name, Data: data})
	}
	return keys
}

func (ppm *PairingPayloadMarshaller) multiaccountToProtobuf() *protobuf.MultiAccount {
	var colourHashes []*protobuf.MultiAccount_ColourHash
	for _, index := range ppm.multiaccount.ColorHash {
		var i []int64
		for _, is := range index {
			i = append(i, int64(is))
		}

		colourHashes = append(colourHashes, &protobuf.MultiAccount_ColourHash{Index: i})
	}

	var identityImages []*protobuf.MultiAccount_IdentityImage
	for _, ii := range ppm.multiaccount.Images {
		identityImages = append(identityImages, &protobuf.MultiAccount_IdentityImage{
			KeyUid:       ii.KeyUID,
			Name:         ii.Name,
			Payload:      ii.Payload,
			Width:        int64(ii.Width),
			Height:       int64(ii.Height),
			Filesize:     int64(ii.FileSize),
			ResizeTarget: int64(ii.ResizeTarget),
			Clock:        ii.Clock,
		})
	}

	return &protobuf.MultiAccount{
		Name:           ppm.multiaccount.Name,
		Timestamp:      ppm.multiaccount.Timestamp,
		Identicon:      ppm.multiaccount.Identicon,
		ColorHash:      colourHashes,
		ColorId:        ppm.multiaccount.ColorID,
		KeycardPairing: ppm.multiaccount.KeycardPairing,
		KeyUid:         ppm.multiaccount.KeyUID,
		Images:         identityImages,
	}
}

func (ppm *PairingPayloadMarshaller) UnmarshalProtobuf(data []byte) error {
	pb := new(protobuf.LocalPairingPayload)
	err := proto.Unmarshal(data, pb)
	if err != nil {
		return err
	}

	ppm.accountKeysFromProtobuf(pb.Keys)
	ppm.multiaccountFromProtobuf(pb.Multiaccount)
	ppm.password = pb.Password
	return nil
}

func (ppm *PairingPayloadMarshaller) accountKeysFromProtobuf(pbKeys []*protobuf.LocalPairingPayload_Key) {
	if ppm.keys == nil {
		ppm.keys = make(map[string][]byte)
	}

	for _, key := range pbKeys {
		ppm.keys[key.Name] = key.Data
	}
}

func (ppm *PairingPayloadMarshaller) multiaccountFromProtobuf(pbMultiAccount *protobuf.MultiAccount) {
	var colourHash [][]int
	for _, index := range pbMultiAccount.ColorHash {
		var i []int
		for _, is := range index.Index {
			i = append(i, int(is))
		}

		colourHash = append(colourHash, i)
	}

	var identityImages []images.IdentityImage
	for _, ii := range pbMultiAccount.Images {
		identityImages = append(identityImages, images.IdentityImage{
			KeyUID:       ii.KeyUid,
			Name:         ii.Name,
			Payload:      ii.Payload,
			Width:        int(ii.Width),
			Height:       int(ii.Height),
			FileSize:     int(ii.Filesize),
			ResizeTarget: int(ii.ResizeTarget),
			Clock:        ii.Clock,
		})
	}

	ppm.multiaccount = &multiaccounts.Account{
		Name:           pbMultiAccount.Name,
		Timestamp:      pbMultiAccount.Timestamp,
		Identicon:      pbMultiAccount.Identicon,
		ColorHash:      colourHash,
		ColorID:        pbMultiAccount.ColorId,
		KeycardPairing: pbMultiAccount.KeycardPairing,
		KeyUID:         pbMultiAccount.KeyUid,
		Images:         identityImages,
	}
}

type PayloadRepository interface {
	LoadFromSource() error
	StoreToSource() error
}

// PairingPayloadRepository is responsible for loading, parsing, validating and storing PairingServer payload data
type PairingPayloadRepository struct {
	*PairingPayload

	multiaccountsDB *multiaccounts.Database

	keystorePath, keyUID string
}

func NewPairingPayloadRepository(p *PairingPayload, config *PairingPayloadManagerConfig) *PairingPayloadRepository {
	ppr := &PairingPayloadRepository{
		PairingPayload: p,
	}

	if config == nil {
		return ppr
	}

	ppr.multiaccountsDB = config.DB
	ppr.keystorePath = config.KeystorePath
	ppr.keyUID = config.KeyUID
	ppr.password = config.Password
	return ppr
}

func (ppr *PairingPayloadRepository) LoadFromSource() error {
	err := ppr.loadKeys(ppr.keystorePath)
	if err != nil {
		return err
	}

	err = ppr.validateKeys(ppr.password)
	if err != nil {
		return err
	}

	ppr.multiaccount, err = ppr.multiaccountsDB.GetAccount(ppr.keyUID)
	if err != nil {
		return err
	}

	return nil
}

func (ppr *PairingPayloadRepository) loadKeys(keyStorePath string) error {
	ppr.keys = make(map[string][]byte)

	fileWalker := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() || filepath.Dir(path) != keyStorePath {
			return nil
		}

		rawKeyFile, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("invalid account key file: %v", err)
		}

		accountKey := new(keystore.EncryptedKeyJSONV3)
		if err := json.Unmarshal(rawKeyFile, &accountKey); err != nil {
			return fmt.Errorf("failed to read key file: %s", err)
		}

		if len(accountKey.Address) != 40 {
			return fmt.Errorf("account key address has invalid length '%s'", accountKey.Address)
		}

		ppr.keys[fileInfo.Name()] = rawKeyFile

		return nil
	}

	err := filepath.Walk(keyStorePath, fileWalker)
	if err != nil {
		return fmt.Errorf("cannot traverse key store folder: %v", err)
	}

	return nil
}

func (ppr *PairingPayloadRepository) StoreToSource() error {
	err := ppr.validateKeys(ppr.password)
	if err != nil {
		return err
	}

	err = ppr.storeKeys(ppr.keystorePath)
	if err != nil {
		return err
	}

	err = ppr.storeMultiAccount()
	if err != nil {
		return err
	}

	// TODO install PublicKey into settings, probably do this outside of StoreToSource
	return nil
}

func (ppr *PairingPayloadRepository) validateKeys(password string) error {
	for _, key := range ppr.keys {
		k, err := keystore.DecryptKey(key, password)
		if err != nil {
			return err
		}

		err = generator.ValidateKeystoreExtendedKey(k)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ppr *PairingPayloadRepository) storeKeys(keyStorePath string) error {
	for name, data := range ppr.keys {
		accountKey := new(keystore.EncryptedKeyJSONV3)
		if err := json.Unmarshal(data, &accountKey); err != nil {
			return fmt.Errorf("failed to read key file: %s", err)
		}

		if len(accountKey.Address) != 40 {
			return fmt.Errorf("account key address has invalid length '%s'", accountKey.Address)
		}

		err := ioutil.WriteFile(filepath.Join(keyStorePath, name), data, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ppr *PairingPayloadRepository) storeMultiAccount() error {
	return ppr.multiaccountsDB.SaveAccount(*ppr.multiaccount)
}
