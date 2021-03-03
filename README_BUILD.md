# build
git submodule add https://github.com/hyperion-hyn/bls third_party/bls
git submodule add https://github.com/hyperion-hyn/mcl third_party/mcl

make setup
make third_party

npm install

# deploy
