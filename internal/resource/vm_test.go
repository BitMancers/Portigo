package resource

import (
	"log"
	"testing"
)

func TestCreateVMDomainAndSpecBuilder(t *testing.T) {
	vmBuilder := NewLibvirtDomainBuilder()
	vmXMLDomain, vmSpec, err := vmBuilder.SetupOverview(OverviewSetupConfig{
		VMID: 0,
		Name: "",
		Desc: "",
	}).
		SetupNetwork(NetworkSetupConfig{
			Type:          "lan",
			EmulationType: "virtio",
			MacAddr:       "", // blank for randomized
		}).
		SetupStorage(StorageConfig{
			Type:         StorageTypeRaw,
			Pool:         "root",
			Size:         20 << 30, // 20 MB
			Interface:    DiskInterfaceVirtIO,
			InstallMedia: "tmp/file.iso",
		}).
		SetupHardware(HardwareSetupConfig{
			Sockets:        1,
			Cores:          1,
			Threads:        1,
			Memory:         "2GB",
			PCIPassthrough: "",
		}).
		SetupAdvanced(AdvancedSetupConfig{
			VNCPort:                      -1, // -1 means auto allocate
			VNCPassword:                  "randompassword",
			VNCResolution:                "800x480",
			ClockOffset:                  "",
			StartupShutdownOrder:         0,
			SerialConsole:                true,
			VNWait:                       false,
			StartOnBoot:                  false,
			TPM:                          false,
			EnableCloudInit:              false,
			CloudInit:                    "",
			IgnoreUnimplementedMSRAccess: false,
			QEMUGuestAgent:               true,
		}).
		Build()

	if err != nil {
		log.Fatal(err)
	}

	_ = vmXMLDomain
	_ = vmSpec
}
