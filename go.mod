module github.com/hyperion-hyn/hyn-bridge

go 1.14

require (
	github.com/ethereum/go-ethereum v1.9.24
	github.com/spf13/viper v1.7.1
)

replace github.com/hyperion-hyn/bls v0.0.6 => ./third_party/bls

replace github.com/ethereum/go-ethereum v1.9.24 => ../go-ethereum
