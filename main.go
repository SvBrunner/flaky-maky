package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/SvBrunner/flaky-maky/internal/fileops"
	"github.com/SvBrunner/flaky-maky/internal/inputs"
	"github.com/SvBrunner/flaky-maky/internal/models"
	"github.com/SvBrunner/flaky-maky/internal/registry"
)

func initInputs() tea.Model {
	flake := models.Flake{}
	inputsList := []inputs.Input{
		&inputs.NameInput{},
		&inputs.ChannelInput{},
		&inputs.SystemInput{},
		&inputs.PreConfigInput{},
		&inputs.DirenvInput{},
		&inputs.FinalInput{},
	}
	for index, input := range inputsList {
		if len(inputsList) > index+1 {
			input.InitInput(&flake, inputsList[index+1])
		} else {
			input.InitInput(&flake, nil)
		}
	}

	return inputsList[0]

}
func main() {
	// Check if we have subcommands
	if len(os.Args) > 1 && os.Args[1] == "servers" {
		handleServersCommand()
		return
	}

	syncFlag := flag.Bool("sync", false, "Sync configurations from registry servers")
	flag.Parse()

	// If sync flag is provided, fetch configs from servers
	if *syncFlag {
		if err := fileops.PopulatePreconfigs(); err != nil {
			fmt.Printf("Error preparing config directory: %v\n", err)
			os.Exit(1)
		}
		if err := fileops.SyncPreconfigs(); err != nil {
			fmt.Printf("Error during sync: %v\n", err)
			os.Exit(1)
		}
	}

	// Check if we have any local configs available
	preconfigs, err := fileops.ReadPreconfigurations()
	if err != nil || len(preconfigs) == 0 {
		fmt.Println("Error: No configurations found. Please run 'flaky-maky --sync' first to download configurations from registry servers.")
		os.Exit(1)
	}

	if err := fileops.PopulatePreconfigs(); err != nil {
		fmt.Printf("Error preparing config directory: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(initInputs())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func handleServersCommand() {
	if len(os.Args) < 3 {
		printServersUsage()
		os.Exit(1)
	}

	subcommand := os.Args[2]

	switch subcommand {
	case "add":
		handleServersAdd()
	case "disable":
		handleServersDisable()
	case "enable":
		handleServersEnable()
	case "list":
		handleServersList()
	case "delete":
		handleServersDelete()
	default:
		fmt.Printf("Unknown servers subcommand: %s\n", subcommand)
		printServersUsage()
		os.Exit(1)
	}
}

func handleServersAdd() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	name := addCmd.String("name", "", "Name of the server")
	url := addCmd.String("url", "", "URL of the server")

	addCmd.Parse(os.Args[3:])

	if *name == "" || *url == "" {
		fmt.Println("Error: Both -name and -url are required")
		fmt.Println("Usage: flaky-maky servers add -name <name> -url <url>")
		os.Exit(1)
	}

	if err := registry.AddServer(*name, *url); err != nil {
		fmt.Printf("Error adding server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Server '%s' added successfully\n", *name)
}

func handleServersDisable() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Server name is required")
		fmt.Println("Usage: flaky-maky servers disable <name>")
		os.Exit(1)
	}

	name := os.Args[3]

	if err := registry.DisableServer(name); err != nil {
		fmt.Printf("Error disabling server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Server '%s' disabled successfully\n", name)
}

func handleServersEnable() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Server name is required")
		fmt.Println("Usage: flaky-maky servers enable <name>")
		os.Exit(1)
	}

	name := os.Args[3]

	if err := registry.EnableServer(name); err != nil {
		fmt.Printf("Error enabling server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Server '%s' enabled successfully\n", name)
}

func handleServersDelete() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Server name is required")
		fmt.Println("Usage: flaky-maky servers delete <name>")
		os.Exit(1)
	}

	name := os.Args[3]

	if err := registry.DeleteServer(name); err != nil {
		fmt.Printf("Error deleting server: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Server '%s' deleted successfully\n", name)
}

func handleServersList() {
	servers, err := registry.ListServers()
	if err != nil {
		fmt.Printf("Error listing servers: %v\n", err)
		os.Exit(1)
	}

	if len(servers) == 0 {
		fmt.Println("No servers configured")
		return
	}

	fmt.Println("Configured servers:")
	for _, server := range servers {
		status := "enabled"
		if !server.Enabled {
			status = "disabled"
		}
		fmt.Printf("  • %s (%s) - %s\n", server.Name, server.URL, status)
	}
}

func printServersUsage() {
	fmt.Println("Usage: flaky-maky servers <subcommand> [options]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("  add -name <name> -url <url>  Add a new server")
	fmt.Println("  disable <name>               Disable a server")
	fmt.Println("  enable <name>                Enable a server")
	fmt.Println("  list                         List all configured servers")
}
