pragma solidity 0.5.17;
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";
import "./lib/TokenManager.sol";

interface MintableToken {
    function mint(address beneficiary, uint256 amount) external returns (bool);
}

interface BurnableToken {
    function burnFrom(address account, uint256 amount) external;
}

contract ERC20HecoManager is Ownable {
    mapping(bytes32 => bool) public usedEvents_;
    mapping(address => address) public mappings;

    event Burned(
        address indexed token,
        address indexed sender,
        uint256 amount,
        address recipient
    );

    event Minted(
        address token,
        uint256 amount,
        address recipient,
        bytes32 receiptId
    );

//    address public wallet;
//    modifier onlyWallet {
//        require(msg.sender == wallet, "HecoManager/not-authorized");
//        _;
//    }
//
//    /**
//     * @dev constructor
//     * @param _wallet is the multisig wallet
//     */
//    constructor(address _wallet) public {
//        wallet = _wallet;
//    }

    /**
     * @dev map an atlas token to heco
     * @param tokenManager address to token manager
     * @param hynTokenAddr atlas token address to map
     * @param name of the atlas token
     * @param symbol of the atlas token
     * @param decimals of the atlas token
     */
    function addToken(
        address tokenManager,
        address hynTokenAddr,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) public onlyOwner {
        address heco20TokenAddr = TokenManager(tokenManager).addToken(
            hynTokenAddr,
            name,
            symbol,
            decimals
        );
        mappings[hynTokenAddr] = heco20TokenAddr;
    }

    /**
     * @dev deregister token mapping in the token manager
     * @param tokenManager address to token manager
     * @param hynTokenAddr address to remove token
     */
    function removeToken(address tokenManager, address hynTokenAddr) public onlyOwner {
        TokenManager(tokenManager).removeToken(hynTokenAddr, 0);
    }

    /**
     * @dev burns tokens on harmony to be unlocked on ethereum
     * @param heco20Token heco token address
     * @param amount amount of tokens to burn
     * @param recipient recipient of the unlock tokens on atlas
     */
    function burnToken(
        address heco20Token,
        uint256 amount,
        address recipient
    ) public {
        BurnableToken(heco20Token).burnFrom(msg.sender, amount);
        emit Burned(heco20Token, msg.sender, amount, recipient);
    }

    /**
     * @dev mints tokens corresponding to the tokens locked in the ethereum chain
     * @param heco20Token is the token address for minting
     * @param amount amount of tokens for minting
     * @param recipient recipient of the minted tokens (harmony address)
     * @param receiptId transaction hash of the lock event on ethereum chain
     */
    function mintToken(
        address heco20Token,
        uint256 amount,
        address recipient,
        bytes32 receiptId
    ) public onlyOwner {
        require(
            !usedEvents_[receiptId],
            "HecoManager/The lock event cannot be reused"
        );
        usedEvents_[receiptId] = true;
        MintableToken(heco20Token).mint(recipient, amount);
        emit Minted(heco20Token, amount, recipient, receiptId);
    }
}
