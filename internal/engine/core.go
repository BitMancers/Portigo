package engine

import (
	"log"
	"portigo/data"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	_ "portigo/migrations"
)

var app *pocketbase.PocketBase = nil

func Init() {
	app = pocketbase.New()

	data.RunMigrations(app)
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// registers new "GET /hello" route
		se.Router.GET("/hello", func(re *core.RequestEvent) error {
			return re.String(200, "Hello world!")
		})

		return se.Next()
	})
}
func ShutdownEngine() {
	app.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		// e.App
		// e.IsRestart
		log.Println("App is shutting down")
		return e.Next()
	})
}
func Start() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
func HealthCheck() {}
