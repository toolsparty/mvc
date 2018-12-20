package mvc

import "log"

type LogFunc func(args ...interface{})

func defaultLogFn(args ...interface{}) {
	var out []interface{}
	for _, arg := range args {
		_, isError := arg.(error)
		if isError {
			log.Println("Error:", arg)
			continue
		}

		out = append(out, arg)
	}

	if len(out) > 0 {
		log.Println(out...)
	}
}
