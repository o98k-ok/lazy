package runner

import (
	"os/exec"
)

type ShellOut []byte

type ShellRunner struct {
	*RunnerImp[ShellOut]
}

func NewCmdWithDir(cmd, dir string) *exec.Cmd {
	c := exec.Command("bash", "-c", cmd)
	if len(dir) != 0 {
		c.Dir = dir
	}
	return c
}

func NewShellRunner(cmds ...*exec.Cmd) Runner[ShellOut] {
	var tasks []func() (ShellOut, error)
	//for _, s := range cmds {
	for i := 0; i < len(cmds); i++ {
		cc := cmds[i]
		tasks = append(tasks, func(index int) func() (ShellOut, error) {
			return func() (ShellOut, error) {
				rr, e := cc.Output()
				return rr, e
			}
		}(i))
	}
	return &ShellRunner{
		RunnerImp: &RunnerImp[ShellOut]{
			Tasks: tasks,
		},
	}
}
