package cmd

import (
	"embed"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateRoot embed.FS
var templates = template.Must(template.ParseFS(templateRoot, "*/*"))

type Generator struct {
	Name    string
	Output  string
	Targets []string
	Module  string
	Image   string
	Embed   bool
	Workdir string
	Env     []string
	Volumes []string
}

func (gen *Generator) Run() error {
	// Create output directory if it doesn't exist
	outputDir := filepath.FromSlash(gen.Output)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	tmp, err := os.MkdirTemp("", gen.Name)
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	if err := gen.copyTemplates(tmp); err != nil {
		return err
	}
	return make(tmp)
}

func (gen *Generator) copyTemplates(dest string) error {
	for _, template := range templates.Templates() {
		filename := strings.TrimRight(template.Name(), ".tmpl")
		filepath := filepath.Join(dest, filename)

		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = template.Execute(file, gen)
		if err != nil {
			return err
		}
	}

	return nil
}

func make(cwd string) error {
	var cmd *exec.Cmd
	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "make")
	} else {
		cmd = exec.Command("make")
	}
	cmd.Dir = cwd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
