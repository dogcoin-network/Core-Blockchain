package systemcontract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
)

// ValidatorsInteractiveABI contains all methods to interactive with validator contracts.
const ValidatorsInteractiveABI = `[{"type":"event","name":"LogAddToTopValidators","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogCreateValidator","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"address","name":"fee","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogDistributeBlockReward","inputs":[{"type":"address","name":"coinbase","internalType":"address","indexed":true},{"type":"uint256","name":"blockReward","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false},{"type":"address[]","name":"To","internalType":"address[]","indexed":false},{"type":"uint64[]","name":"Gass","internalType":"uint64[]","indexed":false}],"anonymous":false},{"type":"event","name":"LogEditValidator","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"address","name":"fee","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogReactive","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogRemoveFromTopValidators","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogRemoveValidator","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"hb","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogRemoveValidatorIncoming","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"hb","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogStake","inputs":[{"type":"address","name":"staker","internalType":"address","indexed":true},{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"staking","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogUnstake","inputs":[{"type":"address","name":"staker","internalType":"address","indexed":true},{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"amount","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogUpdateValidator","inputs":[{"type":"address[]","name":"newSet","internalType":"address[]","indexed":false}],"anonymous":false},{"type":"event","name":"LogWithdrawProfits","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"address","name":"fee","internalType":"address","indexed":true},{"type":"uint256","name":"hb","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogWithdrawStaking","inputs":[{"type":"address","name":"staker","internalType":"address","indexed":true},{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"amount","internalType":"uint256","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"withdrawStakingRewardEv","inputs":[{"type":"address","name":"user","internalType":"address","indexed":false},{"type":"address","name":"validator","internalType":"address","indexed":false},{"type":"uint256","name":"reward","internalType":"uint256","indexed":false},{"type":"uint256","name":"timeStamp","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"function","stateMutability":"view","outputs":[{"type":"uint16","name":"","internalType":"uint16"}],"name":"MaxValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"MinimalStakingCoin","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ProposalAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"PunishContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"StakingLockPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ValidatorContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"WithdrawProfitPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address payable"}],"name":"additionalWallet","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"additionalWalletPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnStopAmount","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"contractCreator","inputs":[{"type":"address","name":"","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"contractPartPercent","inputs":[]},{"type":"function","stateMutability":"payable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"createOrEditValidator","inputs":[{"type":"address","name":"feeAddr","internalType":"address payable"},{"type":"string","name":"moniker","internalType":"string"},{"type":"string","name":"identity","internalType":"string"},{"type":"string","name":"website","internalType":"string"},{"type":"string","name":"email","internalType":"string"},{"type":"string","name":"details","internalType":"string"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"currentValidatorSet","inputs":[{"type":"uint256","name":"","internalType":"uint256"}]},{"type":"function","stateMutability":"payable","outputs":[],"name":"distributeBlockReward","inputs":[{"type":"address[]","name":"_to","internalType":"address[]"},{"type":"uint64[]","name":"_gass","internalType":"uint64[]"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address[]","name":"","internalType":"address[]"}],"name":"getActiveValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"},{"type":"uint256","name":"","internalType":"uint256"},{"type":"uint256","name":"","internalType":"uint256"}],"name":"getStakingInfo","inputs":[{"type":"address","name":"staker","internalType":"address"},{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address[]","name":"","internalType":"address[]"}],"name":"getTopValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"total","internalType":"uint256"},{"type":"uint256","name":"len","internalType":"uint256"}],"name":"getTotalStakeOfActiveValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"string","name":"","internalType":"string"},{"type":"string","name":"","internalType":"string"},{"type":"string","name":"","internalType":"string"},{"type":"string","name":"","internalType":"string"},{"type":"string","name":"","internalType":"string"}],"name":"getValidatorDescription","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address payable"},{"type":"uint8","name":"","internalType":"enum Validators.Status"},{"type":"uint256","name":"","internalType":"uint256"},{"type":"uint256","name":"","internalType":"uint256"},{"type":"uint256","name":"","internalType":"uint256"},{"type":"uint256","name":"","internalType":"uint256"},{"type":"address[]","name":"","internalType":"address[]"}],"name":"getValidatorInfo","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"highestValidatorsSet","inputs":[{"type":"uint256","name":"","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"initialize","inputs":[{"type":"address[]","name":"vals","internalType":"address[]"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"initialized","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"isActiveValidator","inputs":[{"type":"address","name":"who","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"isTopValidator","inputs":[{"type":"address","name":"who","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"lastRewardTime","inputs":[{"type":"address","name":"","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"minimumValidatorStaking","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"reflectionPercentSum","inputs":[{"type":"address","name":"","internalType":"address"},{"type":"uint256","name":"","internalType":"uint256"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"removeValidator","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"removeValidatorIncoming","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"setContractCreator","inputs":[{"type":"address","name":"_contract","internalType":"address"}]},{"type":"function","stateMutability":"payable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"stake","inputs":[{"type":"address","name":"validator","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"stakeTime","inputs":[{"type":"address","name":"","internalType":"address"},{"type":"address","name":"","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"stakerPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalBurnt","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalJailedHB","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalStake","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"tryReactive","inputs":[{"type":"address","name":"validator","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"unstake","inputs":[{"type":"address","name":"validator","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"updateActiveValidatorSet","inputs":[{"type":"address[]","name":"newSet","internalType":"address[]"},{"type":"uint256","name":"epoch","internalType":"uint256"}]},{"type":"function","stateMutability":"pure","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"validateDescription","inputs":[{"type":"string","name":"moniker","internalType":"string"},{"type":"string","name":"identity","internalType":"string"},{"type":"string","name":"website","internalType":"string"},{"type":"string","name":"email","internalType":"string"},{"type":"string","name":"details","internalType":"string"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"validatorPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"viewStakeReward","inputs":[{"type":"address","name":"_staker","internalType":"address"},{"type":"address","name":"_validator","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"withdrawProfits","inputs":[{"type":"address","name":"validator","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"withdrawStaking","inputs":[{"type":"address","name":"validator","internalType":"address"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"withdrawStakingReward","inputs":[{"type":"address","name":"validator","internalType":"address"}]}]`

const PunishInteractiveABI = `[{"type":"event","name":"LogDecreaseMissedBlocksCounter","inputs":[],"anonymous":false},{"type":"event","name":"LogPunishValidator","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"function","stateMutability":"view","outputs":[{"type":"uint16","name":"","internalType":"uint16"}],"name":"MaxValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"MinimalStakingCoin","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ProposalAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"PunishContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"StakingLockPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ValidatorContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"WithdrawProfitPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnStopAmount","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"cleanPunishRecord","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"contractPartPercent","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"decreaseMissedBlocksCounter","inputs":[{"type":"uint256","name":"epoch","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"decreaseRate","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getPunishRecord","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"getPunishValidatorsLen","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"initialize","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"initialized","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"minimumValidatorStaking","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"punish","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"punishThreshold","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"punishValidators","inputs":[{"type":"uint256","name":"","internalType":"uint256"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"removeThreshold","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"stakerPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalBurnt","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"validatorPartPercent","inputs":[]}]`

const ProposalInteractiveABI = `[{"type":"event","name":"LogCreateProposal","inputs":[{"type":"bytes32","name":"id","internalType":"bytes32","indexed":true},{"type":"address","name":"proposer","internalType":"address","indexed":true},{"type":"address","name":"dst","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogPassProposal","inputs":[{"type":"bytes32","name":"id","internalType":"bytes32","indexed":true},{"type":"address","name":"dst","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogRejectProposal","inputs":[{"type":"bytes32","name":"id","internalType":"bytes32","indexed":true},{"type":"address","name":"dst","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogSetUnpassed","inputs":[{"type":"address","name":"val","internalType":"address","indexed":true},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"LogVote","inputs":[{"type":"bytes32","name":"id","internalType":"bytes32","indexed":true},{"type":"address","name":"voter","internalType":"address","indexed":true},{"type":"bool","name":"auth","internalType":"bool","indexed":false},{"type":"uint256","name":"time","internalType":"uint256","indexed":false}],"anonymous":false},{"type":"function","stateMutability":"view","outputs":[{"type":"uint16","name":"","internalType":"uint16"}],"name":"MaxValidators","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"MinimalStakingCoin","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ProposalAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"PunishContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"StakingLockPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"","internalType":"address"}],"name":"ValidatorContractAddr","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint64","name":"","internalType":"uint64"}],"name":"WithdrawProfitPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"burnStopAmount","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"contractPartPercent","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"createProposal","inputs":[{"type":"address","name":"dst","internalType":"address"},{"type":"string","name":"details","internalType":"string"}]},{"type":"function","stateMutability":"nonpayable","outputs":[],"name":"initialize","inputs":[{"type":"address[]","name":"vals","internalType":"address[]"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"initialized","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"minimumValidatorStaking","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"pass","inputs":[{"type":"address","name":"","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"proposalLastingPeriod","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"proposer","internalType":"address"},{"type":"address","name":"dst","internalType":"address"},{"type":"string","name":"details","internalType":"string"},{"type":"uint256","name":"createTime","internalType":"uint256"},{"type":"uint16","name":"agree","internalType":"uint16"},{"type":"uint16","name":"reject","internalType":"uint16"},{"type":"bool","name":"resultExist","internalType":"bool"}],"name":"proposals","inputs":[{"type":"bytes32","name":"","internalType":"bytes32"}]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"setUnpassed","inputs":[{"type":"address","name":"val","internalType":"address"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"stakerPartPercent","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"totalBurnt","inputs":[]},{"type":"function","stateMutability":"view","outputs":[{"type":"uint256","name":"","internalType":"uint256"}],"name":"validatorPartPercent","inputs":[]},{"type":"function","stateMutability":"nonpayable","outputs":[{"type":"bool","name":"","internalType":"bool"}],"name":"voteProposal","inputs":[{"type":"bytes32","name":"id","internalType":"bytes32"},{"type":"bool","name":"auth","internalType":"bool"}]},{"type":"function","stateMutability":"view","outputs":[{"type":"address","name":"voter","internalType":"address"},{"type":"uint256","name":"voteTime","internalType":"uint256"},{"type":"bool","name":"auth","internalType":"bool"}],"name":"votes","inputs":[{"type":"address","name":"","internalType":"address"},{"type":"bytes32","name":"","internalType":"bytes32"}]}]`

const SysGovInteractiveABI = `
[
    {
		"inputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			}
		],
		"name": "finishProposalById",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint32",
				"name": "index",
				"type": "uint32"
			}
		],
		"name": "getPassedProposalByIndex",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
        		"internalType": "uint256",
        		"name": "action",
        		"type": "uint256"
        	},
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getPassedProposalCount",
		"outputs": [
			{
				"internalType": "uint32",
				"name": "",
				"type": "uint32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_admin",
				"type": "address"
			}
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

const AddrListInteractiveABI = `
[
	{
	  "inputs": [],
	  "name": "blackLastUpdatedNumber",
	  "outputs": [
		{
		  "internalType": "uint256",
		  "name": "",
		  "type": "uint256"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "devVerifyEnabled",
	  "outputs": [
		{
		  "internalType": "bool",
		  "name": "",
		  "type": "bool"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "getBlacksFrom",
	  "outputs": [
		{
		  "internalType": "address[]",
		  "name": "",
		  "type": "address[]"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "getBlacksTo",
	  "outputs": [
		{
		  "internalType": "address[]",
		  "name": "",
		  "type": "address[]"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [
		{
		  "internalType": "uint32",
		  "name": "i",
		  "type": "uint32"
		}
	  ],
	  "name": "getRuleByIndex",
	  "outputs": [
		{
		  "internalType": "bytes32",
		  "name": "",
		  "type": "bytes32"
		},
		{
		  "internalType": "uint128",
		  "name": "",
		  "type": "uint128"
		},
		{
		  "internalType": "enum AddressList.CheckType",
		  "name": "",
		  "type": "uint8"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "initializeV2",
	  "outputs": [],
	  "stateMutability": "nonpayable",
	  "type": "function"
	},
	{
	  "inputs": [
		{
		  "internalType": "address",
		  "name": "_admin",
		  "type": "address"
		}
	  ],
	  "name": "initialize",
	  "outputs": [],
	  "stateMutability": "nonpayable",
	  "type": "function"
	},
	{
	  "inputs": [
		{
		  "internalType": "address",
		  "name": "addr",
		  "type": "address"
		}
	  ],
	  "name": "isDeveloper",
	  "outputs": [
		{
		  "internalType": "bool",
		  "name": "",
		  "type": "bool"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "rulesLastUpdatedNumber",
	  "outputs": [
		{
		  "internalType": "uint256",
		  "name": "",
		  "type": "uint256"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	},
	{
	  "inputs": [],
	  "name": "rulesLen",
	  "outputs": [
		{
		  "internalType": "uint32",
		  "name": "",
		  "type": "uint32"
		}
	  ],
	  "stateMutability": "view",
	  "type": "function"
	}
]`

const ValidatorsV1InteractiveABI = `[
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "name": "activeValidators",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "distributeBlockReward",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "getTopValidators",
        "outputs": [
            {
                "internalType": "address[]",
                "name": "",
                "type": "address[]"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address[]",
                "name": "_candidates",
                "type": "address[]"
            },
            {
                "internalType": "address[]",
                "name": "_manager",
                "type": "address[]"
            },
            {
                "internalType": "address",
                "name": "_admin",
                "type": "address"
            }
        ],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address[]",
                "name": "newSet",
                "type": "address[]"
            },
            {
                "internalType": "uint256",
                "name": "epoch",
                "type": "uint256"
            }
        ],
        "name": "updateActiveValidatorSet",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    }
]`

const PunishV1InteractiveABI = `[
    {
      "inputs": [],
      "name": "initialize",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
]`

// DevMappingPosition is the position of the state variable `devs`.
// Since the state variables are as follow:
//    bool public initialized;
//    bool public devVerifyEnabled;
//    address public admin;
//    address public pendingAdmin;
//
//    mapping(address => bool) private devs;
//
//    //NOTE: make sure this list is not too large!
//    address[] blacksFrom;
//    address[] blacksTo;
//    mapping(address => uint256) blacksFromMap;      // address => index+1
//    mapping(address => uint256) blacksToMap;        // address => index+1
//
//    uint256 public blackLastUpdatedNumber; // last block number when the black list is updated
//    uint256 public rulesLastUpdatedNumber;  // last block number when the rules are updated
//    // event check rules
//    EventCheckRule[] rules;
//    mapping(bytes32 => mapping(uint128 => uint256)) rulesMap;   // eventSig => checkIdx => indexInArray+1
//
// according to [Layout of State Variables in Storage](https://docs.soliditylang.org/en/v0.8.4/internals/layout_in_storage.html),
// and after optimizer enabled, the `initialized`, `enabled` and `admin` will be packed, and stores at slot 0,
// `pendingAdmin` stores at slot 1, so the position for `devs` is 2.
const DevMappingPosition = 2

var (
	BlackLastUpdatedNumberPosition = common.BytesToHash([]byte{0x07})
	RulesLastUpdatedNumberPosition = common.BytesToHash([]byte{0x08})
)

var (
	ValidatorsContractName   = "validators"
	PunishContractName       = "punish"
	ProposalContractName     = "proposal"
	SysGovContractName       = "governance"
	AddressListContractName  = "address_list"
	ValidatorsV1ContractName = "validators_v1"
	PunishV1ContractName     = "punish_v1"
	ValidatorsContractAddr   = common.HexToAddress("0x000000000000000000000000000000000000f000")
	PunishContractAddr       = common.HexToAddress("0x000000000000000000000000000000000000f001")
	ProposalAddr             = common.HexToAddress("0x000000000000000000000000000000000000f002")
	SysGovContractAddr       = common.HexToAddress("0x000000000000000000000000000000000000F003")
	AddressListContractAddr  = common.HexToAddress("0x000000000000000000000000000000000000F004")
	ValidatorsV1ContractAddr = common.HexToAddress("0x000000000000000000000000000000000000F005")
	PunishV1ContractAddr     = common.HexToAddress("0x000000000000000000000000000000000000F006")
	// SysGovToAddr is the To address for the system governance transaction, NOT contract address
	SysGovToAddr = common.HexToAddress("0x000000000000000000000000000000000000ffff")

	abiMap map[string]abi.ABI
)

func init() {
	abiMap = make(map[string]abi.ABI, 0)
	tmpABI, _ := abi.JSON(strings.NewReader(ValidatorsInteractiveABI))
	abiMap[ValidatorsContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(PunishInteractiveABI))
	abiMap[PunishContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(ProposalInteractiveABI))
	abiMap[ProposalContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(SysGovInteractiveABI))
	abiMap[SysGovContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(AddrListInteractiveABI))
	abiMap[AddressListContractName] = tmpABI

	tmpABI, _ = abi.JSON(strings.NewReader(ValidatorsV1InteractiveABI))
	abiMap[ValidatorsV1ContractName] = tmpABI
	tmpABI, _ = abi.JSON(strings.NewReader(PunishV1InteractiveABI))
	abiMap[PunishV1ContractName] = tmpABI
}

func GetInteractiveABI() map[string]abi.ABI {
	return abiMap
}

func GetValidatorAddr(blockNum *big.Int, config *params.ChainConfig) *common.Address {
	if config.IsRedCoast(blockNum) {
		return &ValidatorsV1ContractAddr
	}
	return &ValidatorsContractAddr
}

func GetPunishAddr(blockNum *big.Int, config *params.ChainConfig) *common.Address {
	if config.IsRedCoast(blockNum) {
		return &PunishV1ContractAddr
	}
	return &PunishContractAddr
}
