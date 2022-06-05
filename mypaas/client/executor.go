//go:build !windows
// +build !windows

package client

import (
	"github.com/thifnmi/mypaas/exec"
)

var Execut exec.Executor

func Executor() exec.Executor {
	if Execut == nil {
		Execut = exec.OsExecutor{}
	}
	return Execut
}
