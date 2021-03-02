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
module.exports = function(deployer) {
  // deployer.deploy(ConvertLib);
  // deployer.link(ConvertLib, MetaCoin);
  // deployer.deploy(MetaCoin);

};
