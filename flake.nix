{
  description = "blockfont - Unicode block letter rendering for terminal applications";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go development
            go
            gopls
            golangci-lint
            gotools
            go-tools
            delve

            # Pre-commit and linting
            pre-commit

            # Documentation
            python312
            python312Packages.mkdocs
            python312Packages.mkdocs-material
            python312Packages.mkdocs-minify-plugin

            # Git and utilities
            git
            gh
          ];

          shellHook = ''
            echo "blockfont development environment"
            echo "Go version: $(go version)"
            echo ""
            echo "Available commands:"
            echo "  nix run .#build    - Build the library"
            echo "  nix run .#test     - Run tests"
            echo "  nix run .#lint     - Run linters"
            echo "  nix run .#docs     - Build documentation"
            echo "  nix run .#example  - Run example application"
          '';
        };

        # Convenience apps
        apps = {
          build = {
            type = "app";
            program = toString (pkgs.writeShellScript "build" ''
              cd ${toString ./.}
              ${pkgs.go}/bin/go build ./...
            '');
          };

          test = {
            type = "app";
            program = toString (pkgs.writeShellScript "test" ''
              cd ${toString ./.}
              ${pkgs.go}/bin/go test -v ./...
            '');
          };

          lint = {
            type = "app";
            program = toString (pkgs.writeShellScript "lint" ''
              cd ${toString ./.}
              ${pkgs.golangci-lint}/bin/golangci-lint run ./...
            '');
          };

          docs = {
            type = "app";
            program = toString (pkgs.writeShellScript "docs" ''
              cd ${toString ./.}/docs
              ${pkgs.python312Packages.mkdocs}/bin/mkdocs build
            '');
          };

          docs-serve = {
            type = "app";
            program = toString (pkgs.writeShellScript "docs-serve" ''
              cd ${toString ./.}/docs
              ${pkgs.python312Packages.mkdocs}/bin/mkdocs serve
            '');
          };
        };

        # Package
        packages.default = pkgs.buildGoModule {
          pname = "blockfont";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;  # Will be updated after go mod tidy
        };
      }
    );
}
