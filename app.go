package mvc

import (
	"context"

	"github.com/pkg/errors"
)

type App struct {
	ctx context.Context

	config Config
	router Router
	logger LogFunc

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
	if app.logger != nil {
		app.logger(args...)
		return
	}

	defaultLogFn(args...)
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

func (app *App) View(name string) View {
	if view, exists := app.views[name]; exists {
		return view
	}

	return nil
}

func (app *App) Model(name string) Model {
	if model, exists := app.models[name]; exists {
		return model
	}

	return nil
}

func CreateApp(config *AppConfig) (*App, error) {
	app := &App{
		ctx: config.Context,

		config: config.Config,
		router: config.Router,
		logger: config.Logger,

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
