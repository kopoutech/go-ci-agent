package job

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (p *Program) Execute(ctx context.Context, c Capture) error {
	cmd := exec.Command(p.Executable, p.Args...)

	cmd.Dir = filepath.Join(p.baseExecutableDirectory, p.Dir)
	if c.Stdout {
		file, _ := os.Create(c.StdoutFile)
		defer file.Close()
		stdwriter := bufio.NewWriter(file)
		if c.Console {
			w1 := io.MultiWriter(os.Stdout, stdwriter)
			cmd.Stdout = w1
		} else {
			cmd.Stdout = stdwriter
		}
	}

	if c.Stderr {
		file2, _ := os.Create(c.StderrFile)
		defer file2.Close()
		errwriter := bufio.NewWriter(file2)
		if c.Console {
			w := io.MultiWriter(os.Stderr, errwriter)
			cmd.Stderr = w
		} else {
			cmd.Stderr = errwriter
		}
	}

	log.Printf("%+v", cmd)
	if err := cmd.Run(); err != nil {
		log.Println("could not run command: ", err)
		return err
	}

	return nil
}

func (e *Execution) Run(ctx context.Context) error {
	sp := e.Source.SetPath(e.Config.BaseDir)

	ok, err := e.Source.Download()
	if err != nil {
		log.Printf("Error in fetching source : " + err.Error())
		return err
	}
	if !ok {
		log.Printf("Unable to fetch source")
	}

	e.Program.SetExecutionPath(sp)
	err = e.Program.Execute(ctx, e.Capture)
	if err != nil {
		log.Printf("Error in executing program")
		return err
	}
	return nil
}
