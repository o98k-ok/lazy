package alfred

import (
	"github.com/urfave/cli/v2"
)

type BindFunc func([]string)

type Application struct {
	app *cli.App
}

func NewApp(desc string) *Application {
	app := &Application{
		app: cli.NewApp(),
	}

	app.app.Description = desc
	return app
}

func (a *Application) Run(arguments []string) error {
	return a.app.Run(arguments)
}

func (a *Application) Bind(name string, fn BindFunc) *Application {
	a.app.Commands = append(a.app.Commands, &cli.Command{
		Name: name,
		Action: func(context *cli.Context) error {
			fn(context.Args().Slice())
			return nil
		},
	})
	return a
}

func (a *Application) DefaultBind(fn BindFunc) *Application {
	a.app.CommandNotFound = func(context *cli.Context, s string) {
		fn(context.Args().Slice())
	}
	return a
}
