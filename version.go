package cidranger

import (
	"net"

	rnet "github.com/yl2chen/cidranger/net"
)

type rangerFactory func(rnet.IPVersion) Ranger

type versionedRanger struct {
	ipV4Ranger Ranger
	ipV6Ranger Ranger
}

func newVersionedRanger(factory rangerFactory) Ranger {
	_ = "STUB: not implemented"
	return *new(Ranger)
}

func (v *versionedRanger) Insert(entry RangerEntry) error { _ = "STUB: not implemented"; return nil }

func (v *versionedRanger) Remove(network net.IPNet) (RangerEntry, error) {
	_ = "STUB: not implemented"
	return *new(RangerEntry), nil
}

func (v *versionedRanger) Contains(ip net.IP) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *versionedRanger) ContainingNetworks(ip net.IP) ([]RangerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *versionedRanger) CoveredNetworks(network net.IPNet) ([]RangerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Len returns number of networks in ranger.
func (v *versionedRanger) Len() int { _ = "STUB: not implemented"; return 0 }

func (v *versionedRanger) getRangerForIP(ip net.IP) (Ranger, error) {
	_ = "STUB: not implemented"
	return *new(Ranger), nil
}
