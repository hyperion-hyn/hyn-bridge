const ERC20AtlasManager = artifacts.require("ERC20AtlasManager");
const MyERC20 = artifacts.require("MyERC20");
const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");

module.exports = async function (callback) {
    let hecoManagerAddr = ''
    let hynTokenAddress = ''
    let rpTokenAddress = ''
    let tokenManagerAddr = ''

    const networkType = process.argv[process.argv.length - 1];
    if (networkType === 'development') {
        hecoManagerAddr = '0x4F39a925CF5781451BA3D640bFd333d7720a7555'   // dev address
        hynTokenAddress = '0x55cc4e6d6e13ce403ce617e0240bf4405aa4d972'     // dev token address
        rpTokenAddress = ''
        tokenManagerAddr = '0x419Fd24f4C0044faB8f472EED64619679023b347' //dev
    } else if (networkType === 'hecotest') {
        hecoManagerAddr = '0xdFc01fE7Bac53B4d72518d5ad0e1131b83628fF3'
        hynTokenAddress = '0x1990f4c2D9cbB7587e1864812d0403e52fa32f03'  // heco test token
        rpTokenAddress = '0x786710cc6A52F4c07F5A2eCb2f03BCd00CE1d102'
        tokenManagerAddr = '0xf50005d5F49b25CA032F95A56A38B1cB634C69fF' //heco test
    }

    let tokenAddress = rpTokenAddress;

    try {
        console.log('xxxx1');

        // let abi = web3.eth.abi.encodeParameters(['address','string', 'string', 'uint8'], ['0x562D6AFA2A0aD94c8B2946e23C96E27F3cD023e8', 'RP TOKEN', 'RP', 18]);
        // let abi = web3.eth.abi.encodeParameters(['address','string', 'string', 'uint8'], ['0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE', 'Hyperion Token', 'HYN', 18]);
        // console.log(abi)
        // callback();
        // return;

        let hecoManager = await ERC20HecoManager.at(hecoManagerAddr)

        let recipient = '0x9AE56FD9862CbF6188b495D8eE6A45EE261E5306'
        let receiptId = '0x6d71a1bfc43e477124663e48de191930630df6c1eb76c6095ea1b86bbb65dbd3'
        let mintAmount = web3.utils.toWei('300', 'ether')
        console.log('xxxx2');

        let token = await BridgedToken.at(tokenAddress)
        let tokenBanaceOfRecipient = await token.balanceOf(recipient);
        console.log('xxxx3', web3.utils.fromWei(tokenBanaceOfRecipient, 'ether'));

        let tx = await hecoManager.mintToken(token.address, mintAmount, recipient, receiptId)
        console.log(tx)

        tokenBanaceOfRecipient = await token.balanceOf(recipient);
        console.log('xxxx3-2', web3.utils.fromWei(tokenBanaceOfRecipient, 'ether'));

        console.log('xxxx4');
    } catch (e) {
        console.error(e)
    }
    callback();
}
