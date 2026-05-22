package resource

import "time"

type VolumeSpec struct {
	Driver string `json:"driver"`
	// size in bytes
	Size     int    `json:"size"`
	Path     string `json:"path"`
	Capacity int    `json:"capacity"`
	Used     int    `json:"used"`

	FileSystem string      `json:"filesystem"` // ext4 | xfs | btrfs | zfs
	MountPoint string      `json:"mountpoint,omitempty"`
	ReadOnly   bool        `json:"readonly"`
	Conditions []Condition `json:"conditions,omitempty"`
	LastError  string      `json:"last_error"`
	CreatedAt  time.Time   `json:"created_at"`
}

type VolumeStatus string

const (
	VolumePending  VolumeStatus = "pending"
	VolumeCreating VolumeStatus = "creating"
	VolumeReady    VolumeStatus = "ready"
	VolumeAttached VolumeStatus = "attached"
	VolumeFailed   VolumeStatus = "failed"
)

type Volume Resource[VolumeSpec, VolumeStatus]
