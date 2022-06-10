package server

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestPairingServerSuite(t *testing.T) {
	suite.Run(t, new(PairingServerSuite))
}

type PairingServerSuite struct {
	suite.Suite
	TestPairingServerComponents
}

func (s *PairingServerSuite) SetupSuite() {
	s.SetupPairingServerComponents(s.T())
}

func (s *PairingServerSuite) TestPairingServer_StartPairing() {
	modes := []Mode{
		Receiving,
		Sending,
	}

	for _, m := range modes {
		s.PS.mode = m

		// Random payload
		data := make([]byte, 32)
		_, err := rand.Read(data)
		s.Require().NoError(err)

		if m == Sending {
			s.PS.MountPayload(data)
		}

		err = s.PS.StartPairing()
		s.Require().NoError(err)

		// Give time for the sever to be ready, hacky I know, I'll iron this out
		time.Sleep(10 * time.Millisecond)

		cp, err := s.PS.MakeConnectionParams()
		s.Require().NoError(err)

		qr, err := cp.ToString()
		s.Require().NoError(err)

		// Client reads QR code and parses the connection string
		ccp := new(ConnectionParams)
		err = ccp.FromString(qr)
		s.Require().NoError(err)

		c, err := NewPairingClient(ccp)
		s.Require().NoError(err)

		if m == Receiving {
			c.MountPayload(data)
		}

		err = c.PairAccount()
		s.Require().NoError(err)

		s.Require().Equal(s.PS.payload.ToSend(), c.payload.Received())
		s.Require().Equal(s.PS.payload.Received(), c.payload.ToSend())

		// Reset the server's PayloadManager
		s.PS.payload = new(PayloadManager)
	}
}
