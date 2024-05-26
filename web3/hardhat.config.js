/** @type import('hardhat/config').HardhatUserConfig */
require('dotenv').config();
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const RPC_URL = process.env.RPC_URL;

module.exports = {
  defaultNetwork: "eth_sepolia",
  networks: {
    hardhat: {
      chainId: 11155111,
    },
    eth_sepolia: {
      url: RPC_URL,
      accounts: [`0x${PRIVATE_KEY}`]
    }
  },

  solidity: {
    version: "0.8.17",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
};
