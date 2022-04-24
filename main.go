package main

import (
	"math/rand"
	"runtime"
	"smartparking/cmd"
	"smartparking/pkg/banner"
	"time"
)

// @title SMART PARKING API
// @version 0.0.1
// @description API documentation of smartparking service
// @contact.name API Support
// @contact.email kuanysheveldar123@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1/
func main() {
	banner.Default("assets/banner.txt.tmpl", map[string]interface{}{
		"now":      time.Now().Format(time.ANSIC),
		"numCPU":   runtime.NumCPU(),
		"GOOS":     runtime.GOOS,
		"GOARCH":   runtime.GOARCH,
		"Compiler": runtime.Compiler,
	})

	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
