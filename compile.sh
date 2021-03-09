truffle compile --all

mkdir -p build/abi
mkdir -p build/bin

cat ./build/contracts/ERC20AtlasManager.json | jq '.abi' > ./build/abi/ERC20AtlasManager.abi
cat ./build/contracts/ERC20AtlasManager.json | jq -r '.bytecode' > ./build/bin/ERC20AtlasManager.bin
abigen --bin=./build/bin/ERC20AtlasManager.bin --abi=./build/abi/ERC20AtlasManager.abi --type=ERC20AtlasManager --pkg=main --out=./deploy/auto_gen_erc20_atlas_manager.go

cat ./build/contracts/MultiSigWallet.json | jq '.abi' > ./build/abi/MultiSigWallet.abi
cat ./build/contracts/MultiSigWallet.json | jq -r '.bytecode' > ./build/bin/MultiSigWallet.bin
abigen --bin=./build/bin/MultiSigWallet.bin --abi=./build/abi/MultiSigWallet.abi --type=MultiSigWallet --pkg=main --out=./deploy/auto_gen_multi_sig_wallet.go
/Users/moo/project/github/go-ethereum/build/bin/abigen --bin=./build/bin/MultiSigWallet.bin --abi=./build/abi/MultiSigWallet.abi --type=MultiSigWallet --pkg=heco --out=./deploy/heco/auto_gen_multi_sig_wallet.go


cat ./build/contracts/ERC20HecoManager.json | jq '.abi' > ./build/abi/ERC20HecoManager.abi
cat ./build/contracts/ERC20HecoManager.json | jq -r '.bytecode' > ./build/bin/ERC20HecoManager.bin
/Users/moo/project/github/go-ethereum/build/bin/abigen --bin=./build/bin/ERC20HecoManager.bin --abi=./build/abi/ERC20HecoManager.abi --type=ERC20HecoManager --pkg=heco --out=./deploy/heco/auto_gen_erc20_heco_manager.go

cat ./build/contracts/TokenManager.json | jq '.abi' > ./build/abi/TokenManager.abi
cat ./build/contracts/TokenManager.json | jq -r '.bytecode' > ./build/bin/TokenManager.bin
/Users/moo/project/github/go-ethereum/build/bin/abigen --bin=./build/bin/TokenManager.bin --abi=./build/abi/TokenManager.abi --type=TokenManager --pkg=heco --out=./deploy/heco/auto_gen_token_manager.go

