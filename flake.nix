{
  description = "Go dev environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            golangci-lint
            delve
            git
            just
          ];

          # optional, aber praktisch:
          env = {
            GOPATH = "/home/sven/go";
          };

          shellHook = ''
            export PATH="$GOPATH/bin:$PATH"
            echo "Go dev shell ready 🐹"
            go version
          '';
        };
      }
    );
}
