{
  description = "{{.Description}}";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    { flake-parts, ... }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [{{range .Systems}}
        "{{.}}"{{end}}
      ];

      perSystem =
        { pkgs, system, ... }:
        {
          devShells.default = pkgs.mkShell {
            packages = with pkgs; [{{range .Packages}}
              {{.}}{{end}}
            ];

            env = { {{range.EnvironmentVariables}}
              {{.Name}} = "{{.Value}}"{{end}}
            };

            shellHook = ''{{range .ShellHooks}}
              {{.}}{{end}}
            '';
          };
        };
    };
}
