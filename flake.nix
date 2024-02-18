{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    devenv.url = "github:cachix/devenv";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = { self, nixpkgs, devenv, ... } @ inputs:
    let
      pkgs = nixpkgs.legacyPackages."x86_64-linux";
    in
    {
      devShell.x86_64-linux = devenv.lib.mkShell {
        inherit inputs pkgs;

        modules = [
          ({ pkgs, lib, ... }: {

            # This is your devenv configuration
            packages = with pkgs; [
              pkgs.buf
              pkgs.protobuf
              pkgs.protoc-gen-go
              pkgs.go-protobuf

              pkgs.tilt
              pkgs.go
              pkgs.gopls
              pkgs.gotests
              pkgs.gosec
              pkgs.govulncheck
              pkgs.kubernetes-helm
              pkgs.kind
              pkgs.ctlptl
              pkgs.golangci-lint
            ];

            enterShell = ''
            '';
          })
        ];
      };
    };
}
