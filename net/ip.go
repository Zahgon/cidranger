/*
Package net provides utility functions for working with IPs (net.IP).
*/
package net

import (
	"fmt"
	"net"
)

// IPVersion is version of IP address.
type IPVersion string

// Helper constants.
const (
	IPv4Uint32Count = 1
	IPv6Uint32Count = 4

	BitsPerUint32 = 32
	BytePerUint32 = 4

	IPv4 IPVersion = "IPv4"
	IPv6 IPVersion = "IPv6"
)

// ErrInvalidBitPosition is returned when bits requested is not valid.
var ErrInvalidBitPosition = fmt.Errorf("bit position not valid")

// ErrVersionMismatch is returned upon mismatch in network input versions.
var ErrVersionMismatch = fmt.Errorf("Network input version mismatch")

// ErrNoGreatestCommonBit is an error returned when no greatest common bit
// exists for the cidr ranges.
var ErrNoGreatestCommonBit = fmt.Errorf("No greatest common bit")

// NetworkNumber represents an IP address using uint32 as internal storage.
// IPv4 usings 1 uint32, while IPv6 uses 4 uint32.
type NetworkNumber []uint32

// NewNetworkNumber returns a equivalent NetworkNumber to given IP address,
// return nil if ip is neither IPv4 nor IPv6.
func NewNetworkNumber(ip net.IP) NetworkNumber {
	_ = "STUB: not implemented"
	return *new(NetworkNumber)
}

// ToV4 returns ip address if ip is IPv4, returns nil otherwise.
func (n NetworkNumber) ToV4() NetworkNumber { _ = "STUB: not implemented"; return *new(NetworkNumber) }

// ToV6 returns ip address if ip is IPv6, returns nil otherwise.
func (n NetworkNumber) ToV6() NetworkNumber { _ = "STUB: not implemented"; return *new(NetworkNumber) }

// ToIP returns equivalent net.IP.
func (n NetworkNumber) ToIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// Equal is the equality test for 2 network numbers.
func (n NetworkNumber) Equal(n1 NetworkNumber) bool { _ = "STUB: not implemented"; return false }

// Next returns the next logical network number.
func (n NetworkNumber) Next() NetworkNumber { _ = "STUB: not implemented"; return *new(NetworkNumber) }

// Previous returns the previous logical network number.
func (n NetworkNumber) Previous() NetworkNumber {
	_ = "STUB: not implemented"
	return *new(NetworkNumber)
}

// Bit returns uint32 representing the bit value at given position, e.g.,
// "128.0.0.0" has bit value of 1 at position 31, and 0 for positions 30 to 0.
func (n NetworkNumber) Bit(position uint) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// Mod 31 to get array index.

// LeastCommonBitPosition returns the smallest position of the preceding common
// bits of the 2 network numbers, and returns an error ErrNoGreatestCommonBit
// if the two network number diverges from the first bit.
// e.g., if the network number diverges after the 1st bit, it returns 131 for
// IPv6 and 31 for IPv4 .
func (n NetworkNumber) LeastCommonBitPosition(n1 NetworkNumber) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Network represents a block of network numbers, also known as CIDR.
type Network struct {
	net.IPNet
	Number NetworkNumber
	Mask   NetworkNumberMask
}

// NewNetwork returns Network built using given net.IPNet.
func NewNetwork(ipNet net.IPNet) Network { _ = "STUB: not implemented"; return *new(Network) }

// Masked returns a new network conforming to new mask.
func (n Network) Masked(ones int) Network { _ = "STUB: not implemented"; return *new(Network) }

// Contains returns true if NetworkNumber is in range of Network, false
// otherwise.
func (n Network) Contains(nn NetworkNumber) bool { _ = "STUB: not implemented"; return false }

// Contains returns true if Network covers o, false otherwise
func (n Network) Covers(o Network) bool { _ = "STUB: not implemented"; return false }

// LeastCommonBitPosition returns the smallest position of the preceding common
// bits of the 2 networks, and returns an error ErrNoGreatestCommonBit
// if the two network number diverges from the first bit.
func (n Network) LeastCommonBitPosition(n1 Network) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Equal is the equality test for 2 networks.
func (n Network) Equal(n1 Network) bool { _ = "STUB: not implemented"; return false }

func (n Network) String() string { _ = "STUB: not implemented"; return "" }

// NetworkNumberMask is an IP address.
type NetworkNumberMask NetworkNumber

// Mask returns a new masked NetworkNumber from given NetworkNumber.
func (m NetworkNumberMask) Mask(n NetworkNumber) (NetworkNumber, error) {
	_ = "STUB: not implemented"
	return *new(NetworkNumber), nil
}

// NextIP returns the next sequential ip.
func NextIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// PreviousIP returns the previous sequential ip.
func PreviousIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
