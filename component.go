package mvc

type Component interface {
	Name() (string, error)
	SetApp(app *App)
}
