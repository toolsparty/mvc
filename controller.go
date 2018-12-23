package mvc

import (
	"fmt"

	"github.com/pkg/errors"
)

type Controller interface {
	Component

	Init() error
	Actions() (Actions, error)
	BeforeAction(action Action) Action
	AfterAction(action Action) Action
}

type controllers map[string]Controller

func (list *controllers) add(controllers ...Controller) error {
	c := *list

	for _, controller := range controllers {
		name, err := controller.Name()
		if err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("getting controller %T name failed", controller),
			)
		}

		c[name] = controller
	}

	*list = c

	return nil
}

func (list controllers) init(app *App) error {
	for _, controller := range list {
		controller.SetApp(app)

		if err := controller.Init(); err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("initializing controller %T failed", controller),
			)
		}
	}

	return nil
}

type BaseController struct {
	App *App
}

func (BaseController) Name() (string, error) {
	return "", errors.New("controller name undefined")
}

func (c *BaseController) SetApp(app *App) {
	c.App = app
}

func (BaseController) Actions() (Actions, error) {
	return nil, errors.New("controller actions undefined")
}

func (BaseController) Init() error {
	return nil
}

func (BaseController) BeforeAction(action Action) Action {
	return action
}

func (BaseController) AfterAction(action Action) Action {
	return action
}
