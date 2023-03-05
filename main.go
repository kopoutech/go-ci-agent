package main

import (
	"context"

	"github.com/jptalukdar/go-ci-agent/pkg/job"
	"github.com/jptalukdar/go-ci-agent/pkg/sources/git"
)

func main() {

	ctx := context.Background()
	// err := p.Execute(ctx, c)
	// if err != nil {
	// 	log.Printf("Error in executing %s", err)
	// }
	e := job.Execution{
		Source: git.GitSource{
			Url:    "https://github.com/realpython/python-scripts.git",
			Branch: "master",
		},
		Program: job.Program{
			Executable: "python",
			Args:       []string{"29_json_to_yaml.py", "29_json_test.json"},
			Dir:        "tmp/scripts",
		},
		Capture: job.Capture{
			Stdout:     true,
			StdoutFile: "stdoutput",
			Stderr:     true,
			StderrFile: "stderr",
		},
		Config: job.Config{
			BaseDir: "tmp",
		},
	}
	e.Run(ctx)
}
