// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64

package tracekernel

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfInspKlEventT struct {
	Target [20]int8
	Tuple  struct {
		Saddr   struct{ V6addr [16]uint8 }
		Daddr   struct{ V6addr [16]uint8 }
		Sport   uint16
		Dport   uint16
		L3Proto uint16
		L4Proto uint8
		Pad     uint8
	}
	SkbMeta struct {
		Netns    uint32
		Mark     uint32
		Ifindex  uint32
		Len      uint32
		Mtu      uint32
		SkState  uint32
		Protocol uint16
		Pad      uint16
	}
	Pid       uint32
	Cpu       uint32
	Direction uint32
	_         [4]byte
	Latency   uint64
	Point1    uint64
	Point2    uint64
	Point3    uint64
	Point4    uint64
}

type bpfRxlatencyT struct {
	Rcv         uint64
	Rcvfinish   uint64
	Local       uint64
	Localfinish uint64
}

type bpfTxlatencyT struct {
	Queuexmit uint64
	Local     uint64
	Output    uint64
	Finish    uint64
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KlatencyIpFinishOutput2      *ebpf.ProgramSpec `ebpf:"klatency_ip_finish_output2"`
	KlatencyIpLocal              *ebpf.ProgramSpec `ebpf:"klatency_ip_local"`
	KlatencyIpLocalDeliver       *ebpf.ProgramSpec `ebpf:"klatency_ip_local_deliver"`
	KlatencyIpLocalDeliverFinish *ebpf.ProgramSpec `ebpf:"klatency_ip_local_deliver_finish"`
	KlatencyIpOutput             *ebpf.ProgramSpec `ebpf:"klatency_ip_output"`
	KlatencyIpQueueXmit          *ebpf.ProgramSpec `ebpf:"klatency_ip_queue_xmit"`
	KlatencyIpRcv                *ebpf.ProgramSpec `ebpf:"klatency_ip_rcv"`
	KlatencyIpRcvFinish          *ebpf.ProgramSpec `ebpf:"klatency_ip_rcv_finish"`
	ReportConsume                *ebpf.ProgramSpec `ebpf:"report_consume"`
	ReportKfree                  *ebpf.ProgramSpec `ebpf:"report_kfree"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	InspKernelrxEntry *ebpf.MapSpec `ebpf:"insp_kernelrx_entry"`
	InspKerneltxEntry *ebpf.MapSpec `ebpf:"insp_kerneltx_entry"`
	InspKlatencyEvent *ebpf.MapSpec `ebpf:"insp_klatency_event"`
	InspKlatencyStack *ebpf.MapSpec `ebpf:"insp_klatency_stack"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	InspKernelrxEntry *ebpf.Map `ebpf:"insp_kernelrx_entry"`
	InspKerneltxEntry *ebpf.Map `ebpf:"insp_kerneltx_entry"`
	InspKlatencyEvent *ebpf.Map `ebpf:"insp_klatency_event"`
	InspKlatencyStack *ebpf.Map `ebpf:"insp_klatency_stack"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.InspKernelrxEntry,
		m.InspKerneltxEntry,
		m.InspKlatencyEvent,
		m.InspKlatencyStack,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KlatencyIpFinishOutput2      *ebpf.Program `ebpf:"klatency_ip_finish_output2"`
	KlatencyIpLocal              *ebpf.Program `ebpf:"klatency_ip_local"`
	KlatencyIpLocalDeliver       *ebpf.Program `ebpf:"klatency_ip_local_deliver"`
	KlatencyIpLocalDeliverFinish *ebpf.Program `ebpf:"klatency_ip_local_deliver_finish"`
	KlatencyIpOutput             *ebpf.Program `ebpf:"klatency_ip_output"`
	KlatencyIpQueueXmit          *ebpf.Program `ebpf:"klatency_ip_queue_xmit"`
	KlatencyIpRcv                *ebpf.Program `ebpf:"klatency_ip_rcv"`
	KlatencyIpRcvFinish          *ebpf.Program `ebpf:"klatency_ip_rcv_finish"`
	ReportConsume                *ebpf.Program `ebpf:"report_consume"`
	ReportKfree                  *ebpf.Program `ebpf:"report_kfree"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KlatencyIpFinishOutput2,
		p.KlatencyIpLocal,
		p.KlatencyIpLocalDeliver,
		p.KlatencyIpLocalDeliverFinish,
		p.KlatencyIpOutput,
		p.KlatencyIpQueueXmit,
		p.KlatencyIpRcv,
		p.KlatencyIpRcvFinish,
		p.ReportConsume,
		p.ReportKfree,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfeb.o
var _BpfBytes []byte
