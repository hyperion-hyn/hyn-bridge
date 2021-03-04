const MyERC20 = artifacts.require("MyERC20");
const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");

const truffleAssert = require('truffle-assertions');

let tokenManager;
let hecoManager;

// for test
let tokenManagerAddress = '0xc6526A3B7E35a1ACDdAA85E5879AD301fa46724e'
let hecoManagerAddress = '0x33bB65f976659AB3b7142812c158aEd404c73611'
let rpTokenAddress = '0x562D6AFA2A0aD94c8B2946e23C96E27F3cD023e8'

// mainnet
// let tokenManagerAddress = ''
// let hecoManagerAddress = ''
// let rpTokenAddress = '0x88880126bC73107118f18000309d10dB9f1d6a14'

const HYN_ADDRESS = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

contract('Init contract', (accounts) => {
    let admin = accounts[0];  // deploy all contract

    beforeEach("Initialize", async () => {
        tokenManager = await TokenManager.at(tokenManagerAddress)
        hecoManager = await ERC20HecoManager.at(hecoManagerAddress)
        // rpToken = await MyERC20.at(rpTokenAddress)
    });

    it("test all", async () => {
        // rely
        await tokenManager.rely(hecoManager.address)

        //add hyn token
        const hynName = 'Hyperion Token';
        const hynSymbol = 'HYN';
        const hynDecimals = 18;
        let hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
        if (hecoHynTokenAddr == '0x0000000000000000000000000000000000000000') {
            await hecoManager.addToken(tokenManager.address, HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
            hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
        }
        console.log('hecoHynTokenAddr', hecoHynTokenAddr)

        const rpName = 'RP Token';
        const rpSymbol = 'RP';
        const rpDecimals = 18;
        let hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
        if (hecoRpTokenAddr == '0x0000000000000000000000000000000000000000') {
            await hecoManager.addToken(tokenManager.address, rpTokenAddress, rpName, rpSymbol, rpDecimals)
            hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
        }
        console.log('hecoRpTokenAddr', hecoRpTokenAddr)

    });

});
