package templates

import "embed"

//go:embed templates/*
var DefaultTemplates embed.FS

//go:embed flake-template.nix.tpl
var FlakeTemplate embed.FS
