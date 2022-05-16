with import <nixpkgs> {}; with goPackages;

buildGoPackage rec {
  name = "jobgosh";
  buildInputs = [ net osext ];
  goPackagePath = "github.com/jumang4423/jobgosh";
}