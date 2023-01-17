package alfred

import "github.com/urfave/cli/v2"

type BindFunc func([]string)

type BindArg map[string]interface{}

func Get[T int64 | []int64 | string | bool | []string](arg BindArg, name string) (T, bool) {
	v, ok := arg[name].(T)
	return v, ok
}

func Pack[T int64 | []int64 | string | bool | []string](arg BindArg, name string, val T) {
	arg[name] = val
}

type ComplexBindFnc func(BindArg)

type OptionSetting func(cmd *cli.Command)

func Usage(usage string) OptionSetting {
	return func(cmd *cli.Command) {
		cmd.Usage = usage
	}
}

func ImportIntFlag(name, desc string, defaults *int64) OptionSetting {
	if defaults == nil {
		return func(cmd *cli.Command) {
			cmd.Flags = append(cmd.Flags, &cli.Int64Flag{
				Name:  name,
				Usage: desc,
			})
		}
	}
	return func(cmd *cli.Command) {
		cmd.Flags = append(cmd.Flags, &cli.Int64Flag{
			Name:  name,
			Value: *defaults,
			Usage: desc,
		})
	}
}

func ImportStringFlag(name, desc string, defaults *string) OptionSetting {
	if defaults == nil {
		return func(cmd *cli.Command) {
			cmd.Flags = append(cmd.Flags, &cli.StringFlag{
				Name:  name,
				Usage: desc,
			})
		}
	}

	return func(cmd *cli.Command) {
		cmd.Flags = append(cmd.Flags, &cli.StringFlag{
			Name:  name,
			Value: *defaults,
			Usage: desc,
		})
	}
}

func ImportBoolFlag(name, desc string, defaults *bool) OptionSetting {
	if defaults == nil {
		return func(cmd *cli.Command) {
			cmd.Flags = append(cmd.Flags, &cli.BoolFlag{
				Name:  name,
				Usage: desc,
			})
		}
	}

	return func(cmd *cli.Command) {
		cmd.Flags = append(cmd.Flags, &cli.BoolFlag{
			Name:  name,
			Value: *defaults,
			Usage: desc,
		})
	}
}

func ImportStringSliceFlag(name, desc string) OptionSetting {
	return func(cmd *cli.Command) {
		cmd.Flags = append(cmd.Flags, &cli.StringSliceFlag{
			Name:  name,
			Usage: desc,
		})
	}
}

func ImportIntSliceFlag(name, desc string) OptionSetting {
	return func(cmd *cli.Command) {
		cmd.Flags = append(cmd.Flags, &cli.Int64SliceFlag{
			Name:  name,
			Usage: desc,
		})
	}
}
