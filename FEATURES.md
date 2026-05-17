> [!IMPORTANT]
> This is for self discussion and talking to myself

## What are we trying to do here?

We are trying to build a software infrastructure management solution kind of like proxmox and coolify but a bit more declarative and UI friendly.
*Coolify* is kind of already that and we are trying to build smth similar but with GO.
So it'll be mostly portable without installer scripts and all.

Platform support should be limited to Linux but not limited to a distro.
Features:
- Managing virtual machines and containers.
    - Docker containers(with podman backend support)
    - LXC containers
    - QEMU Vitual Machines
- Managing terminal/shell, resource allocation and restriction, network management on each of the containers/VM
- Managing networking across these machines and containers
- Managing storage and other portigo nodes
- ZFS or other system support
- Lua API which will allow users to script infrastructure as code and start multiple VMs and networks in it(kind of like kubernetes?)
- Encrypted KEYS management
- User / RBAC System
- Infrastructure snapshots
- Engine state reconciliation

Mostly we need a GUI for doing most of the stuff we do.
and also a scripting language like lua or smth which will help us declare the infrastructure.

