pragma solidity 0.5.17;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";

contract ERC20AtlasManager is Ownable {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    IERC20 public constant HYN_ADDRESS = IERC20(0x00eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee);

    mapping(bytes32 => bool) public usedEvents_;

    event Locked(
        address indexed token,
        address indexed sender,
        uint256 amount,
        address recipient
    );

    event Unlocked(
        address token,
        uint256 amount,
        address recipient,
        bytes32 receiptId
    );

    address public wallet;
    modifier onlyWallet {
        require(msg.sender == wallet, "AtlasManager/not-authorized");
        _;
    }

    /**
     * @dev constructor
     * @param _wallet is the multisig wallet
     */
    constructor(address _wallet) public {
        wallet = _wallet;
    }

    /**
     * @dev lock tokens to be minted on heco chain
     * @param hynTokenAddr is the atlas token contract address
     * @param amount amount of tokens to lock
     * @param recipient recipient address on the heco chain
     */
    function lockToken(
        address hynTokenAddr,
        uint256 amount,
        address recipient
    ) public {
        require(
            recipient != address(0),
            "AtlasManager/recipient is a zero address"
        );
        require(amount > 0, "AtlasManager/zero token locked");
        IERC20 hynToken = IERC20(hynTokenAddr);
        uint256 _balanceBefore = hynToken.balanceOf(msg.sender);
        hynToken.safeTransferFrom(msg.sender, address(this), amount);
        uint256 _balanceAfter = hynToken.balanceOf(msg.sender);
        uint256 _actualAmount = _balanceBefore.sub(_balanceAfter);
        emit Locked(address(hynToken), msg.sender, _actualAmount, recipient);
    }

    /**
     * @dev lock tokens for a user address to be minted on heco chain
     * @param hynTokenAddr is the atlas token contract address
     * @param userAddr is token holder address
     * @param amount amount of tokens to lock
     * @param recipient recipient address on the heco chain
     */
    function lockTokenFor(
        address hynTokenAddr,
        address userAddr,
        uint256 amount,
        address recipient
    ) public onlyWallet {
        require(
            recipient != address(0),
            "AtlasManager/recipient is a zero address"
        );
        require(amount > 0, "AtlasManager/zero token locked");
        IERC20 hynToken = IERC20(hynTokenAddr);
        uint256 _balanceBefore = hynToken.balanceOf(userAddr);
        hynToken.safeTransferFrom(userAddr, address(this), amount);
        uint256 _balanceAfter = hynToken.balanceOf(userAddr);
        uint256 _actualAmount = _balanceBefore.sub(_balanceAfter);
        emit Locked(address(hynToken), userAddr, _actualAmount, recipient);
    }

    /**
     * @dev unlock tokens after burning them on heco chain
     * @param hynTokenAddr is the atlas token contract address
     * @param amount amount of unlock tokens
     * @param recipient recipient of the unlock tokens
     * @param receiptId transaction hash of the burn event on heco chain
     */
    function unlockToken(
        address hynTokenAddr,
        uint256 amount,
        address recipient,
        bytes32 receiptId
    ) public onlyWallet {
        require(
            !usedEvents_[receiptId],
            "AtlasManager/The burn event cannot be reused"
        );
        IERC20 hynToken = IERC20(hynTokenAddr);
        usedEvents_[receiptId] = true;
        hynToken.safeTransfer(recipient, amount);
        emit Unlocked(hynTokenAddr, amount, recipient, receiptId);
    }

    /**
     * @dev lock HYNs to be minted on heco chain
     * @param amount amount of tokens to lock
     * @param recipient recipient address on the harmony chain
     */
    function lockHyn(
        uint256 amount,
        address recipient
    ) public payable {
        require(
            recipient != address(0),
            "AtlasManager/recipient is a zero address"
        );
        require(msg.value == amount, "HynManager/zero token locked");
        emit Locked(address(HYN_ADDRESS), msg.sender, amount, recipient);
    }

    /**
     * @dev unlock HYNs after burning them on heco chain
     * @param amount amount of unlock tokens
     * @param recipient recipient of the unlock tokens
     * @param receiptId transaction hash of the burn event on harmony chain
     */
    function unlockHyn(
        uint256 amount,
        address payable recipient,
        bytes32 receiptId
    ) public onlyWallet {
        require(
            !usedEvents_[receiptId],
            "AtlasManager/The burn event cannot be reused"
        );
        usedEvents_[receiptId] = true;
        recipient.transfer(amount);
        emit Unlocked(address(HYN_ADDRESS), amount, recipient, receiptId);
    }
}
