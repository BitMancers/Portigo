package resource

type ContainerSpec struct {
	Runtime       ContainerRuntime  `db:"runtime" json:"runtime"`
	Image         string            `db:"image" json:"image"`
	Command       []string          `db:"command" json:"command"`
	Env           map[string]string `db:"env,omitempty" json:"env,omitempty"`
	Ports         PortMapping       `db:"ports,omitempty" json:"ports,omitempty"`
	Volumes       VolumeMount       `db:"volumes,omitempty" json:"volumes,omitempty"`
	Limits        ResourceLimits    `db:"limits" json:"limits"`
	RestartPolicy RestartPolicy     `db:"restart_policy" json:"restart_policy"`
}

type ContainerRuntime string

const (
	RuntimeDocker ContainerRuntime = "docker"
	RuntimePodman ContainerRuntime = "podman"
)

type RestartPolicy string

const (
	RestartAlways RestartPolicy = "always"
	RestartNever  RestartPolicy = "never"
	// Restart when the container exits with non-zero code
	RestartOnFailure RestartPolicy = "failure"
)

type ResourceLimits struct {
	// storage capacity in
	Storage string `db:"storage,omitempty" json:"storage,omitempty"`
	// memory limits
	Memory string `db:"memory" json:"memory"`
	// cpu cores
	CPU int32 `db:"cpu" json:"cpu"`
}

type PortMapping struct {
	HostPort      int `db:"host_port" json:"host_port"`
	ContainerPort int `db:"container_port" json:"container_port"`
	// TCP | UDP
	Protocol string `db:"protocol" json:"protocol"`
}

type VolumeMount struct {
	Volume   string `db:"volume" json:"volume"`
	Target   string `db:"target" json:"target"`
	ReadOnly bool   `db:"read_only" json:"read_only"`
}
