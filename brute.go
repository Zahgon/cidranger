package cidranger

import (
	"net"
)

// bruteRanger is a brute force implementation of Ranger.  Insertion and
// deletion of networks is performed on an internal storage in the form of
// map[string]net.IPNet (constant time operations).  However, inclusion tests are
// always performed linearly at no guaranteed traversal order of recorded networks,
// so one can assume a worst case performance of O(N).  The performance can be
// boosted many ways, e.g. changing usage of net.IPNet.Contains() to using masked
// bits equality checking, but the main purpose of this implementation is for
// testing because the correctness of this implementation can be easily guaranteed,
// and used as the ground truth when running a wider range of 'random' tests on
// other more sophisticated implementations.
type bruteRanger struct {
	ipV4Entries map[string]RangerEntry
	ipV6Entries map[string]RangerEntry
}

// newBruteRanger returns a new Ranger.
func newBruteRanger() Ranger { _ = "STUB: not implemented"; return *new(Ranger) }

// Insert inserts a RangerEntry into ranger.
func (b *bruteRanger) Insert(entry RangerEntry) error { _ = "STUB: not implemented"; return nil }

// Remove removes a RangerEntry identified by given network from ranger.
func (b *bruteRanger) Remove(network net.IPNet) (RangerEntry, error) {
	_ = "STUB: not implemented"
	return *new(RangerEntry), nil
}

// Contains returns bool indicating whether given ip is contained by any
// network in ranger.
func (b *bruteRanger) Contains(ip net.IP) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ContainingNetworks returns all RangerEntry(s) that given ip contained in.
func (b *bruteRanger) ContainingNetworks(ip net.IP) ([]RangerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CoveredNetworks returns the list of RangerEntry(s) the given ipnet
// covers.  That is, the networks that are completely subsumed by the
// specified network.
func (b *bruteRanger) CoveredNetworks(network net.IPNet) ([]RangerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Len returns number of networks in ranger.
func (b *bruteRanger) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *bruteRanger) getEntriesByVersion(ip net.IP) (map[string]RangerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
