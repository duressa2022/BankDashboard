package config

import (
	"working.com/bank_dash/package/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() (Application, error) {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		return Application{}, err
	}
	app.Env = env
	app.Mongo = NewMongoDatabase(app.Env)
	return *app, nil
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
