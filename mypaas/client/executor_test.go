//go:build !windows
// +build !windows

package client

import (
	"github.com/thifnmi/mypaas/exec"
	"github.com/thifnmi/mypaas/exec/exectest"
	"gopkg.in/check.v1"
)

func (s *S) TestExecutor(c *check.C) {
	Execut = &exectest.FakeExecutor{}
	c.Assert(Executor(), check.DeepEquals, Execut)
	Execut = nil
	c.Assert(Executor(), check.DeepEquals, exec.OsExecutor{})
}
