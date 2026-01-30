{
  description = "Recall devshell and package";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        # Your devshell stays the same
        devShells.default = pkgs.mkShell {
          name = "recall-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
          ];
        };

        packages.recall = pkgs.buildGoModule {
          pname = "recall";
          version = "0.1.0";

          src = self;

          vendorHash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "Gamified to-do list in Go";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };

        # Optional: expose recall as an app for nix run
        apps.recall = {
          type = "app";
          program = "${self.packages.${system}.recall}/bin/recall";
        };
      });
}
