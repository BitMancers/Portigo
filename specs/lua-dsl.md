# Lua DSL spec for Portigo

## Resources list

- vm
- container
- lxc
- network
- volume
- snapshot
- secret
- node

## Declare resources like

```lua
-- only the syntax
<resource_name>(name, config)

```

### Example
```lua

network("db", {
    driver = "bridge",
    cidr = "10.1.0.0/24"
})

vm("debian-postgres", {
    type = "image", -- specify type for the vm, iso or image
    file = "images/debian-14.qcow2",
    auto_convert_to_raw = true,
    ram = "2gb",
    storage = "50gb",
    cpu = 2,
    network = "db"
})

storage("<storage-name>", {
    driver = "zfs", -- zfs, ext4
    pool = "tank/vms", -- can ONLY be specified in zfs
    create_if_doesnt_exists = true,
    size = "50GB"
})

container("redis-docker", {
    network = "docker-net",
})

```


