package engine

import (
	"log"
	"os"

	"github.com/BitMancers/Portigo/data"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"libvirt.org/go/libvirt"
)

func Init(lvConn *libvirt.Connect) {
	State.PB = pocketbase.New()
	State.LVConn = lvConn

	data.RunMigrations(State.PB)

	State.PB.OnServe().BindFunc(func(se *core.ServeEvent) error {

		se.Router.GET("/hello", func(re *core.RequestEvent) error {
			return re.String(200, "Hello world!")
		})

		se.Router.GET("/{path...}", apis.Static(os.DirFS("ui/build"), false))

		return se.Next()
	})
}
func Start() {
	if err := State.PB.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func Shutdown() {
	State.PB.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		// e.App
		// e.IsRestart
		log.Println("App is shutting down")
		return e.Next()
	})
}

func HealthCheck() {
}
