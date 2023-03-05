package job

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
)

func (p *Program) Execute(ctx context.Context, c Capture) error {
	cmd := exec.Command(p.Executable, p.Args...)

	cmd.Dir = p.Dir
	if c.Stdout {
		file, _ := os.Create(c.StdoutFile)
		defer file.Close()
		stdwriter := bufio.NewWriter(file)
		cmd.Stdout = stdwriter
	}

	if c.Stderr {
		file2, _ := os.Create(c.StderrFile)
		defer file2.Close()
		errwriter := bufio.NewWriter(file2)
		cmd.Stderr = errwriter
	}

	log.Printf("%+v", cmd)
	if err := cmd.Run(); err != nil {
		log.Println("could not run command: ", err)
		return err
	}

	return nil
}

func (e *Execution) Run(ctx context.Context) error {
	ok, err := e.Source.Download(e.Config.BaseDir)
	if err != nil {
		log.Printf("Error in fetching source : " + err.Error())
		return err
	}
	if !ok {
		log.Printf("Unable to fetch source")
	}

	err = e.Program.Execute(ctx, e.Capture)
	if err != nil {
		log.Printf("Error in executing program")
		return err
	}
	return nil
}
