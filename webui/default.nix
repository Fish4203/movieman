# save this as shell.nix
{ pkgs ? import <nixpkgs> {}}:

let
  pkgs = import (builtins.fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/archive/e89cf1c932006531f454de7d652163a9a5c86668.tar.gz"; 
  }) {};
  
  nodejs = pkgs.elmPackages.nodejs;  
in
pkgs.mkShell {
  inherit nodejs;
  packages = [ 
    pkgs.hello
    pkgs.bun
    nodejs
  ];
}

