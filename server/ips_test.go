package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/status-im/status-go/server/servertest"
)

func TestIPsTestingSuite(t *testing.T) {
	suite.Run(t, new(IPsTestingSuite))
}

type IPsTestingSuite struct {
	suite.Suite
	servertest.TestLoggerComponents
}

func (s *IPsTestingSuite) SetupSuite() {
	s.SetupLoggerComponents()
}

func (s *IPsTestingSuite) TestConnectionParams_GetLocalAddressesForPairingServer() {

	allIps := [][]net.IP{
		{
			net.IPv4(127, 0, 0, 1),
			net.IPv6loopback,
		},
		{
			net.IPv4(192, 168, 1, 42),
			net.IP{0xfc, 0x80, 0, 0, 0, 0, 0, 0, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			net.IPv4(11, 12, 13, 14),
		},
		{
			net.IP{0xfc, 0x80, 0, 0, 0, 0, 0, 0, 0xff, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
		{
			net.IP{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0xff, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
	}

	// First NI is a loop-back
	ni0 := allIps[0]
	s.Require().NotNil(ni0[0].To4())
	s.Require().True(ni0[0].IsLoopback())
	s.Require().Len(ni0[1], net.IPv6len)
	s.Require().True(ni0[1].IsLoopback())

	// Second NI's both IP addresses fits the needs. IPv6 should be filtered out.
	ni1 := allIps[1]
	s.Require().NotNil(ni1[0].To4())
	s.Require().True(ni1[0].IsGlobalUnicast())
	s.Require().True(ni1[0].IsPrivate())
	s.Require().Len(ni1[1], net.IPv6len)
	s.Require().True(ni1[1].IsGlobalUnicast())
	s.Require().True(ni1[1].IsPrivate())

	// Next NI should be filtered out as non-private
	ni2 := allIps[2]
	s.Require().NotNil(ni2[0].To4())
	s.Require().False(ni2[0].IsPrivate())

	// Next NI fits the needs, should be taken,
	// as no preferred IPv4 is available on this NI.
	ni3 := allIps[3]
	s.Require().Len(ni3[0], net.IPv6len)
	s.Require().True(ni3[0].IsGlobalUnicast())
	s.Require().True(ni3[0].IsPrivate())

	// Last NI has a link-local unicast address,
	// which should be filtered out as non-private.
	ni4 := allIps[4]
	s.Require().Len(ni4[0], net.IPv6len)
	s.Require().True(ni4[0].IsLinkLocalUnicast())
	s.Require().False(ni4[0].IsGlobalUnicast())
	s.Require().False(ni4[0].IsPrivate())

	ips := filterAddressesForPairingServer(allIps)
	s.Require().Len(ips, 2)

	var ip1, ip2 net.IP

	if ips[0].To4() != nil {
		ip1 = ips[0]
		ip2 = ips[1]
	} else {
		ip1 = ips[1]
		ip2 = ips[0]
	}

	s.Require().NotNil(ip1.To4())
	s.Require().NotNil(ni1[0].To4())
	s.Require().Equal(ip1.To4(), ni1[0].To4())
	s.Require().Equal(ip2, ni3[0])
}

func (s *IPsTestingSuite) TestConnectionParams_FindReachableAddresses() {
	var remoteIps []net.IP
	var localNets []net.IPNet
	var ips []net.IP

	// Test 1
	remoteIps = []net.IP{
		net.IPv4(10, 1, 2, 3),
		net.IPv4(172, 16, 2, 42),
		net.IPv4(192, 168, 1, 42),
	}
	localNets = []net.IPNet{
		{IP: net.IPv4(192, 168, 1, 43), Mask: net.IPv4Mask(255, 255, 255, 0)},
	}
	ips = findReachableAddresses(remoteIps, localNets)
	s.Require().Len(ips, 1)
	s.Require().Equal(ips[0], remoteIps[2])

	// Test 2
	remoteIps = []net.IP{
		net.IPv4(10, 1, 2, 3),
		net.IPv4(172, 16, 2, 42),
		net.IPv4(192, 168, 1, 42),
	}
	localNets = []net.IPNet{
		{IP: net.IPv4(10, 1, 1, 1), Mask: net.IPv4Mask(255, 255, 0, 0)},
		{IP: net.IPv4(172, 16, 2, 43), Mask: net.IPv4Mask(255, 255, 255, 0)},
		{IP: net.IPv4(192, 168, 2, 43), Mask: net.IPv4Mask(255, 255, 255, 0)},
	}
	ips = findReachableAddresses(remoteIps, localNets)
	s.Require().Len(ips, 2)
	s.Require().Equal(ips[0], remoteIps[0])
	s.Require().Equal(ips[1], remoteIps[1])

	// Test 3
	remoteIps = []net.IP{
		net.IPv4(10, 1, 2, 3),
		net.IPv4(172, 16, 2, 42),
		net.IPv4(192, 168, 1, 42),
	}
	localNets = []net.IPNet{}
	ips = findReachableAddresses(remoteIps, localNets)
	s.Require().Len(ips, 0)

	// Test 4
	remoteIps = []net.IP{}
	localNets = []net.IPNet{}
	ips = findReachableAddresses(remoteIps, localNets)
	s.Require().Len(ips, 0)
}

func (s *IPsTestingSuite) TestConnectionParams_RealNetworksTest() {

	//  This test is intended to be run manually.
	//  1. set `printDetails` to true
	//  2. run Part 1 on 2 devices
	//  3. copy printed results to Part 2
	//  4. update expected results in Part 3
	//  5. run Part 3

	// printing is disabled by default to avoid showing sensitive information
	const printDetails = false

	printLocalAddresses := func(in [][]net.IP) {
		fmt.Println("{")
		for _, a := range in {
			fmt.Println("    {")
			for _, v := range a {
				fmt.Println("        net.ParseIP(\"", v.String(), "\"),")

			}
			fmt.Println("    },")
		}
		fmt.Println("}")
	}

	printNets := func(in []net.IPNet) {
		fmt.Println("{")
		for _, n := range in {
			fmt.Println("    \"", n.String(), "\",")
		}
		fmt.Println("}")
	}

	parseNets := func(in []string) []net.IPNet {
		var out []net.IPNet
		for _, v := range in {
			_, network, err := net.ParseCIDR(v)
			s.Require().NoError(err)
			out = append(out, *network)
		}
		return out
	}

	//	Part 1:
	//	print needed stuff. Run on both machines.

	addrs, err := getLocalAddresses()
	s.Require().NoError(err)
	s.Logger.Info("MacOS:", zap.Any("addrs", addrs))

	if printDetails {
		printLocalAddresses(addrs)
	}

	nets, err := getAllAvailableNetworks()
	s.Require().NoError(err)
	s.Logger.Info("MacOS:", zap.Any("nets", nets))

	if printDetails {
		printNets(nets)
	}

	//	Part 2:
	//	Input all printed devices details below

	macNIs := [][]net.IP{
		{
			net.ParseIP("127.0.0.1"),
			net.ParseIP("::1"),
			net.ParseIP("fe80::1"),
		},
		{
			net.ParseIP("fe80::c1f:ee0d:1476:dd9a"),
			net.ParseIP("192.168.1.36"),
		},
		{
			net.ParseIP("172.16.9.1"),
		},
	}

	macNets := parseNets([]string{
		"127.0.0.1/8",
		"::1/128",
		"fe80::1/64",
		"fe80::c1f:ee0d:1476:dd9a/64",
		"192.168.1.36/24",
		"172.16.9.1/24",
	})

	winNIs := [][]net.IP{
		{
			net.ParseIP("fe80::6fd7:5ce4:554f:165a"),
			net.ParseIP("192.168.1.33"),
		},
		{
			net.ParseIP("fe80::ffa5:98e1:285c:42eb"),
			net.ParseIP("10.0.85.2"),
		},
		{
			net.ParseIP("::1"),
			net.ParseIP("127.0.0.1"),
		},
	}

	winNets := parseNets([]string{
		"fe80::6fd7:5ce4:554f:165a/64",
		"192.168.1.33/24",
		"fe80::ffa5:98e1:285c:42eb/64",
		"10.0.85.2/32",
		"::1/128",
		"127.0.0.1/8",
	})

	//	Part 3:
	//	The test itself

	// Windows as server, Mac as client
	winIPs := filterAddressesForPairingServer(winNIs)
	winReachableIps := findReachableAddresses(winIPs, macNets)
	s.Require().Len(winReachableIps, 1)
	s.Require().Equal(winReachableIps[0].String(), "192.168.1.33")

	// Windows as server, Mac as client
	macIPs := filterAddressesForPairingServer(macNIs)
	macReachableIps := findReachableAddresses(macIPs, winNets)
	s.Require().Len(macReachableIps, 1)
	s.Require().Equal(macReachableIps[0].String(), "192.168.1.36")
}
