package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	purifyComments()
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

func purifyComments() {
	fmt.Println("\n[2/4] Purificando código: Removendo comentários (//)...")
	count := 0

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		if strings.Contains(path, ".git") || strings.Contains(path, "vendor") {
			return nil
		}

		if processFile(path) {
			count++
		}
		return nil
	})

	if err == nil {
		fmt.Printf("      OK: %d arquivos Go purificados para produção.\n", count)
	}
}

func processFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	var output bytes.Buffer
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//go:") || strings.HasPrefix(trimmed, "// +build") {
			output.WriteString(line)
			if err == io.EOF {
				break
			}
			continue
		}

		cleanLine := ""
		inString := false
		var stringChar rune

		runes := []rune(line)
		foundComment := false
		for i := 0; i < len(runes); i++ {
			r := runes[i]

			if (r == '"' || r == '`' || r == '\'') && (i == 0 || runes[i-1] != '\\') {
				if !inString {
					inString = true
					stringChar = r
				} else if stringChar == r {
					inString = false
				}
			}

			if !inString && i+1 < len(runes) && runes[i] == '/' && runes[i+1] == '/' {

				comment := string(runes[i:])
				lowerComment := strings.ToLower(comment)
				if strings.Contains(lowerComment, "nolint") || strings.Contains(lowerComment, "lint:") || strings.Contains(lowerComment, "noqa") {
					cleanLine = strings.TrimRight(line, "\r\n")
				}
				foundComment = true
				break
			}
			cleanLine += string(r)
		}

		if foundComment || strings.TrimSpace(cleanLine) != "" || strings.TrimSpace(line) == "" {
			output.WriteString(strings.TrimRight(cleanLine, "\r\n") + "\n")
		}

		if err == io.EOF {
			break
		}
	}

	_ = file.Close()
	_ = os.WriteFile(path, output.Bytes(), 0644)
	return true
}

func runLinting() {
	fmt.Println("\n[3/4] Executando Automação de Estilo, Padrões e Qualidade (CI Sync)...")

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
	fmt.Println("\n[4/4] Auditoria de Segurança e Vulnerabilidades (CI Sync)...")

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
