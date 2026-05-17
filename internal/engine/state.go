package engine

type State int

const (
	STATE_CREATED State = iota
	STATE_STARTED
	STATE_STOPPED
	STATE_RUNNING
	STATE_TERMINATED
)

type IState interface {
	CheckState() State
	SetState() State
}
