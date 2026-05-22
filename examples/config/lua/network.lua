network(...)

--[[
  ```text
Network
 ├── Type
 ├── Address Space
 ├── Connectivity Rules
 ├── DHCP/IPAM
 ├── DNS
 ├── Routing
 ├── Encryption
 ├── Attachments
 └── Cross-node Behavior
  ```
  ]]

network("<network-name>", {
	driver = "<driver-type>", -- driver-types: bridge|nat|isolated|vlan|vxlan|macvlan|wireguard|host
	cidr = "x.x.x.x/x",

	dhcp = true or false,
	outbound = true or false,
})

-- giving it to a vm or a container should be like for an example network

network("backend-net", {
	driver = "bridge",
	cidr = "10.10.0.0/24",
})

network("frontend-net", {
	driver = "bridge",
	cidr = "192.168.0.0/24",
})

vm("database-vm", {
	network = "backend",
})

container("api", {
	network = "backend",
})

-- multiple network per resource
container("graphql", {
	networks = {
		"frontend",
		"backend",
	},
})

-- or something like

container("graphql", {
	networks = {
		{ network = "frontend", ip = "192.168.0.1" },
		{ network = "backend", ip = "10.10.0.1" },
	},
})

-- Api types
--[[
  | Type      | Purpose             |
  | --------- | ------------------- |
  | bridge    | local host bridge   |
  | nat       | outbound internet   |
  | isolated  | internal only       |
  | vlan      | physical VLAN       |
  | vxlan     | cross-node overlay  |
  | macvlan   | direct LAN exposure |
  | wireguard | encrypted mesh      |
  | host      | host networking     |
]]

-- Example: NAT

network("public-net", {
  driver = "nat",
  cidr = "172.16.0.0/24"
  dhcp = true ,
  outbound = true
})


-- Example: Isolated

-- no internet access, no uplink
network("internal-db-net", {
  driver = "isolated",
  cidr = "10.100.0.0/24"
})

-- Example: VLAN-backend network
network("servers-net", {
  driver = "vlan",
  interface = "eno1",
  vlan = 100, 
  cidr = "192.168.100.0/24"
})

-- Example: Cross-Node Overlay
-- containers on different nodes can communicate directly
network('cluster-net', {
  driver = "vxlan",
  cidr = "10.200.0.0/16",
  encryption = true
})

-- Explicit network attachment
attach("my-vm", "backend-network", {
  ip = "10.10.0.5"
})

-- IPAM
-- Automatic ip allocation
network("backend-net", {
  cidr = "10.10.0.0/24",
  ipam = { mode = "auto" },
})

-- Static IPs

vm("db", {
    networks = {
        {
            network = "backend-net",
            ip = "10.10.0.10"
        }
    }
})
