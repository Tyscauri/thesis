# thesis
Permanent repository for the masterthesis "How Distributed Ledger Technology Facilitates The Establishment Of A Circular Economy"

## Scope:
Most of the repository contains required software from the Iota repository https://github.com/iotaledger/wasp - "develop" branch
My own work consists of 
thesis/wasp/contracts/wasm/circularity_artifact/schema.yaml
thesis/wasp/contracts/wasm/circularity_artifact/src/circularity_artifact.rs
thesis/wasp/contracts/wasm/circularity_artifact/test/circularity_artifact_test.go

compiling to the files
thesis/wasp/contracts/wasm/circularity_artifact/pkg

The Iota Schema Tool supported the development of the artifact. Further information regarding the schema tool could be found under https://wiki.iota.org/smart-contracts/guide/schema/usage

## Artifact deployment:
### Requirements:
Go 1.16
RocksDB
Linux Ubuntu (used Ubuntu 20.04.4 LTS, 64-bit)

- Clone the repository
- Move to thesis/wasp directory in the terminal
- "make install" (not tested)
- add wasp-cli, wasp and schema to the PATH variable
- move to the test folder directory in the terminal
- "go test" runs the test cases on the local machine
