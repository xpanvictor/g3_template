package bootstrap

import "github.com/xpanvictor/g3_template.git/pkg/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

// App Initialize the configurations for app
func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (a *Application) CleanUp() {
	// clean db
	CloseMongoDBConnection(a.Mongo)
}
