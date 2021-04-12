package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	dir, err := os.MkdirTemp("", "gorapp")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	defer os.RemoveAll(dir)

	src, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	err = os.WriteFile(dir+"/main.go", src, 0600)
	if err != nil {
		log.Println("could not write file:", err)
		return
	}

	run(false, dir, "go", "version")
	run(false, dir, "go", "mod", "init", "rungo.local/run")
	run(false, dir, "go", "mod", "tidy")

	args := []string{"run", "main.go"}
	for _, v := range os.Args[2:] {
		args = append(args, v)
	}
	err = run(true, dir, "go", args...)

	if err != nil {
		log.Println("error:", err)
		return
	}
}

func run(output bool, dir, command string, args ...string) error {
	//fmt.Printf("running %v: %v %v\n", dir, command, args)
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	if output == true {
		commandOut, _ := cmd.StdoutPipe()
		errOut, _ := cmd.StderrPipe()
		go func() {
			io.Copy(os.Stdout, commandOut)
		}()
		go func() {
			io.Copy(os.Stderr, errOut)
		}()
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
