package job_test

import (
	"context"
	"log"
	"testing"

	"github.com/jptalukdar/go-ci-agent/pkg/job"
)

func TestProgram(t *testing.T) {
	p := job.Program{
		Executable: "echo",
		Args:       []string{"hello", "there"},
	}

	c := job.Capture{
		Stdout:     true,
		StdoutFile: "stdoutput",
	}

	ctx := context.Background()
	err := p.Execute(ctx, c)

	log.Printf(err.Error())

}
