package main

import "runtime"

// Version e Commit sao preenchidos em release com go build -ldflags (ver README).
// Em desenvolvimento permanecem "dev" se nao injetados.
var (
	Version = "dev"
	Commit  = "dev"
)

func goRuntimeVersion() string {
	return runtime.Version()
}
