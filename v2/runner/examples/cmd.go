package main

import (
	"fmt"
	"os/exec"

	"github.com/o98k-ok/lazy/v2/runner"
)

func main() {
	cmds := []*exec.Cmd{
		runner.NewCmdWithDir("/bin/ls -l", ""),
		runner.NewCmdWithDir("/bin/ls -l", "/"),
	}

	res, _ := runner.NewShellRunner(cmds...).Parallel(2)
	fmt.Println(res[0].Err, string(res[0].Elem))
	fmt.Println(res[1].Err, string(res[1].Elem))
}
