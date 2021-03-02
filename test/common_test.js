const truffleAssert = require('truffle-assertions');


contract('Common Testing', (accounts) => {
    beforeEach("Initialize", async () => {
        console.log('xxx1-1', await web3.eth.getAccounts())
        console.log('xxx1-2', web3.eth.accounts.wallet)
        console.log('xxx1-3', web3.utils.toWei('1', 'ether'))
        console.log('xxx2', process.env.NETWORK)
    });

    it("test common", async () => {
        console.log('xxx3', process.env.NETWORK)
    });

});
