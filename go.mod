module github.com/pincheng0101/likecoin-chain-tx-exporter

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.11
	github.com/lib/pq v1.10.2 // indirect
	github.com/likecoin/likechain v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
