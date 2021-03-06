package mvc

import (
	"fmt"

	"github.com/pkg/errors"
)

// v is should be context
type Action func(v interface{}) error

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
			action = controller.BeforeAction(action)
			c[path] = controller.AfterAction(action)
		}
	}

	*list = c

	return nil
}
