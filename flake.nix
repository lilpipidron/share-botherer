{
  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixos-unstable";
    };
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };
  outputs = {
    self,
    nixpkgs,
    gomod2nix,
    ...
  }: let
    version = "0.0.1";
    supportedSystems = [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ];
    forEachSupportedSystem = f:
      nixpkgs.lib.genAttrs supportedSystems (system:
        f {
          pkgs = import nixpkgs {
            inherit system;
            overlays = [self.overlays.default];
          };
        });
  in {
    overlays.default = final: prev: {
      buildGoApplication = gomod2nix.legacyPackages.${prev.stdenv.system}.buildGoApplication;
      gomod2nixPkg = gomod2nix.packages.${prev.stdenv.system}.default;
    };
    devShells = forEachSupportedSystem ({pkgs}: {
      default = pkgs.mkShell {
        packages = with pkgs; [
          go
          gotools
          golangci-lint
          gomod2nixPkg
          alejandra
        ];
      };
    });
    packages = forEachSupportedSystem ({pkgs, ...}: rec {
      share-botherer = pkgs.buildGoApplication {
        pname = "share-botherer";
        inherit version;
        src = self;
        modules = ./gomod2nix.toml;
      };
      default = share-botherer;
    });
  };
}
