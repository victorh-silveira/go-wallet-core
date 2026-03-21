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

		commandPath := t.command
		if _, err := exec.LookPath(commandPath); err != nil {
			if t.command == "golangci-lint" {
				fmt.Println("      [INFO] Instalando golangci-lint...")
				if err := ensureToolInstalled("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"); err != nil {
					fmt.Printf("      [ERRO] Falha ao instalar golangci-lint: %v\n", err)
					failed = true
					continue
				}
			}
			if resolved, ok := resolveGoToolPath(t.command); ok {
				commandPath = resolved
			} else if _, err := exec.LookPath(commandPath); err != nil {
				fmt.Printf("      [ERRO] %s não instalado localmente.\n", t.name)
				failed = true
				continue
			}
		}

		// #nosec G204
		cmd := exec.Command(commandPath, t.args...)
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

	failed := false
	for _, t := range tools {
		fmt.Printf("      -> %s: Analisando...\n", t.name)

		commandPath := t.command
		if _, err := exec.LookPath(commandPath); err != nil {
			if t.command == "gosec" {
				fmt.Println("      [INFO] Instalando gosec...")
				if err := ensureToolInstalled("go", "install", "github.com/securego/gosec/v2/cmd/gosec@latest"); err != nil {
					fmt.Printf("      [ERRO] Falha ao instalar gosec: %v\n", err)
					failed = true
					continue
				}
			}
			if t.command == "govulncheck" {
				fmt.Println("      [INFO] Instalando govulncheck...")
				if err := ensureToolInstalled("go", "install", "golang.org/x/vuln/cmd/govulncheck@latest"); err != nil {
					fmt.Printf("      [ERRO] Falha ao instalar govulncheck: %v\n", err)
					failed = true
					continue
				}
			}
			if resolved, ok := resolveGoToolPath(t.command); ok {
				commandPath = resolved
			} else if _, err := exec.LookPath(commandPath); err != nil {
				fmt.Printf("      [ERRO] %s não instalado localmente.\n", t.name)
				failed = true
				continue
			}
		}

		// #nosec G204
		cmd := exec.Command(commandPath, t.args...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			fmt.Printf("      [ERRO] %s encontrou pontos de atenção:\n", t.name)
			fmt.Println(stderr.String())
			failed = true
		} else {
			fmt.Printf("      OK: %s sem problemas críticos.\n", t.name)
		}
	}
	if failed {
		fmt.Println("\n[!] ERRO: Auditoria de segurança falhou.")
		os.Exit(1)
	}
	fmt.Println("      OK: Auditoria técnica concluída.")
}

func ensureToolInstalled(command string, args ...string) error {
	// #nosec G204
	cmd := exec.Command(command, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %s", err, stderr.String())
	}
	return nil
}

func resolveGoToolPath(toolName string) (string, bool) {
	// #nosec G204
	cmd := exec.Command("go", "env", "GOPATH")
	output, err := cmd.Output()
	if err != nil {
		return "", false
	}
	gopath := strings.TrimSpace(string(output))
	if gopath == "" {
		return "", false
	}
	toolPath := filepath.Join(gopath, "bin", toolName)
	if _, err := os.Stat(toolPath); err == nil {
		return toolPath, true
	}
	toolPathExe := toolPath + ".exe"
	if _, err := os.Stat(toolPathExe); err == nil {
		return toolPathExe, true
	}
	return "", false
}
