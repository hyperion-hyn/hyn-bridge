const TokenManager = artifacts.require("TokenManager");
const truffleAssert = require('truffle-assertions');

let tokenManagerInstance;

contract('TokenManager Testing', (accounts) => {
    beforeEach("Initialize TokenManager contract instance", async () => {
        tokenManagerInstance = await TokenManager.new();
    });

    it("test all", async () => {
        const userAddToken = accounts[1];

        const rpTokenAddress = '0x562D6AFA2A0aD94c8B2946e23C96E27F3cD023e8';
        const name = 'RP TOKEN';
        const symbol = 'RP';
        const decimals = 18;
        await truffleAssert.fails(
            tokenManagerInstance.addToken(rpTokenAddress, name, symbol, decimals, {from: userAddToken}),
            truffleAssert.ErrorType.REVERT,
            'TokenManager/not-authorized'
        )

        let erc20Token = '';

        await tokenManagerInstance.rely(userAddToken)
        let tx = await tokenManagerInstance.addToken(rpTokenAddress, name, symbol, decimals, {from: userAddToken})
        await truffleAssert.eventEmitted(tx, 'TokenMapAck', (ev) => {
            erc20Token = ev.tokenAck;
            return ev.tokenReq === rpTokenAddress;
        }, "TokenMapAck should be emitted");

        await truffleAssert.fails(
            tokenManagerInstance.addToken(rpTokenAddress, name, symbol, decimals, {from: userAddToken}),
            truffleAssert.ErrorType.REVERT,
            'TokenManager/ethToken already mapped'
        )

        let hecoToken = await tokenManagerInstance.mappedTokens(rpTokenAddress);
        assert.equal(hecoToken, erc20Token);

        await truffleAssert.passes(
            tokenManagerInstance.removeToken(rpTokenAddress, 0, {from: userAddToken})
        );

        hynToken = await tokenManagerInstance.mappedTokens(erc20Token);
        assert.equal('0x0000000000000000000000000000000000000000', hynToken);

        // deny
        await  tokenManagerInstance.deny(userAddToken)
        await truffleAssert.fails(
            tokenManagerInstance.addToken(rpTokenAddress, name, symbol, decimals, {from: userAddToken}),
            truffleAssert.ErrorType.REVERT,
            'TokenManager/not-authorized'
        )

    });

});
