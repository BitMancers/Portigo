infra("ProjectName", {
	network("local-net-1", {
		cidr = "10.1.0.0/24",
	}),
	-- storage("zfs-ssd", {
	-- 	driver = "zfs", -- zfs, ext4
	-- 	pool = "tank/vms", -- can ONLY be specified in zfs
	-- }),

	storage("ext-ssd", {
		driver = "ext4", -- zfs, ext4
		storage = "/dev/sda1",
		mount = "/mnt/sda1",
	}),

	vm("debian-db", {
		type = "image", -- specify type for
		file = "images/debian-14.qcow2",
		ram = "2gb",
		storage = "50gb",
		cpu = 2,
	}),

	docker("my-redis", {
		image = "tensorflow/tensorflow",
		cpu = 1,
		ram = "500mb",
		network = "local-net-1",
	}),
})
