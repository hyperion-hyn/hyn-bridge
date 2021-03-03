const ERC20AtlasManager = artifacts.require("ERC20AtlasManager");
const MyERC20 = artifacts.require("MyERC20");
const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");

const truffleAssert = require('truffle-assertions');

let tokenManager;
let hecoManager;
let atlasManager;
let rpToken;

const HYN_ADDRESS = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

contract('All in one testing', (accounts) => {
  let admin = accounts[0];  // deploy all contract
  let user = accounts[1];   // user account
  let recipient = accounts[2];

  beforeEach("Initialize", async () => {
    atlasManager = await ERC20AtlasManager.new();
    tokenManager = await TokenManager.new();
    hecoManager = await ERC20HecoManager.new();
  });

  it("test all", async () => {
    // rely
    await tokenManager.rely(hecoManager.address)

    // deploy rp token
    const name = 'RP TOKEN';
    const symbol = 'RP';
    const decimals = 18;
    rpToken = await MyERC20.new(name, symbol, decimals);
    // mint 100 rp-token for user/recipient
    await rpToken.mint(user, web3.utils.toWei('100', 'ether'));
    await rpToken.mint(recipient, web3.utils.toWei('100', 'ether'));

    // mapping tokens
    await hecoManager.addToken(tokenManager.address, rpToken.address, name, symbol, decimals)
    let hecoRpTokenAddr = await tokenManager.mappedTokens(rpToken.address);
    let hecoRpToken = await BridgedToken.at(hecoRpTokenAddr);

    // lock 9 rp-token by recipient
    let lockTokenAmount = web3.utils.toWei('9', 'ether');
    await rpToken.approve(atlasManager.address, lockTokenAmount, {from: recipient});
    tx = await atlasManager.lockToken(rpToken.address, lockTokenAmount, recipient, {from: recipient})
    console.log('xxx1', tx.receipt.transactionHash, tx.receipt.gasUsed)

    tx = await hecoManager.mintToken(hecoRpToken.address, lockTokenAmount, recipient, tx.receipt.transactionHash)
    console.log('xxx2', tx.receipt.transactionHash, tx.receipt.gasUsed)
    let hecoRpTokenBalance = await hecoRpToken.balanceOf(recipient)
    assert.equal(hecoRpTokenBalance.toString(), lockTokenAmount.toString())

    // burn and unlock
    let burnAmount = hecoRpTokenBalance;
    await hecoRpToken.approve(hecoManager.address, burnAmount, {from: recipient});
    tx = await hecoManager.burnToken(hecoRpToken.address, burnAmount, recipient, {from: recipient});
    console.log('xxx3', tx.receipt.transactionHash, tx.receipt.gasUsed)
    hecoRpTokenBalance = await hecoRpToken.balanceOf(recipient)
    assert.equal(hecoRpTokenBalance.toString(), '0')

    let rpTokenBalance = await rpToken.balanceOf(recipient)
    let expectBalance = web3.utils.toBN(rpTokenBalance).add(web3.utils.toBN(burnAmount));
    tx = await atlasManager.unlockToken(rpToken.address, burnAmount, recipient, tx.receipt.transactionHash)
    console.log('xxx4', tx.receipt.transactionHash, tx.receipt.gasUsed)
    rpTokenBalance = await rpToken.balanceOf(recipient)
    assert.equal(rpTokenBalance.toString(), expectBalance.toString())


    /// hyn
    await hecoManager.addToken(tokenManager.address, HYN_ADDRESS, 'Hyperion Token', 'HYN', 18)
    let hecoHynTokenAddr = await tokenManager.mappedTokens(HYN_ADDRESS);
    let hecoHynToken = await BridgedToken.at(hecoHynTokenAddr);

    let hynBalance = await web3.eth.getBalance(recipient, 'latest')
    let lockAmount = web3.utils.toWei('9', 'ether');
    tx = await atlasManager.lockHyn(lockAmount, recipient, {from: recipient, value: lockAmount, gasPrice: 1})
    console.log('yyy1', tx.receipt.transactionHash, tx.receipt.gasUsed)
    let expectedHynBalance = web3.utils.toBN(hynBalance).sub(web3.utils.toBN(lockAmount)).sub(web3.utils.toBN(tx.receipt.gasUsed))
    // expectedHynBalance = expectedHynBalance.sub(web3.utils.toBN(tx.receipt.gasUsed))
    hynBalance = await web3.eth.getBalance(recipient, 'latest')
    assert.equal(expectedHynBalance.toString(), hynBalance.toString())

    // mint in heco
    let hecoHynBalance = await hecoHynToken.balanceOf(recipient)
    let expectHecoHynBalance = web3.utils.toBN(hecoHynBalance).add(web3.utils.toBN(lockAmount));
    tx = await hecoManager.mintToken(hecoHynToken.address, lockAmount, recipient, tx.receipt.transactionHash)
    console.log('yyy2', tx.receipt.transactionHash, tx.receipt.gasUsed)
    hecoHynBalance = await hecoHynToken.balanceOf(recipient)
    assert.equal(hecoHynBalance.toString(), expectHecoHynBalance.toString())

    // burn in heco
    let burnHynAmount = hecoRpTokenBalance;
    expectHecoHynBalance = web3.utils.toBN(hecoHynBalance).sub(web3.utils.toBN(burnHynAmount));
    await hecoHynToken.approve(hecoManager.address, burnHynAmount, {from: recipient});
    tx = await hecoManager.burnToken(hecoHynToken.address, burnHynAmount, recipient, {from: recipient});
    console.log('yyy3', tx.receipt.transactionHash, tx.receipt.gasUsed)
    hecoRpTokenBalance = await hecoHynToken.balanceOf(recipient)
    assert.equal(hecoRpTokenBalance.toString(), expectHecoHynBalance.toString())

    hynBalance = await web3.eth.getBalance(recipient, 'latest')
    expectBalance = web3.utils.toBN(hynBalance).add(web3.utils.toBN(burnHynAmount));
    tx = await atlasManager.unlockHyn(burnHynAmount, recipient, tx.receipt.transactionHash)
    console.log('yyy4', tx.receipt.transactionHash, tx.receipt.gasUsed)
    hynBalance = await web3.eth.getBalance(recipient, 'latest')
    assert.equal(hynBalance.toString(), expectBalance.toString())

  });

});
