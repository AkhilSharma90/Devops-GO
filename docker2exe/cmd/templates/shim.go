package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/mattn/go-isatty"
)

type Shim struct {
	Image   string
	Workdir string
	Env     []string
	Volumes []string
	Stdout  io.Writer
	Stderr  io.Writer
}

func (shim *Shim) Exists() bool {
	cmd := shim.docker("inspect", shim.Image)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run() == nil
}

func (shim *Shim) Pull() error {
	return shim.docker("pull", shim.Image).Run()
}

func (shim *Shim) Load(file io.Reader) error {
	cmd := shim.docker("load")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = io.Copy(stdin, file)
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (shim *Shim) Exec(containerArgs []string) error {
	args := []string{"run", "--rm"}

	if isatty.IsTerminal(os.Stdout.Fd()) {
		args = append(args, "-it")
	}

	for _, env := range shim.Env {
		args = append(args, "-e", env)
	}

	for _, volume := range shim.Volumes {
		args = append(args, "-v", volume)
	}

	if shim.Workdir != "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		args = append(args, "-w", shim.Workdir)
		args = append(args, "-v", fmt.Sprintf("%s:%s", cwd, shim.Workdir))
	}

	args = append(args, shim.Image)
	args = append(args, containerArgs...)

	cmd := shim.docker(args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Executing Docker command: docker %s\n", strings.Join(args, " "))
	return cmd.Run()
}

func (shim *Shim) docker(args ...string) *exec.Cmd {
	dockerCmd := "docker"
	if runtime.GOOS == "windows" {
		dockerCmd = "docker.exe"
	}

	if path := os.Getenv("DOCKER"); path != "" {
		dockerCmd = path
	}

	cmd := exec.Command(dockerCmd, args...)
	cmd.Stdout = shim.Stdout
	cmd.Stderr = shim.Stderr
	return cmd
}
