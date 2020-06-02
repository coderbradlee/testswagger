export HOST=0.0.0.0
export PORT=18888
export CHAINPOINT=api.testnet.iotex.one:443
export IoTeXDIDPROXYADDRESS=io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0
go build ./cmd/did-server
./did-server