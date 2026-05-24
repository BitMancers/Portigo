package resource

import (
	"portigo/internal/utils"
	"testing"
	"time"
)

func TestCreatingVM(t *testing.T) {
	vmID := utils.RandStringRunes(10)
	diskID := utils.RandStringRunes(10)
	vm := VM{
		ID:       vmID,
		Kind:     ResourceKindVM,
		Metadata: Metadata{},
		Spec: VMSpec{
			CPU:    CPUConfig{Cores: 1, Sockets: 1},
			Memory: MemoryConfig{DedicatedMB: 1024, Ballooning: true},
			Disks: []DiskConfig{
				DiskConfig{
					ID:        diskID,
					VolumeRef: "example-vol",
					Image:     "vm-image.qcow2",
					Size:      (5 << 30), // 5 MB
					Interface: "virtio",
					Boot:      true,
				},
			},
			Networks: []VMNetworkInterface{
				VMNetworkInterface{
					Name:       vmID + "name",
					Type:       NetworkHost,
					NetworkRef: "<none>",
					MACAddress: "E5:AE:05:7C:D9:FE",
					Model:      "virtio",
					IPAddress:  "192.169.1.88",
				},
			},
			Boot:      BootConfig{},
			Image:     "",
			Autostart: false,
			CloudInit: &CloudInit{},
		},
		Status: VMStatus{
			Phase:       "",
			Node:        "",
			IPs:         []string{},
			Running:     false,
			Uptime:      0,
			LastStarted: time.Now(),
			Conditions: []Condition{
				Condition{
					Type:               "",
					Status:             "",
					LastTransitionTime: time.Time{},
					Reason:             "",
					Message:            "",
				},
			},
		},
	}
	_ = vm
}
