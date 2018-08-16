package mvc

type Router interface {
	Route(app *App) error
}
