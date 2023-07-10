package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jptalukdar/go-ci-agent/pkg/job"
	"github.com/jptalukdar/go-ci-agent/pkg/sources/git"
)

func main() {

	ctx := context.Background()
	// err := p.Execute(ctx, c)
	// if err != nil {
	// 	log.Printf("Error in executing %s", err)
	// }
	// e := job.Execution{
	// 	Source: &git.GitSource{
	// 		Url:    "https://github.com/realpython/python-scripts.git",
	// 		Branch: "master",
	// 	},
	// 	Program: job.Program{
	// 		Executable: "python",
	// 		Args:       []string{"29_json_to_yaml.py", "29_json_test.json"},
	// 		Dir:        "scripts", // Relative to git download path
	// 	},
	// 	Capture: job.Capture{
	// 		Stdout:     true,
	// 		StdoutFile: "stdoutput.io",
	// 		Stderr:     true,
	// 		StderrFile: "stderr.io",
	// 		Console:    true,
	// 	},
	// 	Config: job.Config{
	// 		BaseDir: "tmp",
	// 	},
	// }

	var e job.Execution
	e.Source = &git.GitSource{} // This line prevents json unmarshalling issue

	Data := []byte(`{"source" : {
		"url" : "https://github.com/realpython/python-scripts.git",
		"branch" : "master"
	},
	"program" : {
		"executable" : "python",
		"args" : ["29_json_to_yaml.py" , "29_json_test.json"],
		"dir" : "scripts"
	},
	"capture" : {
		"stdout" : true,
		"stdoutFile" : "stdout.io",
		"stderr" : true,
		"stderrFile" : "stderr.io",
		"console" : true
	},
	"config" : {
		"baseDir" : "/tmp/"
	}}`)

	// decoding country1 struct
	// from json format
	err := json.Unmarshal(Data, &e)
	if err != nil {
		log.Fatal(err)
	}
	e.Run(ctx)
}
