# 🎯 Flaky Maky

An interactive terminal tool for easy creation and management of NixOS flake configurations for development projects.

## 📋 Overview

Flaky Maky is a user-friendly command-line tool that helps NixOS users generate development flakes for their projects. With an interactive wizard and pre-configured templates, creating flakes becomes a breeze.

The preconfigured templates can be self hosted via [flaky-servy](https://github.com/SvBrunner/flaky-servy)

## ✨ Features

- **🎨 Interactive Wizard**: Step-by-step guidance for flake creation
- **📦 Pre-configured Templates**: Access to ready-made configurations from registry servers
- **🔧 Server Management**: Add, enable, disable, and delete configuration servers
- **🔄 Configuration Sync**: Download and synchronize pre-configurations
- **🌐 direnv Support**: Optional integration for automatic development environments
- **🖥️ Multi-Architecture**: Support for various NixOS system architectures

## 📖 Usage

### Getting Started

**Important:** Before first use, configurations must be synchronized:

```bash
flaky-maky --sync
```

### Create a New Flake

Start the interactive wizard:

```bash
flaky-maky
```

The wizard will guide you through the following steps:

1. **Name**: Assign a name to your flake
2. **NixOS Channel**: Select a NixOS channel (e.g., nixos-unstable, nixos-24.05)
3. **System Architecture**: Choose your target architecture (x86_64-linux, aarch64-darwin, etc.)
4. **Template**: Select a pre-configuration (Go, Typst, etc.)
5. **direnv**: Optionally enable direnv integration

### Server Management

#### Add a Server

```bash
flaky-maky servers add -name <name> -url <url>
```

**Example:**
```bash
flaky-maky servers add -name my-server -url https://configs.example.com
```

#### List All Servers

```bash
flaky-maky servers list
```

**Output:**
```
Configured servers:
  • my-server (https://configs.example.com) - enabled
  • backup-server (https://backup.example.com) - disabled
```

#### Enable/Disable a Server

```bash
flaky-maky servers enable my-server

flaky-maky servers disable my-server
```

#### Delete a Server

```bash
flaky-maky servers delete my-server
```

### Synchronize Configurations

Download the latest pre-configurations from all enabled registry servers:

```bash
flaky-maky --sync
```

This should be run regularly to get the latest templates.

## 🛠️ Development

### Prerequisites

- Go 1.25.7 or higher

### Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI Framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI Components

- [yaml.v3](https://gopkg.in/yaml.v3) - YAML Parser

### Build the Project

```bash
go build -o flaky-maky
```

### Run Tests

```bash
go test ./...
```

## 🤝 Contributing

Contributions are welcome! Feel free to create issues or submit pull requests.

### Development Notes

This project also serves as a practice project for Go development. Feedback on code quality and best practices is explicitly welcome!

Parts of the project (Like this README) are generated with AI.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


**Note:** This tool is under active development. Features and API may still change.
