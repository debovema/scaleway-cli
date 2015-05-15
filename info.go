package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kardianos/osext"
)

var cmdInfo = &Command{
	Exec:        runInfo,
	UsageLine:   "info",
	Description: "Display system-wide information",
	Help:        "Display system-wide information.",
}

func runInfo(cmd *Command, args []string) {
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "usage: scw %s\n", cmd.UsageLine)
		os.Exit(1)
	}

	// FIXME: fmt.Printf("Servers: %s\n", "quantity")
	// FIXME: fmt.Printf("Images: %s\n", "quantity")
	fmt.Printf("Debug mode (client): %v\n", os.Getenv("DEBUG") != "")

	fmt.Printf("Organization: %s\n", config.Organization)
	// FIXME: add partially-masked token
	fmt.Printf("API Endpoint: %s\n", os.Getenv("scaleway_api_endpoint"))
	configPath, _ := GetConfigFilePath()
	fmt.Printf("RC file: %s\n", configPath)
	fmt.Printf("User: %s\n", os.Getenv("USER"))
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	hostname, _ := os.Hostname()
	fmt.Printf("Hostname: %s\n", hostname)
	cliPath, _ := osext.Executable()
	fmt.Printf("CLI Path: %s\n", cliPath)

	fmt.Printf("Cache: %s\n", cmd.API.Cache.Path)
	fmt.Printf("  Servers: %d\n", cmd.API.Cache.GetNbServers())
	fmt.Printf("  Images: %d\n", cmd.API.Cache.GetNbImages())
	fmt.Printf("  Snapshots: %d\n", cmd.API.Cache.GetNbSnapshots())
}