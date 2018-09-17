package mvc

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

type ViewParams map[string]interface{}

type View interface {
	Component

	Render(w io.Writer, tpl string, params ViewParams) error
}

type views map[string]View

func (list *views) add(views ...View) error {
	c := *list

	for _, view := range views {
		name, err := view.Name()
		if err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("getting model %T name failed", view),
			)
		}

		c[name] = view
	}

	*list = c

	return nil
}

type BaseView struct {
	App *App
}

func (BaseView) Name() (string, error) {
	return "", errors.New("view name undefined")
}

func (v *BaseView) SetApp(app *App) {
	v.App = app
}
