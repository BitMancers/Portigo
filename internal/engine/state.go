package engine

import (
	"github.com/pocketbase/pocketbase"
	"libvirt.org/go/libvirt"
)

type ResourceState int

const (
	STATE_CREATED ResourceState = iota
	STATE_STARTED
	STATE_STOPPED
	STATE_RUNNING
	STATE_TERMINATED
)

type EngineState struct {
	PB     *pocketbase.PocketBase
	LVConn *libvirt.Connect

	// libvirt conn state
	// db conn state
	// probably pocketbase state?
	// mutexes?
}

var State EngineState

type IResourceState interface {
	CheckState() ResourceState
	SetState() ResourceState
}
