package tc

import (
	"fmt"

	"github.com/mdlayher/netlink"
)

// Qdisc represents the queueing discipline part of traffic control
type Qdisc struct {
	Tc
}

const (
	rtmNewQdisc = 36
	rtmDelQdisc = 37
	rtmGetQdisc = 38
)

// Qdisc allows to read and alter queues
func (tc *Tc) Qdisc() *Qdisc {
	return &Qdisc{*tc}
}

// Add creates a new queueing discipline
func (qd *Qdisc) Add(info *Object) error {
	options, err := validateQdiscObject(rtmNewQdisc, info)
	if err != nil {
		return err
	}
	return qd.action(rtmNewQdisc, netlink.Create|netlink.Excl, info, options)
}

// Replace add/remove a queueing discipline. If the node does not exist yet it is created
func (qd *Qdisc) Replace(info *Object) error {
	options, err := validateQdiscObject(rtmNewQdisc, info)
	if err != nil {
		return err
	}
	return qd.action(rtmNewQdisc, netlink.Create|netlink.Replace, info, options)
}

// Link performs a replace on an existing queueing discipline
func (qd *Qdisc) Link(info *Object) error {
	options, err := validateQdiscObject(rtmNewQdisc, info)
	if err != nil {
		return err
	}
	return qd.action(rtmNewQdisc, netlink.Replace, info, options)
}

// Delte removes a queueing discipline
func (qd *Qdisc) Delte(info *Object) error {
	options, err := validateQdiscObject(rtmDelQdisc, info)
	if err != nil {
		return err
	}
	return qd.action(rtmDelQdisc, netlink.HeaderFlags(0), info, options)
}

// Change modifies a queueing discipline 'in place'
func (qd *Qdisc) Change(info *Object) error {
	options, err := validateQdiscObject(rtmNewQdisc, info)
	if err != nil {
		return err
	}
	return qd.action(rtmNewQdisc, netlink.HeaderFlags(0), info, options)
}

// Get fetches all queueing disciplines
func (qd *Qdisc) Get() ([]Object, error) {
	return qd.get(rtmGetQdisc, &Msg{})
}

func validateQdiscObject(action int, info *Object) ([]tcOption, error) {
	options := []tcOption{}
	if info.Ifindex == 0 {
		return options, fmt.Errorf("Could not set device ID 0")
	}

	// TODO: improve logic and check combinations

	switch info.Kind {
	case "clsact":
		options = append(options, tcOption{Interpretation: vtString, Type: tcaKind, Data: info.Kind})
	default:
		return options, ErrNotImplemented
	}

	if info.Stats != nil || info.XStats != nil || info.Stats2 != nil || info.FqCodel != nil || info.BPF != nil {
		return options, ErrNotImplemented
	}

	if info.EgressBlock != 0 {
		options = append(options, tcOption{Interpretation: vtUint32, Type: tcaEgressBlock, Data: info.EgressBlock})
	}
	if info.IngressBlock != 0 {
		options = append(options, tcOption{Interpretation: vtUint32, Type: tcaIngressBlock, Data: info.IngressBlock})
	}
	if info.HwOffload != 0 {
		options = append(options, tcOption{Interpretation: vtUint8, Type: tcaHwOffload, Data: info.HwOffload})
	}
	if info.Chain != 0 {
		options = append(options, tcOption{Interpretation: vtUint32, Type: tcaChain, Data: info.Chain})
	}
	return options, nil
}