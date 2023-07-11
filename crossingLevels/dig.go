package crossingLevels

import (
	"fmt"
	"go.uber.org/dig"
)

type Greeter struct {
	name string
}

func NewGreeter(name string) *Greeter {
	return &Greeter{name: name}
}

func (g *Greeter) Greet() {
	fmt.Printf("Hello, %s!\n", g.name)
}

type App struct {
	greeter *Greeter
}

func NewApp(greeter *Greeter) *App {
	return &App{greeter: greeter}
}

func (a *App) Run() {
	a.greeter.Greet()
}

func T1() {
	container := dig.New()

	container.Provide(NewGreeter)
	container.Provide(NewApp)

	err := container.Invoke(func(app *App) {
		app.Run()
	})
	if err != nil {
		panic(err)
	}
}
