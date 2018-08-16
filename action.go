package mvc

import (
	"fmt"

	"github.com/pkg/errors"
)

// should be context
type Action func(ctx interface{}) error

type Actions map[string]Action

func (list *Actions) add(controllers ...Controller) error {
	c := *list

	for _, controller := range controllers {
		actions, err := controller.Actions()
		if err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("getting actions for controller %T failed", controller),
			)
		}

		for path, action := range actions {
			_, ok := c[path]
			if ok {
				return errors.New(
					fmt.Sprintf("action %T is exists", action),
				)
			}
			c[path] = action
		}
	}

	*list = c

	return nil
}
