export HOST=0.0.0.0
export PORT=18888
export CHAINPOINT=xx
export IoTeXDIDPROXYADDRESS=yyy
export GASPRICE=1000
export GASLIMIT=10000
go build ./cmd/did-server
./did-server