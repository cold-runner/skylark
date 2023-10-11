package main

import (
	"github.com/cold-runner/skylark/internal/core"
)

func main() {
	core.InitApp()
	core.App.InstallRouter().Run()
}
