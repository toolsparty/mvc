package mvc

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type App struct {
	ctx context.Context

	config      Config
	router      Router
	controllers controllers
	models      models
	views       views
	actions     Actions

	debug bool
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Config() Config {
	return app.config
}

func (app *App) Actions() Actions {
	return app.actions
}

func (app *App) Log(args ...interface{}) {
	for _, arg := range args {
		if app.debug {
			_, isError := arg.(error)
			if isError {
				log.Println("Error:", arg)
				continue
			}
		}

		log.Println(args...)
	}
}

func (app *App) SetDebug(debug bool) {
	app.debug = debug
}

func (app *App) Run() error {
	err := app.controllers.init(app)
	if err != nil {
		return errors.Wrap(err, "initializing controllers failed")
	}

	err = app.router.Route(app)
	if err != nil {
		return errors.Wrap(err, "running application failed")
	}

	return nil
}

func (app *App) View(name string) (View, error) {
	if view, exists := app.views[name]; exists {
		return view, nil
	}

	return nil, errors.New("view '" + name + "' not found")
}

func CreateApp(config *AppConfig) (*App, error) {
	app := &App{
		config:      config.Config,
		router:      config.Router,
		controllers: make(controllers, len(config.Controllers)),
		models:      make(models, len(config.Models)),
		views:       make(views, len(config.Views)),
		actions:     make(Actions),
	}

	err := app.controllers.add(config.Controllers...)
	if err != nil {
		return nil, errors.Wrap(err, "adding controllers failed")
	}

	err = app.actions.add(config.Controllers...)
	if err != nil {
		return nil, errors.Wrap(err, "adding controllers actions failed")
	}

	err = app.models.add(config.Models...)
	if err != nil {
		return nil, errors.Wrap(err, "adding models failed")
	}

	err = app.views.add(config.Views...)
	if err != nil {
		return nil, errors.Wrap(err, "adding views failed")
	}

	return app, nil
}
