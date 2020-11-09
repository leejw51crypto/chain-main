with import <nixpkgs> { };
mkShell {
  buildInputs = [
    python38
    protobuf3_13
    grpc
    git
  ];
  shellHook = ''                                                                                       
  ./convert.sh                                                                                          '';
}
