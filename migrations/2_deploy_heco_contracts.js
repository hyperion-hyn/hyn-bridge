const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");

// roles:
// 1. admin: deploy TokenManager,
// 2. heco
//      a. master: deploy hecoManager
//      b. hecoUser:
// 3. atlas
//      a. master:   deploy atlasManager
//      b. atlasUser:
module.exports = function (deployer, network, accounts) {
  let rpTokenAddress = '0x562D6AFA2A0aD94c8B2946e23C96E27F3cD023e8'
  if(network == 'heco') {
    rpTokenAddress = '0x88880126bC73107118f18000309d10dB9f1d6a14'
  }
  const HYN_ADDRESS = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
  let tokenManager, hecoManager;
  deployer.then(async () => {
    // if(network == 'hecotest') {
    //
    // }
    await deployer.deploy(TokenManager);
    tokenManager = await TokenManager.deployed()

    await deployer.deploy(ERC20HecoManager);
    hecoManager = await ERC20HecoManager.deployed()

    // rely
    await tokenManager.rely(hecoManager.address)

    //add hyn token
    const hynName = 'Hyperion Token';
    const hynSymbol = 'HYN';
    const hynDecimals = 18;
    let hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
    if(hecoHynTokenAddr == '0x0000000000000000000000000000000000000000') {
      await hecoManager.addToken(tokenManager.address, HYN_ADDRESS, hynName, hynSymbol, hynDecimals)
      hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS)
    }
    console.log('hecoHynTokenAddr', hecoHynTokenAddr)

    // add rp token
    const rpName = 'RP Token';
    const rpSymbol = 'RP';
    const rpDecimals = 18;
    let hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
    if(hecoRpTokenAddr == '0x0000000000000000000000000000000000000000') {
      await hecoManager.addToken(tokenManager.address, rpTokenAddress, rpName, rpSymbol, rpDecimals)
      hecoRpTokenAddr = await tokenManager.mappedTokens(rpTokenAddress)
    }
    console.log('hecoRpTokenAddr', hecoRpTokenAddr)
  });
};
