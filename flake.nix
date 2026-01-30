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
          version = "2026.01.30-a";

          src = self;

          vendorHash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=";

			go = pkgs.go_1_25_6;
          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A minimal to-do list program with a few amenities";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };

        apps.recall = {
          type = "app";
          program = "${self.packages.${system}.recall}/bin/recall";
        };
      });
}
