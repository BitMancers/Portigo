package resource

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestCreatingVMSpecXMLParsing(t *testing.T) {
	// QEMU emulated guest on x86_64
	actualDomainValue := LibvirtDomain{
		XMLName: xml.Name{
			Space: "",
			Local: "domain",
		},
		Type:          "qemu",
		Name:          "QEMU-fedora-i686",
		UUID:          "c7a5fdbd-cdaf-9455-926a-d65c16db1809",
		Memory:        219200,
		CurrentMemory: 219200,
		VCPU:          2,
		OS: OS{
			Type: OSType{
				Value:   "hvm",
				Arch:    "i686",
				Machine: "pc",
			},
			Boot: Boot{
				Dev: "cdrom",
			},
		},
		Devices: Devices{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disks: []Disk{
				{
					Type:   "file",
					Device: "cdrom",
					Source: DiskSource{
						File: "tmp/fedora.iso",
					},
					Target: DiskTarget{
						Dev: "hdc",
					},
					ReadOnly: &struct{}{},
				},
				{
					Type:   "file",
					Device: "disk",
					Source: DiskSource{
						File: "tmp/fedora.img",
					},
					Target: DiskTarget{
						Dev: "hda",
					},
				},
			},
			Interfaces: []VMNetInterface{
				{
					Type: "network",
					Source: InterfaceSource{
						Network: "default",
					},
				},
			},
			Graphics: Graphics{
				Type: "vnc",
				Port: -1,
			},
		},
	}

	actualXMLValue, err := xml.Marshal(actualDomainValue)
	if err != nil {
		t.Fatal("ERROR: marshalling xml data: ", err)
	}

	fmt.Println(string(actualXMLValue))

	// QEMU emulated guest on x86_64, taken from https://libvirt.org/drvqemu.html#id26
	expectedXMLValue :=
		`<domain type='qemu'>
  <name>QEMU-fedora-i686</name>
  <uuid>c7a5fdbd-cdaf-9455-926a-d65c16db1809</uuid>
  <memory>219200</memory>
  <currentMemory>219200</currentMemory>
  <vcpu>2</vcpu>
  <os>
    <type arch='i686' machine='pc'>hvm</type>
    <boot dev='cdrom'/>
  </os>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='cdrom'>
      <source file='tmp/fedora.iso'/>
      <target dev='hdc'/>
      <readonly/>
    </disk>
    <disk type='file' device='disk'>
      <source file='tmp/fedora.img'/>
      <target dev='hda'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
    </interface>
    <graphics type='vnc' port='-1'/>
  </devices>
</domain>`

	var expectedDomain LibvirtDomain
	err = xml.Unmarshal([]byte(expectedXMLValue), &expectedDomain)
	if err != nil {
		t.Fatal("Unable to unmarshal value: ", err)
	}

	if !reflect.DeepEqual(expectedDomain, actualDomainValue) {
		t.Fatalf("XML values does not match on either side\nLeft: %+#v\nRight: %+#v", expectedDomain, actualDomainValue)
	}

	// KVM hardware accelerated guest on  i686
	actualDomainValue = LibvirtDomain{
		XMLName: xml.Name{
			Space: "",
			Local: "domain",
		},
		Type:   "kvm",
		Name:   "demo2",
		UUID:   "4dea24b3-1d52-d8f3-2516-782e98a23fa0",
		Memory: 131072,
		VCPU:   1,
		OS: OS{
			Type: OSType{
				Value: "hvm",
				Arch:  "i686",
			},
		},
		Clock: VMClock{
			Sync: "localtime",
		},
		Devices: Devices{
			Emulator: "/usr/bin/qemu-kvm",
			Disks: []Disk{
				{
					Type:   "file",
					Device: "disk",
					Source: DiskSource{
						File: "tmp/demo2.img",
					},
					Target: DiskTarget{
						Dev: "hda",
					},
				},
			},
			Interfaces: []VMNetInterface{
				{
					Type: "network",
					Source: InterfaceSource{
						Network: "default",
					},
					MAC: MacAddr{
						Address: "24:42:53:21:52:45",
					},
				},
			},
			Graphics: Graphics{
				Type:   "vnc",
				Port:   -1,
				Keymap: "de",
			},
		},
	}

	expectedXMLValue = `
<domain type='kvm'>
  <name>demo2</name>
  <uuid>4dea24b3-1d52-d8f3-2516-782e98a23fa0</uuid>
  <memory>131072</memory>
  <vcpu>1</vcpu>
  <os>
    <type arch="i686">hvm</type>
  </os>
  <clock sync="localtime"/>
  <devices>
    <emulator>/usr/bin/qemu-kvm</emulator>
    <disk type='file' device='disk'>
      <source file='tmp/demo2.img'/>
      <target dev='hda'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
      <mac address='24:42:53:21:52:45'/>
    </interface>
    <graphics type='vnc' port='-1' keymap='de'/>
  </devices>
</domain>
`

	var expectedDomain2 LibvirtDomain
	err = xml.Unmarshal([]byte(expectedXMLValue), &expectedDomain2)
	if err != nil {
		t.Fatal("ERROR: unmarshalling xml data", err)
	}

	if !reflect.DeepEqual(expectedDomain2, actualDomainValue) {
		t.Fatalf("\nLeft: %#v \nRight: %#v\n", expectedDomain2, actualDomainValue)
	}

	expectedXMLValue =
		`<domain type='hvf'>
  <name>hvf-demo</name>
  <uuid>4dea24b3-1d52-d8f3-2516-782e98a23fa0</uuid>
  <memory>131072</memory>
  <vcpu>1</vcpu>
  <os>
    <type arch="x86_64">hvm</type>
  </os>
  <features>
    <acpi/>
  </features>
  <clock sync="localtime"/>
  <devices>
    <emulator>/usr/local/bin/qemu-system-x86_64</emulator>
    <controller type='scsi' index='0' model='virtio-scsi'/>
    <disk type='volume' device='disk'>
      <driver name='qemu' type='qcow2'/>
      <source pool='default' volume='myos'/>
      <target bus='scsi' dev='sda'/>
    </disk>
    <interface type='user'>
      <mac address='24:42:53:21:52:45'/>
      <model type='virtio'/>
    </interface>
    <graphics type='vnc' port='-1'/>
  </devices>
</domain>`

	var expectedDomain3 LibvirtDomain
	actualDomainValue = LibvirtDomain{
		XMLName: xml.Name{
			Space: "",
			Local: "domain",
		},
		Type:   "hvf",
		Name:   "hvf-demo",
		UUID:   "4dea24b3-1d52-d8f3-2516-782e98a23fa0",
		Memory: 131072,
		VCPU:   1,
		OS: OS{
			Type: OSType{
				Value: "hvm",
				Arch:  "x86_64",
			},
			Boot: Boot{},
		},
		Features: Features{
			ACPI: &struct{}{},
		},
		Clock: VMClock{
			Sync: "localtime",
		},
		Devices: Devices{
			Emulator: "/usr/local/bin/qemu-system-x86_64",
			Controller: DiskController{
				Type:  "scsi",
				Index: "0",
				Model: "virtio-scsi",
			},
			Disks: []Disk{
				{
					Type:   "volume",
					Device: "disk",
					Source: DiskSource{
						Pool:   "default",
						Volume: "myos",
					},
					Target: DiskTarget{
						Dev: "sda",
						Bus: "scsi",
					},
					Driver: DiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
				},
			},
			Interfaces: []VMNetInterface{
				{
					Type:   "user",
					Source: InterfaceSource{},
					MAC: MacAddr{
						Address: "24:42:53:21:52:45",
					},
					Model: VMInterfaceModel{
						Type: "virtio",
					},
				},
			},
			Graphics: Graphics{
				Type: "vnc",
				Port: -1,
			},
		},
	}

	err = xml.Unmarshal([]byte(expectedXMLValue), &expectedDomain3)
	if err != nil {
		t.Fatal("ERROR: unmarshalling xml data", err)
	}

	if !reflect.DeepEqual(expectedDomain3, actualDomainValue) {
		t.Fatalf("\nLeft: %#v \nRight: %#v\n", expectedDomain3, actualDomainValue)
	}
}
