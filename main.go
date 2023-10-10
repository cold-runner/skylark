package main

import (
	"github.com/cold-runner/skylark/internal/core"
)

func main() {
	core.DependencyInjection()
	core.App.InstallRouter().Run()
}
