const ERC20AtlasManager = artifacts.require("ERC20AtlasManager");
const MyERC20 = artifacts.require("MyERC20");
const truffleAssert = require('truffle-assertions');

let atlasManager;
let atlas20Token;

contract('ERC20AtlasManager Testing', (accounts) => {
  let admin = accounts[0];  // deploy all contract
  let user = accounts[1];   // user account
  let recipient = accounts[2];

  beforeEach("Initialize contract instance", async () => {
    atlasManager = await ERC20AtlasManager.new();

    // deploy atlas20Token
    const name = 'RP TOKEN';
    const symbol = 'RP';
    const decimals = 18;
    atlas20Token = await MyERC20.new(name, symbol, decimals);
  });

  it("test all", async () => {
    let hynTokenAddress = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE';

    console.log('account', accounts[0]);

    let balance = await web3.eth.getBalance(recipient, 'latest')
    console.log('balance begin:', web3.utils.fromWei(balance, 'ether'), balance)

    let lockHynAmount = web3.utils.toWei('1', 'ether');
    let tx = await atlasManager.lockHyn(lockHynAmount, recipient, {from: recipient, value: lockHynAmount});
    await truffleAssert.eventEmitted(tx, 'Locked', (ev) => {
      return ev.token === hynTokenAddress && ev.amount == lockHynAmount && ev.recipient === recipient && ev.sender === recipient;
    }, "Locked event should be emitted");

    balance = await web3.eth.getBalance(recipient, 'latest')
    console.log('balance after lock:', web3.utils.fromWei(balance, 'ether'))

    await atlas20Token.mint(recipient, web3.utils.toWei('100', 'ether'));
    let lockTokenAmount = web3.utils.toWei('1', 'ether');
    await atlas20Token.approve(atlasManager.address, lockTokenAmount, {from: recipient});
    tx = await atlasManager.lockToken(atlas20Token.address, lockTokenAmount, recipient, {from: recipient})
    await truffleAssert.eventEmitted(tx, 'Locked', (ev) => {
      return ev.token === atlas20Token.address && ev.amount == lockTokenAmount && ev.recipient === recipient && ev.sender === recipient;
    }, "Locked event should be emitted");

    let tokenBalance = await atlas20Token.balanceOf(recipient)
    console.log('tokenBalance', tokenBalance.toString())
    assert.equal(web3.utils.toWei('99', 'ether'), tokenBalance)


    // unlock
    const receiptId = '0x041a37cedac8dd5e898751523e2c6bef7e97f1f9abde27eb659702f72dd1fe02'
    let hynBalance = await web3.eth.getBalance(recipient)
    const expectBalance = web3.utils.toBN(hynBalance).add(web3.utils.toBN(lockHynAmount));
    tx = await atlasManager.unlockHyn(lockHynAmount, recipient, receiptId)
    await truffleAssert.eventEmitted(tx, 'Unlocked', (ev) => {
      return ev.token === hynTokenAddress && ev.amount == lockHynAmount && ev.recipient === recipient && ev.receiptId === receiptId;
    }, "Unlocked event should be emitted");

    // await hecoManager.mintToken(heco20Token.address, web3.utils.toWei('1', 'ether'), recipient, '0x041a37cedac8dd5e898751523e2c6bef7e97f1f9abde27eb659702f72dd1fe02')

    /// hyn token 0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
    // const addressHyn = '0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee'
    // const nameHyn = 'Hyperion Token';
    // const symbolHyn = 'HYN';
    // const decimalsHyn = 18;
    // truffleAssert.passes(
    //     hecoManager.addToken(tokenManagerAddr, addressHyn, nameHyn, symbolHyn, decimalsHyn)
    // );


    // await heco20Token.approve(hecoManager.address, web3.utils.toWei('1', 'ether'))
    // hecoManager.mint
  });

});
