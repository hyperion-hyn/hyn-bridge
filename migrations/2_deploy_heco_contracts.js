const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");
const MultiSigWallet = artifacts.require("MultiSigWallet");

const HECO_MANAGER_JSON = require('../build/contracts/ERC20HecoManager.json')
const ATLAS_MANAGER_JSON = require('../build/contracts/ERC20AtlasManager.json')

let hecoManagerContract;
let atlasManagerContract;


module.exports = function (deployer, network, accounts) {
    let rpTokenAddress = '0x562D6AFA2A0aD94c8B2946e23C96E27F3cD023e8'
    if (network == 'heco') {
        rpTokenAddress = '0x88880126bC73107118f18000309d10dB9f1d6a14'
    }
    const HYN_ADDRESS = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
    hecoManagerContract = new web3.eth.Contract(HECO_MANAGER_JSON.abi)

    let tokenManager, hecoManager;
    deployer.then(async () => {
        // if(network == 'hecotest') {
        // }
        let committeeAddr1 = accounts[0];
        multiSigWalletHeco = await deployer.deploy(MultiSigWallet, [committeeAddr1], 1)

        tokenManager = await deployer.deploy(TokenManager);
        hecoManager = await deployer.deploy(ERC20HecoManager, multiSigWalletHeco.address);

        // rely
        console.log('begin rely', hecoManager.address)
        await tokenManager.rely(hecoManager.address)
        // await tokenManager.rely(accounts[0])
        console.log('rely done')

        //add hyn token
        const hynName = 'Hyperion Token';
        const hynSymbol = 'HYN';
        const hynDecimals = 18;
        let hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
        if (hecoHynTokenAddr == '0x0000000000000000000000000000000000000000') {
            // console.log('deploy BridgedToken', HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
            // let hecoHynToken = await deployer.deploy(BridgedToken, HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
            // console.log('add minter', hecoManager.address)
            // await hecoHynToken.addMinter(hecoManager.address)
            // console.log('registerToken', HYN_ADDRESS, hecoHynToken.address)
            // await tokenManager.registerToken(HYN_ADDRESS, hecoHynToken.address)

            console.log('begin addHynToken', tokenManager.address, HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
            let abi = hecoManagerContract.methods.addToken(tokenManager.address, HYN_ADDRESS, hynName, hynSymbol, hynDecimals).encodeABI()
            await multiSigWalletHeco.submitTransaction(hecoManager.address, 0, abi, {from: committeeAddr1});
            // await hecoManager.addToken(tokenManager.address, HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
            hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
            console.log('addHynToken done', hecoHynTokenAddr)
        }
        console.log('hecoHynTokenAddr', hecoHynTokenAddr)

        // add rp token
        const rpName = 'RP Token';
        const rpSymbol = 'RP';
        const rpDecimals = 18;
        let hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
        if (hecoRpTokenAddr == '0x0000000000000000000000000000000000000000') {
            // console.log('deploy BridgedToken', rpTokenAddress, rpName, rpSymbol, rpDecimals)
            // let hecoRpToken = await deployer.deploy(BridgedToken, rpTokenAddress, rpName, rpSymbol, rpDecimals)
            // console.log('add minter', hecoManager.address)
            // await hecoRpToken.addMinter(hecoManager.address)
            // console.log('registerToken', rpTokenAddress, hecoRpToken.address)
            // await tokenManager.registerToken(rpTokenAddress, hecoRpToken.address)

            console.log('begin addToken', tokenManager.address, rpTokenAddress, rpName, rpSymbol, rpDecimals)
            let abi = hecoManagerContract.methods.addToken(tokenManager.address, rpTokenAddress, rpName, rpSymbol, rpDecimals).encodeABI()
            await multiSigWalletHeco.submitTransaction(hecoManager.address, 0, abi, {from: committeeAddr1});
            // await hecoManager.addToken(tokenManager.address, rpTokenAddress, rpName, rpSymbol, rpDecimals)

            hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
            console.log('addRpToken done', hecoRpTokenAddr)
        }
        console.log('hecoRpTokenAddr', hecoRpTokenAddr)
    });
};
