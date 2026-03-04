{
  description = "{{.Description}}";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "{{.Channel}}";
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
