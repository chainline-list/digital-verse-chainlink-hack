package registry

import (
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
	"gitlab.com/digitalverse/blockchain/pkg/blockchain/generic"
)

func CreateRegistry(c *blockchain.Config) map[string]blockchain.Network {
	return map[string]blockchain.Network{
		"heco": generic.NewGenericNetwork(
			"https://testnet.hecoinfo.com/tx/%s",
			c.HecoEndpointUrl,
			c.HecoDeployWalletPk,
			c.HecoNftContractAddress,
			-1,
		),
		"eth": generic.NewGenericNetwork(
			"https://rinkeby.etherscan.io/tx/%s",
			c.RinkebyEndpointUrl,
			c.RinkebyDeployWalletPk,
			c.RinkebyNftContractAddress,
			-1,
		),
		"polygon": generic.NewGenericNetwork(
			"https://mumbai.polygonscan.com/tx/%s",
			c.PolygonEndpointUrl,
			c.PolygonDeployWalletPk,
			c.PolygonNftContractAddress,
			c.PolygonShardID,
		),
	}
}
