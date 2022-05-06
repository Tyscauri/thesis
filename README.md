# Thesis
This is the permanent repository corresponding to the masterthesis "How Distributed Ledger Technology Facilitates the Establishment of a Circular Economy"

## Scope:
Most of the repository contains required software from the Iota repository https://github.com/iotaledger/wasp - "develop" branch  
My own work consists of  
thesis/wasp/contracts/wasm/circularity_artifact/schema.yaml  
thesis/wasp/contracts/wasm/circularity_artifact/src/circularity_artifact.rs  
thesis/wasp/contracts/wasm/circularity_artifact/test/circularity_artifact_test.go  
  

The Iota Schema Tool supported the development of the artifact. Further information regarding the Schema Tool can be found under  https://wiki.iota.org/smart-contracts/guide/schema/usage  
More information provided regarding the Iota node and smart contracts can be found here:  https://wiki.iota.org/smart-contracts/guide/chains_and_nodes/running-a-node

## Artifact deployment:
### Requirements:
- Go 1.16
- RocksDB
- Linux Ubuntu (used Ubuntu 20.04.4 LTS, 64-bit)

### Steps to deploy the artifact and run the testcases
- Clone the repository
- Move to the directory in the terminal under thesis/wasp
- "make install" (not tested)
- add wasp-cli, wasp and schema to the PATH variable
- move to the test folder directory in the terminal under thesis/wasp/contracts/wasm/circularity_artifact/test/
- "go test" runs the test cases on the local machine

For recompiling
- move to thesis/wasp/contracts/wasm/circularity_artifact/
- "schema -rust -go"
- which compiles to the files in  
thesis/wasp/contracts/wasm/circularity_artifact/pkg  

