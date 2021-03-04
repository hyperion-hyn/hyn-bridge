const truffleAssert = require('truffle-assertions');

const ERC20AtlasManager = artifacts.require("ERC20AtlasManager");
const MyERC20 = artifacts.require("MyERC20");
const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");

contract('Common Testing', (accounts) => {
    it("test common", async () => {
        let token = await BridgedToken.at('0x1990f4c2D9cbB7587e1864812d0403e52fa32f03')
        // let tx = await token.mint('0x89A9855032047fAF65BAA95F43128af6EE5721eD', web3.utils.toWei('1', 'ether'))
        let name = await token.name();
        console.log(name)
    });

});
