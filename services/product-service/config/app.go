package config

import "go.mongodb.org/mongo-driver/mongo"

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = initEnv()
	app.Mongo = OpenMongoDBConnection(*app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(&app.Mongo)
}
