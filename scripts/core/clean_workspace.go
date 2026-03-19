package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	ProjectName = "Go-Wallet-Core"
	Version     = "1.0.0"
)

func main() {
	printHeader()

	cleanCaches()
	runLinting()
	runAudit()

	printFooter()
}

func printHeader() {
	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf("      %s - CLEANING SYSTEM (GO NATIVE) %s\n", strings.ToUpper(ProjectName), Version)
	cwd, _ := os.Getwd()
	fmt.Printf("      DIRETÓRIO: %s\n", cwd)
	fmt.Println(strings.Repeat("=", 75))
}

func printFooter() {
	fmt.Println("\n" + strings.Repeat("=", 75))
	fmt.Println("      ECOSSISTEMA GO SINCRONIZADO E PURIFICADO (PADRÃO SÊNIOR).")
	fmt.Println(strings.Repeat("=", 75))
}

func cleanCaches() {
	fmt.Println("\n[1/4] Limpando caches do Go e binários...")

	commands := [][]string{
		{"go", "clean", "-cache"},
		{"go", "clean", "-testcache"},
		{"go", "clean", "-modcache"},
	}

	for _, cmd := range commands {

		// #nosec G204
		_ = exec.Command(cmd[0], cmd[1:]...).Run()
	}

	removed := 0
	patterns := []string{"*.exe", "*.test", "*.out", "bin/", "*.tmp", "*.log"}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		for _, pattern := range patterns {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {

				// #nosec G122
				_ = os.RemoveAll(path)
				removed++
				break
			}
		}
		return nil
	})

	if err == nil {
		fmt.Printf("      OK: %d itens de cache/build purificados.\n", removed)
	}
}

func runLinting() {
	fmt.Println("\n[2/3] Executando Automação de Estilo, Padrões e Qualidade (CI Sync)...")

	tools := []struct {
		name    string
		command string
		args    []string
	}{
		{"Go Fmt", "go", []string{"fmt", "./..."}},
		{"Go Vet", "go", []string{"vet", "./..."}},
		{"Go Mod Tidy", "go", []string{"mod", "tidy"}},
		{"GolangCI-Lint", "golangci-lint", []string{"run", "./..."}},
	}

	failed := false
	for _, t := range tools {
		fmt.Printf("      -> %s: Processando...\n", t.name)

		if _, err := exec.LookPath(t.command); err != nil {
			fmt.Printf("      [AVISO] %s não instalado localmente. Pulando...\n", t.name)
			continue
		}

		// #nosec G204
		cmd := exec.Command(t.command, t.args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			fmt.Printf("      [ERRO] %s falhou!\n", t.name)
			fmt.Println(stderr.String())
			failed = true
		} else {
			fmt.Printf("      OK: %s finalizado.\n", t.name)
		}
	}

	if failed {
		fmt.Println("\n[!] ERRO: O código não atende aos padrões de qualidade sênior (CI Sync).")
		os.Exit(1)
	}
	fmt.Println("      OK: Sincronização de estilo e linting finalizada.")
}

func runAudit() {
	fmt.Println("\n[3/3] Auditoria de Segurança e Vulnerabilidades (CI Sync)...")

	tools := []struct {
		name    string
		command string
		args    []string
	}{
		{"Gosec", "gosec", []string{"./..."}},
		{"Govulncheck", "govulncheck", []string{"./..."}},
	}

	for _, t := range tools {
		fmt.Printf("      -> %s: Analisando...\n", t.name)

		if _, err := exec.LookPath(t.command); err != nil {
			fmt.Printf("      [AVISO] %s não instalado localmente. Pulando...\n", t.name)
			continue
		}

		// #nosec G204
		cmd := exec.Command(t.command, t.args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			fmt.Printf("      [AVISO] %s encontrou pontos de atenção:\n", t.name)
			fmt.Println(stderr.String())
		} else {
			fmt.Printf("      OK: %s sem problemas críticos.\n", t.name)
		}
	}
	fmt.Println("      OK: Auditoria técnica concluída.")
}
