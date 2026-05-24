package resource

import (
	"context"
	"fmt"
	"time"

	"libvirt.org/go/libvirt"
)

type VMNetworkInterface struct {
	Name       string      `db:"name" json:"name"`
	Type       NetworkType `db:"type" json:"type"`
	NetworkRef string      `db:"network_ref" json:"network_ref"` // reference to a Network Resource
	MACAddress string      `db:"mac" json:"mac"`
	Model      string      `db:"model" json:"model,omitempty"` // vitio,e1000,rt8139
	IPAddress  string      `db:"ip_addr" json:"ip_addr,omitempty"`
}

type VMSpec struct {
	CPU       CPUConfig            `db:"cpu" json:"cpu"`
	Memory    MemoryConfig         `db:"memory" json:"memory"`
	Disks     []DiskConfig         `db:"disks" json:"disks"`
	Networks  []VMNetworkInterface `db:"networks" json:"networks"`
	Boot      BootConfig           `db:"boot" json:"boot"`
	Image     string               `db:"image" json:"image"`
	Autostart bool                 `db:"autostart" json:"autostart"`
	CloudInit *CloudInit           `db:"cloud_init" json:"cloud_init,omitempty"`
}

type VMStatus struct {
	Phase       VMPhase     `db:"phase" json:"phase"`
	Node        string      `db:"node" json:"node"`
	IPs         []string    `db:"ips" json:"ips"`
	Running     bool        `db:"running" json:"running"`
	Uptime      int16       `db:"uptime" json:"uptime"`
	LastStarted time.Time   `db:"last_started" json:"last_started"`
	Conditions  []Condition `db:"conditions" json:"conditions"`
}

type CPUConfig struct {
	Cores   int `json:"cores"`
	Sockets int `json:"sockets"`
}
type MemoryConfig struct {
	DedicatedMB int  `json:"dedicated_mb"`
	Ballooning  bool `json:"ballooning"`
}

type DiskInterface string

const (
	DiskInterfaceSata   DiskInterface = "sata"
	DiskInterfaceSCSI   DiskInterface = "scsi"
	DiskInterfaceVirtIO DiskInterface = "virtio"
	DiskInterfaceNVME   DiskInterface = "nvme"
	DiskInterfaceIDE    DiskInterface = "ide"
)

type DiskConfig struct {
	ID        string        `json:"id"`
	VolumeRef string        `json:"volume-ref"`
	Image     string        `json:"image"`
	Size      int64         `json:"size"`      // Size in bytes
	Interface DiskInterface `json:"interface"` // 'sata' | 'scsi' | 'virtio' | 'nvme' | 'ide'
	Boot      bool          `json:"boot"`
}

// TODO: fill in the struct
type BootConfig struct {
}

// TODO: fill in the struct
type CloudInit struct {
}

func InitLibvirt() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer conn.Close()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	fmt.Printf("%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println("ERROR: ", err)
		dom.Free()
	}
}

type VMPhase string

const (
	VMPhaseCreate   VMPhase = "create"
	VMPhaseRunning  VMPhase = "running"
	VMPhaseShutdown VMPhase = "shutdown"
	VMPhaseDelete   VMPhase = "delete"
	VMPhaseStart    VMPhase = "started"
)

type VM Resource[VMSpec, VMStatus]

func (vm *VM) Create(ctx context.Context, res VM) error {
	return nil
}
func (vm *VM) Start(ctx context.Context, name string) error {
	return nil
}
func (vm *VM) Stop(ctx context.Context, name string) error            { return nil }
func (vm *VM) Delete(ctx context.Context, name string) error          { return nil }
func (vm *VM) Snapshot(ctx context.Context, name string) error        { return nil }
func (vm *VM) Inspect(ctx context.Context, name string) (VM, error)   { return VM{}, nil }
func (vm *VM) Reconcile(ctx context.Context, name string) (VM, error) { return VM{}, nil }
