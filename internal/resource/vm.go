package resource

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/google/uuid"
	"libvirt.org/go/libvirt"
)

type VMArch string

const (
	ArchX86   VMArch = "x86"
	ArchArm   VMArch = "arm"
	ArchPPC   VMArch = "ppc"
	ArchRISCV VMArch = "risc-v"
	Archs390x VMArch = "s390x"
	ArchSPARC VMArch = "sparc"
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
	Arch      VMArch               `db:"arch" json:"arch"`
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

var LibvirtConn *libvirt.Connect = nil

func InitLibvirt() (*libvirt.Connect, error) {
	conn, err := libvirt.NewConnect("qemu:///system")

	return conn, err
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

type DomainType string

const (
	TypeQemu DomainType = "qemu"
	TypeKVM  DomainType = "kvm"
	TypeHVF  DomainType = "hvf"
)

type LibvirtDomain struct {
	XMLName       xml.Name `xml:"domain"`
	Type          string   `xml:"type,attr"`
	Name          string   `xml:"name"`
	UUID          string   `xml:"uuid"`
	Memory        int64    `xml:"memory"`
	CurrentMemory int64    `xml:"currentMemory,omitempty"`
	VCPU          int64    `xml:"vcpu"`
	OS            OS       `xml:"os,omitempty"`
	Devices       Devices  `xml:"devices,omitempty"`
	Features      Features `xml:"features,omitempty"`
	Clock         VMClock  `xml:"clock,omitempty"`
}

type DiskController struct {
	Type  string `xml:"type,attr"`
	Index string `xml:"index,attr"`
	Model string `xml:"model,attr"`
}

type OS struct {
	Type OSType `xml:"type"`
	Boot Boot   `xml:"boot"`
}

type OSType struct {
	Value   string `xml:",chardata"`
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr,omitempty"`
}

type Boot struct {
	Dev string `xml:"dev,attr"`
}

type Features struct {
	ACPI *struct{} `xml:"acpi,omitempty"`
}

type Devices struct {
	Emulator   string           `xml:"emulator"`
	Disks      []Disk           `xml:"disk"`
	Interfaces []VMNetInterface `xml:"interface,omitempty"`
	Graphics   Graphics         `xml:"graphics"`
	Controller DiskController   `xml:"controller,omitempty"`
}

type Disk struct {
	Type     string     `xml:"type,attr"`
	Device   string     `xml:"device,attr"`
	Driver   DiskDriver `xml:"driver,omitempty"`
	Source   DiskSource `xml:"source,omitempty"`
	Target   DiskTarget `xml:"target,omitempty"`
	ReadOnly *struct{}  `xml:"readonly,omitempty"`
}

type DiskDriver struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr,omitempty"`
}

type DiskSource struct {
	File   string `xml:"file,attr,omitempty"`
	Pool   string `xml:"pool,attr,omitempty"`
	Volume string `xml:"volume,attr,omitempty"`
}

type DiskTarget struct {
	Dev string `xml:"dev,attr,omitempty"`
	Bus string `xml:"bus,attr,omitempty"`
}

type VMClock struct {
	Sync string `xml:"sync,attr"`
}

type VMNetInterface struct {
	Type   string           `xml:"type,attr"`
	Source InterfaceSource  `xml:"source"`
	MAC    MacAddr          `xml:"mac"`
	Model  VMInterfaceModel `xml:"model,omitempty"`
}
type VMInterfaceModel struct {
	Type string `xml:"type,attr"`
}

type InterfaceSource struct {
	Network string `xml:"network,attr"`
}

type MacAddr struct {
	Address string `xml:"address,attr"`
}

type Graphics struct {
	Type   string `xml:"type,attr"`
	Port   int    `xml:"port,attr"`
	Keymap string `xml:"keymap,attr,omitempty"`
}

type LibvirtDomainBuilder struct {
	domain *LibvirtDomain
	spec   *VM
}

func NewLibvirtDomainBuilder() *LibvirtDomainBuilder {
	return &LibvirtDomainBuilder{
		domain: &LibvirtDomain{
			XMLName: xml.Name{
				Space: "",
				Local: "domain",
			},
			UUID: uuid.NewString(),
		},
	}
}

type OverviewSetupConfig struct {
	VMID int32
	Name string
	Desc string
}

func (lb *LibvirtDomainBuilder) SetupOverview(cfg OverviewSetupConfig) *LibvirtDomainBuilder {
	lb.domain.UUID = uuid.NewString()
	lb.domain.Name = cfg.Name
	lb.domain.XMLName.Local = "domain"

	return nil
}

type StorageType string

const (
	StorageTypeZFS       StorageType = "zfs"
	StorageTypeRaw       StorageType = "raw"
	StorageTypeNoStorage StorageType = "none"
)

type StorageConfig struct {
	Type         StorageType
	Pool         string
	Size         int64
	Interface    DiskInterface
	InstallMedia string
}

func (lb *LibvirtDomainBuilder) SetupStorage(cfg StorageConfig) *LibvirtDomainBuilder {
	return nil
}

type NetworkSetupConfig struct {
	Type          string // LAN(standard switch) or None
	EmulationType string // vitio,e1000,rt8139
	MacAddr       string // blank for randomized address
}

func (lb *LibvirtDomainBuilder) SetupNetwork(cfg NetworkSetupConfig) *LibvirtDomainBuilder {
	return nil
}

type HardwareSetupConfig struct {
	Sockets        int32
	Cores          int32
	Threads        int32
	Memory         string // GB like 4GB or MB like 512MB
	PCIPassthrough string
}

func (lb *LibvirtDomainBuilder) SetupHardware(cfg HardwareSetupConfig) *LibvirtDomainBuilder {
	return nil
}

type AdvancedSetupConfig struct {
	VNCPort                      int32
	VNCPassword                  string
	VNCResolution                string
	ClockOffset                  string
	StartupShutdownOrder         int32
	SerialConsole                bool
	VNWait                       bool
	StartOnBoot                  bool
	TPM                          bool
	EnableCloudInit              bool
	CloudInit                    string
	IgnoreUnimplementedMSRAccess bool
	QEMUGuestAgent               bool
}

func (lb *LibvirtDomainBuilder) SetupAdvanced(cfg AdvancedSetupConfig) *LibvirtDomainBuilder {
	return nil
}

func (lb *LibvirtDomainBuilder) Build() LibvirtDomain { return *lb.domain }

func (vm *VM) Init(ctx context.Context, res VM) error { return nil }

func (vm *VM) Create(ctx context.Context, res VM) error { return nil }

func (vm *VM) Start(ctx context.Context, name string) error { return nil }

func (vm *VM) Stop(ctx context.Context, name string) error { return nil }

func (vm *VM) Delete(ctx context.Context, name string) error { return nil }

func (vm *VM) Snapshot(ctx context.Context, name string) error { return nil }

func (vm *VM) Inspect(ctx context.Context, name string) (VM, error) { return VM{}, nil }

func (vm *VM) Reconcile(ctx context.Context, name string) (VM, error) { return VM{}, nil }
