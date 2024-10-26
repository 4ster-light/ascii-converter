{
  description = "Nix flake to build and run a project using a Makefile";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system: 
      let 
        pkgs = import nixpkgs { inherit system; };
      in {
        devShell = pkgs.mkShell {
          name = "ascii-converter-dev";

          buildInputs = [
            pkgs.bun
            pkgs.go_1_23
            pkgs.gnumake
          ];

          shellHook = ''
            echo "Nix shell ready. You can now use 'make build', 'make run', etc."
          '';
        };
      }
    );
}
