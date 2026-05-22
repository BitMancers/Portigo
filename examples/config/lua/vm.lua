infra("<infrastructure or project-name>", function()
	network("<network-name>", {
		cidr = "10.1.0.0/24",
	})

	storage("<storage-name>", {
		driver = "zfs", -- zfs, ext4
		pool = "tank/vms", -- can ONLY be specified in zfs
	})

	vm("<vm-name>", {
		type = "image", -- specify type for
		file = "images/debian-14.qcow2",
		ram = "2gb",
		storage = "50gb",
		cpu = 2,
	})
end)
