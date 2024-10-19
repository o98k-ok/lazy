package alfred

import (
	"github.com/urfave/cli/v2"
)

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

func (a *Application) BindMore(name string, fn ComplexBindFnc, ops ...OptionSetting) *Application {
	cmd := &cli.Command{
		Name: name,
	}

	for _, op := range ops {
		op(cmd)
	}

	cmd.Action = func(ctx *cli.Context) error {
		value := make(map[string]interface{})
		for _, flag := range cmd.Flags {
			name := flag.Names()[0]
			switch flag.(type) {
			case *cli.Int64Flag:
				Pack(value, name, ctx.Int64(name))
			case *cli.StringFlag:
				Pack(value, name, ctx.String(name))
			case *cli.BoolFlag:
				Pack(value, name, ctx.Bool(name))
			case *cli.Int64SliceFlag:
				Pack(value, name, ctx.Int64Slice(name))
			case *cli.StringSliceFlag:
				Pack(value, name, ctx.StringSlice(name))
			default:
				panic("unsupport flags...")
			}
		}
		fn(value)
		return nil
	}

	a.app.Commands = append(a.app.Commands, cmd)
	return a
}

func (a *Application) DefaultBind(fn BindFunc) *Application {
	a.app.CommandNotFound = func(context *cli.Context, s string) {
		fn(context.Args().Slice())
	}
	return a
}

func GetParamByIdx(args []string, idx int) string {
	if len(args) <= 2 {
		return ""
	}

	args = args[2:]
	if len(args) <= idx {
		return ""
	}
	return args[idx]
}
