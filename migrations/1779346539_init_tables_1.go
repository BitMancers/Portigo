package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		// add up queries...

		vmCollection := core.NewBaseCollection("vm")
		_ = vmCollection

		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
