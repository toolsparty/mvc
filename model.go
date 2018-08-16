package mvc

import (
	"fmt"
	"github.com/pkg/errors"
)

type Model interface {
	Component
}

type models map[string]Model

func (list *models) add(models ...Model) error {
	c := *list

	for _, model := range models {
		name, err := model.Name()
		if err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("getting model %T name failed", model),
			)
		}

		c[name] = model
	}

	*list = c

	return nil
}

func (list models) setApp(app *App) {
	for _, model := range list {
		model.SetApp(app)
	}
}

type BaseModel struct {
	App *App
}

func (BaseModel) Name() (string, error) {
	return "", errors.New("model name undefined")
}

func (m *BaseModel) SetApp(app *App) {
	m.App = app
}
