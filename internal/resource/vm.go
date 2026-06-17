package resource

import (
	"encoding/xml"
	"time"

	"github.com/BitMancers/Portigo/internal/utils"
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

type VMNetworkType string

const (
	VMNetworkDefault VMNetworkType = "network"
	VMNetworkBridge  VMNetworkType = "bridge"
	VMNeworkNone     VMNetworkType = ""
)

type VMNetworkInterface struct {
	Name       string        `db:"name" json:"name"`
	Type       VMNetworkType `db:"type" json:"type"`
	NetworkRef string        `db:"network_ref" json:"network_ref"` // reference to a Network Resource
	MACAddress string        `db:"mac" json:"mac"`
	IPAddress  string        `db:"ip_addr" json:"ip_addr,omitempty"`
}

type VMSpec struct {
	Arch      VMArch               `db:"arch" json:"arch"`
	CPU       CPUConfig            `db:"cpu" json:"cpu"`
	Memory    MemoryConfig         `db:"memory" json:"memory"`
	Disks     []DiskConfig         `db:"disks" json:"disks"`
	Networks  []VMNetworkInterface `db:"networks" json:"networks"`
	Image     string               `db:"image" json:"image"`
	Autostart bool                 `db:"autostart" json:"autostart"`
	CloudInit *CloudInit           `db:"cloud_init" json:"cloud_init,omitempty"`
}

type VMStatus struct {
	Phase       VMPhase     `db:"phase" json:"phase"`
	Infra       string      `db:"node" json:"node"`
	IPs         []string    `db:"ips" json:"ips"`
	Running     bool        `db:"running" json:"running"`
	Uptime      int16       `db:"uptime" json:"uptime"`
	LastStarted time.Time   `db:"last_started" json:"last_started"`
	Conditions  []Condition `db:"conditions" json:"conditions"`
}

type CPUConfig struct {
	Cores   int32 `json:"cores"`
	Sockets int32 `json:"sockets"`
}
type MemoryConfig struct {
	Dedicated int64 `json:"dedicated"`
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
	ID           string        `json:"id"`
	VolumeRef    string        `json:"volume-ref"`
	Pool         string        `json:"pool"`
	Image        string        `json:"image"`
	InstallMedia string        `json:"media"`
	Size         int64         `json:"size"`      // Size in bytes
	Interface    DiskInterface `json:"interface"` // 'sata' | 'scsi' | 'virtio' | 'nvme' | 'ide'
	Boot         bool          `json:"boot"`
}

// TODO: fill in the struct
type BootConfig struct {
	Dev string
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

// Libvirt XML specific struct
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
	VCPU          int32    `xml:"vcpu"`
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

type DiskDeviceType string

const (
	DiskDeviceDisk   DiskDeviceType = "disk"
	DiskDeviceCDROM  DiskDeviceType = "cdrom"
	DiskDeviceFloppy DiskDeviceType = "floppy"
	DiskDeviceLun    DiskDeviceType = "lun"
)

type DiskType string

const (
	DiskTypeFile      DiskType = "file"
	DiskTypeBlock     DiskType = "block"
	DiskTypeDir       DiskType = "dir"
	DiskTypeNetwork   DiskType = "network"
	DiskTypeVolume    DiskType = "volume"
	DiskTypeNVMe      DiskType = "nvme"
	DiskTypeVhostUser DiskType = "vhostuser"
)

type Disk struct {
	Type     DiskType       `xml:"type,attr"`
	Device   DiskDeviceType `xml:"device,attr"` // disk, cdrom, floppy, lun
	Driver   DiskDriver     `xml:"driver,omitempty"`
	Source   DiskSource     `xml:"source,omitempty"`
	Target   DiskTarget     `xml:"target,omitempty"`
	ReadOnly *struct{}      `xml:"readonly,omitempty"`
}

type DiskDriverType string

const (
	DiskDriverQCow2 DiskDriverType = "qcow2"
	DiskDriverRaw   DiskDriverType = "raw"
)

type DiskDriver struct {
	Name string         `xml:"name,attr"`           // "qemu" in most cases
	Type DiskDriverType `xml:"type,attr,omitempty"` // qcow2 and raw
}

type DiskSource struct {
	File   string `xml:"file,attr,omitempty"`
	Pool   string `xml:"pool,attr,omitempty"`
	Volume string `xml:"volume,attr,omitempty"`
}

type DiskBusType string

const (
	DiskTargetBusIDE    DiskBusType = "ide"
	DiskTargetBusSATA   DiskBusType = "sata"
	DiskTargetBusVirtIO DiskBusType = "virtio"
	DiskTargetBusSCSI   DiskBusType = "scsi"
)

type DiskTarget struct {
	Dev string      `xml:"dev,attr,omitempty"`
	Bus DiskBusType `xml:"bus,attr,omitempty"`
}

type VMClock struct {
	Sync string `xml:"sync,attr"`
}

type VMNetInterface struct {
	Type   VMNetworkType    `xml:"type,attr"`
	Source InterfaceSource  `xml:"source"`
	MAC    MacAddr          `xml:"mac"`
	Model  VMInterfaceModel `xml:"model,omitempty"`
}

type NetworkEmulationType string

const (
	NetworkEmulationVirtIO  NetworkEmulationType = "vitio" // default
	NetworkEmulationE1000   NetworkEmulationType = "e1000"
	NetworkEmulationE1000e  NetworkEmulationType = "e1000e"
	NetworkEmulationRTL8139 NetworkEmulationType = "rtl8139"
)

type VMInterfaceModel struct {
	Type NetworkEmulationType `xml:"type,attr"`
}

type InterfaceSource struct {
	Network string `xml:"network,attr"`
}

type MacAddr struct {
	Address string `xml:"address,attr"`
}

type Graphics struct {
	Type     string `xml:"type,attr"` // most probably 'vnc'
	Port     int32  `xml:"port,attr"`
	Keymap   string `xml:"keymap,attr,omitempty"`
	Autoport string `xml:"autoport,attr,omitempty"`
	Listen   string `xml:"listen,attr,omitempty"`
	Passwd   string `xml:"passwd,attr,omitempty"`
}

type LibvirtDomainBuilder struct {
	domain         *LibvirtDomain
	spec           *VM
	conn           *libvirt.Connect
	isoInstalledVM bool
}

func NewLibvirtDomainBuilder(conn *libvirt.Connect) *LibvirtDomainBuilder {
	vmUUID := uuid.NewString()
	return &LibvirtDomainBuilder{
		isoInstalledVM: false,
		domain: &LibvirtDomain{
			XMLName: xml.Name{
				Space: "",
				Local: "domain",
			},
			UUID: vmUUID,
		},
		spec: &VM{
			Kind: "vm",
			Status: VMStatus{
				Phase:       VMPhaseCreate,
				Infra:       "-",
				IPs:         []string{},
				Running:     false,
				Uptime:      0,
				LastStarted: time.Time{},
				Conditions:  []Condition{},
			},
		},
		conn: conn,
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
	lb.spec.ID = cfg.VMID
	lb.spec.Metadata.Name = cfg.Name + string(cfg.VMID)
	lb.spec.Kind = ResourceKindVM

	return lb
}

type StorageType string

const (
	StorageTypeZFS       StorageType = "zfs"
	StorageTypeRaw       StorageType = "raw"
	StorageTypeNoStorage StorageType = "none"
)

type StorageConfig struct {
	StorageType     StorageType // default zfs
	DiskType        DiskType    // default file, qcow2, raw
	Pool            string      // only if zfs
	DiskDeviceType  DiskDeviceType
	DiskBusType     DiskBusType
	Size            int64
	Interface       DiskInterface
	InstallLocation string
	ReadOnly        bool
}

func (lb *LibvirtDomainBuilder) SetupStorage(cfg StorageConfig) *LibvirtDomainBuilder {

	storageSetup := Disk{
		Type: cfg.DiskType,
		Driver: DiskDriver{
			Name: "qemu",
		},
		Source: DiskSource{
			File: cfg.InstallLocation,
		},
		Target: DiskTarget{
			Bus: cfg.DiskBusType,
		},
	}

	specDiskConfig := DiskConfig{
		ID:           uuid.NewString(),
		InstallMedia: cfg.InstallLocation,
		Size:         cfg.Size,
		Interface:    cfg.Interface,
		Boot:         true,
	}
	if utils.MatchExtension(cfg.InstallLocation, "qcow2") {
		storageSetup.Driver.Type = DiskDriverQCow2
	} else if utils.MatchExtension(cfg.InstallLocation, "raw") {
		storageSetup.Driver.Type = DiskDriverRaw
	} else if utils.MatchExtension(cfg.InstallLocation, "iso") {
		lb.spec.Spec.Disks = append(lb.spec.Spec.Disks, specDiskConfig)
		storageSetup.Target.Dev = "hdc"
	}
	if cfg.ReadOnly {
		storageSetup.ReadOnly = &struct{}{}
	}

	if cfg.StorageType == StorageTypeZFS && cfg.Pool != "" {
		storageSetup.Source.Pool = cfg.Pool
	}

	var currentDevIndex int = 0
	for idx, disk := range lb.domain.Devices.Disks {
		if disk.Target.Bus == cfg.DiskBusType {
			currentDevIndex = idx
		}
	}

	nextDev := nextDev(cfg.DiskBusType, currentDevIndex)

	storageSetup.Target.Dev = nextDev

	lb.domain.Devices.Disks = append(lb.domain.Devices.Disks, storageSetup)

	return lb
}

type OSInstallType string

const (
	OSInstallISO   OSInstallType = "iso"
	OSInstallRaw   OSInstallType = "raw"
	OSInstallQCow2 OSInstallType = "qcow2"
)

type OSSetupConfig struct {
	Arch         string
	InstallType  OSInstallType
	FileLocation string
}

func (lb *LibvirtDomainBuilder) SetupOS(cfg OSSetupConfig) *LibvirtDomainBuilder {
	lb.domain.OS = OS{
		Type: OSType{
			Value:   "hvm",
			Arch:    cfg.Arch,
			Machine: "pc", // TODO: may need to changed
		},
	}

	if cfg.InstallType == OSInstallISO {
		lb.isoInstalledVM = true
		lb.SetupStorage(StorageConfig{
			StorageType:     StorageTypeZFS,
			DiskType:        DiskTypeFile,
			Pool:            "rpool",
			DiskDeviceType:  DiskDeviceCDROM,
			Interface:       DiskInterfaceSata,
			InstallLocation: cfg.FileLocation,
			ReadOnly:        true,
		})
	}

	return lb
}

type NetworkSetupConfig struct {
	Type          VMNetworkType        // LAN(standard switch), Bridge or None
	Source        string               // network names like default or libvirt created network
	EmulationType NetworkEmulationType // vitio ,e1000, e1000e, rtl8139
	MacAddr       string               // blank for randomized address
}

func (lb *LibvirtDomainBuilder) SetupNetwork(cfg NetworkSetupConfig) *LibvirtDomainBuilder {
	vmNetworkInterface := VMNetInterface{
		Type: cfg.Type,
		Source: InterfaceSource{
			Network: cfg.Source,
		},
		MAC: MacAddr{
			Address: cfg.MacAddr,
		},
		Model: VMInterfaceModel{
			Type: cfg.EmulationType,
		},
	}

	networkSpec := VMNetworkInterface{
		NetworkRef: cfg.Source,
	}

	lb.domain.Devices.Interfaces = append(lb.domain.Devices.Interfaces, vmNetworkInterface)
	lb.spec.Spec.Networks = append(lb.spec.Spec.Networks, networkSpec)

	return lb
}

type HardwareSetupConfig struct {
	Sockets        int32
	Cores          int32
	Threads        int32
	Memory         int64 // in bytes
	PCIPassthrough string
}

func (lb *LibvirtDomainBuilder) SetupHardware(cfg HardwareSetupConfig) *LibvirtDomainBuilder {
	lb.domain.VCPU = cfg.Cores
	lb.domain.CurrentMemory = cfg.Memory
	lb.domain.Memory = cfg.Memory

	lb.spec.Spec.CPU = CPUConfig{
		Cores:   cfg.Cores,
		Sockets: cfg.Sockets,
	}
	lb.spec.Spec.Memory = MemoryConfig{
		Dedicated: cfg.Memory,
	}
	return lb
}

type AdvancedSetupConfig struct {
	VNCPort              int32
	VNCPassword          string
	VNCResolution        string
	ClockOffset          string
	StartupShutdownOrder int32
	SerialConsole        bool
	VNWait               bool
	StartOnBoot          bool
	TPM                  bool
	EnableCloudInit      bool
	CloudInit            string
	QEMUGuestAgent       bool
}

// TODO: add cloud init stuff
func (lb *LibvirtDomainBuilder) SetupAdvanced(cfg AdvancedSetupConfig) *LibvirtDomainBuilder {
	graphicsSetup := Graphics{
		Type:   "vnc",
		Port:   cfg.VNCPort,
		Keymap: "en",
		Listen: "0.0.0.0",
		Passwd: cfg.VNCPassword,
	}

	if cfg.VNCPort == -1 {
		graphicsSetup.Autoport = "yes"
	}

	lb.spec.Spec.Autostart = cfg.StartOnBoot
	lb.domain.Devices.Graphics = graphicsSetup

	return lb
}

func (lb *LibvirtDomainBuilder) Build() (LibvirtDomain, VM, error) {

	if lb.isoInstalledVM {
		isoFile := lb.domain.Devices.Disks[0].Source.File
		lb.SetupStorage(StorageConfig{
			StorageType:     StorageTypeZFS,
			DiskType:        DiskTypeFile,
			Pool:            lb.spec.Spec.Disks[0].Pool,
			DiskDeviceType:  DiskDeviceDisk,
			DiskBusType:     lb.domain.Devices.Disks[0].Target.Bus,
			Size:            lb.spec.Spec.Disks[0].Size,
			Interface:       DiskInterfaceSata,
			InstallLocation: "",
			ReadOnly:        false,
		})
	}
	return *lb.domain, *lb.spec, nil
}
