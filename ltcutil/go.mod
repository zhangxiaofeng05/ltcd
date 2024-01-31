module github.com/ltcsuite/ltcd/ltcutil

go 1.16

require (
	github.com/aead/siphash v1.0.1
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/kkdai/bstream v1.0.0
	github.com/ltcsuite/ltcd v0.23.4
	github.com/ltcsuite/ltcd/btcec/v2 v2.1.3
	github.com/ltcsuite/ltcd/chaincfg/chainhash v1.0.2
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

replace github.com/ltcsuite/ltcd => ../
replace github.com/ltcsuite/ltcd/btcec/v2 => ../btcec
replace github.com/ltcsuite/ltcd/chaincfg/chainhash => ../chaincfg/chainhash