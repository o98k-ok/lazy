package alfred

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	t.Run("test new app", func(t *testing.T) {
		desc := "testing mode"
		app := NewApp(desc)
		assert.Equal(t, desc, app.app.Description)
	})

	cases := []struct {
		name    string
		cmdline string
		expect  []string
		success bool
	}{
		{
			"test single select",
			"./entry select server01",
			[]string{"server01"},
			true,
		},
		{
			"test double trace",
			"./entry trace server01 server03",
			[]string{"traceserver01", "traceserver03"},
			true,
		},
		{
			"test failed case",
			"./entry fuck server01 server03",
			[]string{},
			true,
		},
	}

	app := NewApp("")
	var got []string
	app.Bind("select", func(is []string) {
		got = is
	})
	app.Bind("trace", func(is []string) {
		is = lo.Map[string, string](is, func(t string, _ int) string {
			return "trace" + t
		})
		got = is
	})
	app.DefaultBind(func(i []string) {
		got = []string{}
	})

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := app.Run(strings.Split(c.cmdline, " "))
			assert.Equal(t, c.success, err == nil)
			assert.Equal(t, c.expect, got)
		})
	}
}
