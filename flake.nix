{
  description = "doggl";

  inputs = { nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable"; };

  outputs = { self, nixpkgs, ... }:
    let
      supportedSystems = [ "aarch64-darwin" "x86_64-darwin" "x86_64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
          v = builtins.replaceStrings [ "\n" "\r" ] [ "" "" ]
            (builtins.readFile ./version);
        in rec {
          doggl = pkgs.buildGo124Module {
            pname = "doggl";
            version = v;
            src = ./.;
            vendorHash = "sha256-+Ezs6+YOOIESXrQneAQAsfvo3L6LwIiBx3LEybgEqBw=";
            ldflags = [ "-X=github.com/tubenhirn/doggl/cmd.AppVersion=${v}" ];
          };

          goreleaser-build = pkgs.stdenv.mkDerivation {
            pname = "doggl-goreleaser";
            version = v;
            src = ./.;
            buildInputs = [ pkgs.goreleaser pkgs.git pkgs.go_1_24 ];
            buildPhase = ''
              export HOME=$TMPDIR
              export GORELEASER_CURRENT_TAG="v${v}"
              export APP_VERSION="${v}"
              goreleaser --clean --config .goreleaser.yaml
            '';
            installPhase = ''
              mkdir -p $out
              cp -r dist/* $out/
            '';
          };

          default = doggl;
        });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
          v = builtins.replaceStrings [ "\n" "\r" ] [ "" "" ]
            (builtins.readFile ./version);
        in {
          default = pkgs.mkShell {
            buildInputs = [
              pkgs.go_1_24
              pkgs.goreleaser
              pkgs.git
              pkgs.direnv
              pkgs.zsh
              pkgs.semantic-release
              pkgs.yarn
            ];
            env = {
              GOFLAGS = "-mod=mod";
              APP_VERSION = "${v}";
            };
            shellHook = ''
              if [ -z "$IN_NIX_SHELL_ZSH" ] && [ -z "$ZSH_VERSION" ]; then
                 export IN_NIX_SHELL_ZSH=1
                 exec zsh
              fi
              echo "🐶 Welcome to the doggl development shell!"
              echo "If you use 1password, run: export GITHUB_TOKEN=\$(op read op://yourvault/github/GITHUB_TOKEN)"
            '';
          };
        });
    };
}
