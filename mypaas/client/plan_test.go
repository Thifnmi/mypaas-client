package client

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/thifnmi/mypaas/cmd"
	"github.com/thifnmi/mypaas/cmd/cmdtest"
	"gopkg.in/check.v1"
)

func (s *S) TestPlanListInfo(c *check.C) {
	c.Assert((&PlanList{}).Info(), check.NotNil)
}

func (s *S) TestPlanListBytes(c *check.C) {
	var stdout, stderr bytes.Buffer
	result := `[
	{"name": "test",  "memory": 536870912, "swap": 268435456, "cpushare": 100, "default": false},
	{"name": "test2", "memory": 536870912, "swap": 268435456, "cpushare": 200, "default": true}
]`
	expected := `+-------+-----------------+-----------+-----------+---------+
| Name  | CPU             | Memory    | Swap      | Default |
+-------+-----------------+-----------+-----------+---------+
| test  | 100 (CPU share) | 536870912 | 268435456 | false   |
| test2 | 200 (CPU share) | 536870912 | 268435456 | true    |
+-------+-----------------+-----------+-----------+---------+
`
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	trans := &cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Message: string(result), Status: http.StatusOK},
		CondFunc: func(req *http.Request) bool {
			return strings.HasSuffix(req.URL.Path, "/plans") && req.Method == "GET"
		},
	}
	client := cmd.NewClient(&http.Client{Transport: trans}, nil, manager)
	command := PlanList{}
	command.Flags().Parse(true, []string{"-b"})
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, expected)
}

func (s *S) TestPlanListHuman(c *check.C) {
	var stdout, stderr bytes.Buffer
	result := `[
	{"name": "test",  "memory": 536870912, "swap": 268435456, "cpushare": 100, "default": false},
	{"name": "test2", "memory": 536870912, "swap": 268435456, "cpushare": 200, "default": true}
]`
	expected := `+-------+-----------------+--------+-------+---------+
| Name  | CPU             | Memory | Swap  | Default |
+-------+-----------------+--------+-------+---------+
| test  | 100 (CPU share) | 512Mi  | 256Mi | false   |
| test2 | 200 (CPU share) | 512Mi  | 256Mi | true    |
+-------+-----------------+--------+-------+---------+
`
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	trans := &cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Message: string(result), Status: http.StatusOK},
		CondFunc: func(req *http.Request) bool {
			return strings.HasSuffix(req.URL.Path, "/plans") && req.Method == "GET"
		},
	}
	client := cmd.NewClient(&http.Client{Transport: trans}, nil, manager)
	command := PlanList{}
	// command.Flags().Parse(true, []string{"-h"})
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, expected)
}

func (s *S) TestPlanListOverride(c *check.C) {
	var stdout, stderr bytes.Buffer
	result := `[
	{"name": "test",  "memory": 536870912, "swap": 268435456, "cpushare": 100, "default": false, "override": {"cpumilli": 300, "memory": 268435456}}
]`
	expected := `+------+----------------+------------------+-------+---------+
| Name | CPU            | Memory           | Swap  | Default |
+------+----------------+------------------+-------+---------+
| test | 30% (override) | 256Mi (override) | 256Mi | false   |
+------+----------------+------------------+-------+---------+
`
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	trans := &cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Message: string(result), Status: http.StatusOK},
		CondFunc: func(req *http.Request) bool {
			return strings.HasSuffix(req.URL.Path, "/plans") && req.Method == "GET"
		},
	}
	client := cmd.NewClient(&http.Client{Transport: trans}, nil, manager)
	command := PlanList{}
	// command.Flags().Parse(true, []string{"-h"})
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, expected)
}

func (s *S) TestPlanListCPUMilli(c *check.C) {
	var stdout, stderr bytes.Buffer
	result := `[
	{"name": "test",  "memory": 536870912, "swap": 0, "cpumilli": 500, "default": false, "override": {"cpumilli": null, "memory": null}}
]`
	expected := `+------+-----+--------+---------+
| Name | CPU | Memory | Default |
+------+-----+--------+---------+
| test | 50% | 512Mi  | false   |
+------+-----+--------+---------+
`
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	trans := &cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Message: string(result), Status: http.StatusOK},
		CondFunc: func(req *http.Request) bool {
			return strings.HasSuffix(req.URL.Path, "/plans") && req.Method == "GET"
		},
	}
	client := cmd.NewClient(&http.Client{Transport: trans}, nil, manager)
	command := PlanList{}
	// command.Flags().Parse(true, []string{"-h"})
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, expected)
}

func (s *S) TestPlanListEmpty(c *check.C) {
	var stdout, stderr bytes.Buffer
	context := cmd.Context{
		Args:   []string{},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	trans := &cmdtest.ConditionalTransport{
		Transport: cmdtest.Transport{Message: "", Status: http.StatusNoContent},
		CondFunc: func(req *http.Request) bool {
			return strings.HasSuffix(req.URL.Path, "/plans") && req.Method == "GET"
		},
	}
	client := cmd.NewClient(&http.Client{Transport: trans}, nil, manager)
	command := PlanList{}
	err := command.Run(&context, client)
	c.Assert(err, check.IsNil)
	c.Assert(stdout.String(), check.Equals, "No plans available.\n")
}
