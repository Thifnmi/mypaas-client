package client

import (
	"github.com/thifnmi/mypaas/fs"
	"github.com/thifnmi/mypaas/fs/fstest"
	"gopkg.in/check.v1"
)

func (s *S) TestFileSystem(c *check.C) {
	fsystem = &fstest.RecordingFs{}
	c.Assert(filesystem(), check.DeepEquals, fsystem)
	fsystem = nil
	c.Assert(filesystem(), check.DeepEquals, fs.OsFs{})
}
