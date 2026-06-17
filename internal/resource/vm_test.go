package resource

import (
	"log"
	"testing"

	"libvirt.org/go/libvirt"
)

func TestCreateVMDomainAndSpecBuilder(t *testing.T) {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		t.Fatal("Error creating libvirt connection", err)
	}
	vmBuilder := NewLibvirtDomainBuilder(conn)
	vmXMLDomain, vmSpec, err := vmBuilder.
		SetupOverview(OverviewSetupConfig{
			VMID: 0,
			Name: "Alpine Test VM",
			Desc: "Alpine test vm",
		}).
		SetupNetwork(NetworkSetupConfig{
			Type:          VMNetworkDefault,
			Source:        "default",
			EmulationType: NetworkEmulationVirtIO,
			MacAddr:       RandomMacAddress(),
		}).
		SetupOS(OSSetupConfig{
			Arch:         "x86_64",
			InstallType:  OSInstallISO,
			FileLocation: "tmp/alpine.iso",
		}).
		SetupStorage(StorageConfig{
			StorageType:     StorageTypeZFS,
			DiskType:        DiskTypeFile,
			Pool:            "rpool",
			DiskDeviceType:  DiskDeviceDisk,
			DiskBusType:     DiskTargetBusVirtIO,
			Size:            32 << 30,
			Interface:       DiskInterfaceVirtIO,
			InstallLocation: "",
			ReadOnly:        false,
		}).
		Build()

	if err != nil {
		log.Fatal(err)
	}

	_ = vmXMLDomain
	_ = vmSpec
}
