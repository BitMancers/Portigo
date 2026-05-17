package engine

type VM struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	CPU   int    `json:"cpu"`
	RAM   int    `json:"ram"`
	State State  `json:"vm_state"`
}
