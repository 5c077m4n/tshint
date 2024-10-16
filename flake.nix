{
  description = "Go dev flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default =
          with pkgs;
          mkShell {
            buildInputs = [
              # Go deps
              go
              gotools # Go tools like goimports, godoc, and others
              golangci-lint
            ];
            packages = [
              # Dev Env
              git
              neovim
              fd
              fzf
              ripgrep
              fish
              starship
              bat
            ];

            shellHook = ''
              exec ${fish.outPath}/bin/fish
            '';
          };
      }
    );
}
