package resource

import "time"

type VMSpec struct {
	CPU       CPUConfig       `db:"cpu" json:"cpu"`
	Memory    MemoryConfig    `db:"memory" json:"memory"`
	Disks     []DiskConfig    `db:"disks" json:"disks"`
	Networks  []NetworkConfig `db:"networks" json:"networks"`
	Boot      BootConfig      `db:"boot" json:"boot"`
	Image     string          `db:"image" json:"image"`
	Autostart bool            `db:"autostart" json:"autostart"`
	CloudInit *CloudInit      `db:"cloud_init" json:"cloud_init,omitempty"`
}

type VMStatus struct {
	Phase       VMPhase  `json:"phase"`
	Node        string   `json:"node"`
	IPs         []string `json:"ips"`
	Running     bool     `json:"running"`
	Uptime      int16    `json:"uptime"`
	LastStarted *time.Time
	Conditions  []Condition
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
	ID     string `json:"id"`
	Volume string `json:"volume"`
	// Size in gigabytes
	Size int32 `json:"size"`
	// Values: 'sata' | 'scsi' | 'virtio' | 'nvme' | 'ide'
	Interface DiskInterface `json:"interface"`
	Boot      bool          `json:"boot"`
}

// TODO: fill in the struct
type BootConfig struct {
}

// TODO: fill in the struct
type CloudInit struct {
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
