const ERC20HecoManager = artifacts.require("ERC20HecoManager");
const TokenManager = artifacts.require("TokenManager");
const BridgedToken = artifacts.require("BridgedToken");
const MyERC20 = artifacts.require("MyERC20");
const truffleAssert = require('truffle-assertions');

let tokenManager;
let hecoManager;
let atlas20Token;

contract('ERC20HecoManager Testing', (accounts) => {
    let admin = accounts[0];  // deploy all contract
    let user = accounts[1];   // user account
    let recipient = accounts[2];   // user account

    beforeEach("Initialize TokenManager contract instance", async () => {
        tokenManager = await TokenManager.new();
        hecoManager = await ERC20HecoManager.new();

        // rely
        await tokenManager.rely(hecoManager.address)

        // deploy atlas20Token
        const name = 'RP TOKEN';
        const symbol = 'RP';
        const decimals = 18;
        atlas20Token = await MyERC20.new(name, symbol, decimals);
    });

    it("test all", async () => {
        // mint user some token
        await atlas20Token.mint(user, web3.utils.toWei('100', 'ether'));
        assert.equal(await atlas20Token.balanceOf(user), web3.utils.toWei('100', 'ether'));
        assert.equal(await atlas20Token.totalSupply(), web3.utils.toWei('100', 'ether'));

        const tokenManagerAddr = tokenManager.address;

        const rpTokenAddress = atlas20Token.address;
        const name = await atlas20Token.name();
        const symbol = await atlas20Token.symbol();
        const decimals = await atlas20Token.decimals();

        /// rp token
        await truffleAssert.passes(
            hecoManager.addToken(tokenManagerAddr, rpTokenAddress, name, symbol, decimals)
        );

        const heco20TokenAddr = await tokenManager.mappedTokens(rpTokenAddress);
        const heco20Token = await BridgedToken.at(heco20TokenAddr);

        const receiptId = '0x041a37cedac8dd5e898751523e2c6bef7e97f1f9abde27eb659702f72dd1fe02'
        let tx = await hecoManager.mintToken(heco20Token.address, web3.utils.toWei('1', 'ether'), recipient, receiptId)
        await truffleAssert.eventEmitted(tx, 'Minted', (ev) => {
            return ev.token === heco20Token.address && ev.amount == web3.utils.toWei('1', 'ether') && ev.recipient === recipient && ev.receiptId === receiptId;
        }, "Minted event should be emitted");

        await truffleAssert.fails(
            hecoManager.mintToken(heco20Token.address, web3.utils.toWei('1', 'ether'), recipient, receiptId),
            truffleAssert.ErrorType.REVERT,
            'HecoManager/The lock event cannot be reused'
        )

        assert.equal(await heco20Token.balanceOf(recipient), web3.utils.toWei('1', 'ether'));

        // fail, should approve first
        await truffleAssert.fails(
            hecoManager.burnToken(heco20Token.address, web3.utils.toWei('1', 'ether'), recipient, {from: recipient}),
            truffleAssert.ErrorType.REVERT,
            'ERC20: burn amount exceeds allowance'
        )

        await heco20Token.approve(hecoManager.address, web3.utils.toWei('1', 'ether'), {from: recipient})
        tx = await hecoManager.burnToken(heco20Token.address, web3.utils.toWei('1', 'ether'), recipient, {from: recipient})
        await truffleAssert.eventEmitted(tx, 'Burned', (ev) => {
            return ev.token === heco20Token.address && ev.sender === recipient && ev.amount == web3.utils.toWei('1', 'ether') && ev.recipient === recipient;
        }, "Burned event should be emitted");

        assert.equal(await heco20Token.balanceOf(recipient), web3.utils.toWei('0', 'ether'));

        await hecoManager.removeToken(tokenManagerAddr, rpTokenAddress);
        let emptyAddress = await tokenManager.mappedTokens(rpTokenAddress);
        assert.equal(emptyAddress, '0x0000000000000000000000000000000000000000')

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
