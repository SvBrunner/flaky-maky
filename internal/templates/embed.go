package templates

import "embed"

//go:embed flake-template.nix.tpl
var FlakeTemplate embed.FS
