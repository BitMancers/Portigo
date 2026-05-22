package resource

type LXCSPec struct {
	Template   string         `db:"template" json:"template"`
	Privileged bool           `db:"privileged" json:"privileged"`
	Limits     ResourceLimits `db:"limits" json:"limits"`
	RootFS     RootFSConfig   `db:"rootfs" json:"rootfs"`
	Features   LXCFeatures    `db:"features" json:"features"`
	Mounts     []BindMount    `db:"mounts" json:"mounts"`
}

type LXCFeatures struct {
	Nesting bool `db:"nesting" json:"nesting"`
	Fuse    bool `db:"fuse" json:"fuse"`
	Keyctl  bool `db:"keyctl" json:"keyctl"`
}

type RootFSConfig struct {
	Path     string `db:"path" json:"path"`
	ReadOnly string `db:"readonly" json:"readonly"`
	// zfs | btrfs | lvm | overlay, etc
	StorageType string `db:"storage_type" json:"storage_type"`
}

type BindMount struct {
	Source      string `db:"source" json:"source"`
	Destination string `db:"destination" json:"destination"`
	ReadOnly    bool   `db:"readonly" json:"readonly"`
}
