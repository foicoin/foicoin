// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://688c1866a34ee43cec42d23aaa66b1868cc6d55fe58148831adc814512523c2412bc9ab481469dca07f354b8820985a92730a0accd5ee1209619a308d380c6f5@104.42.143.8:30303",//windows server bootnode
	"enode://113ac360330377da000d5fec4457b60b4bc9baa5801cfabc0aa28b722943c1fd974d05fe3717734e9c6c4e9da9cf81ecc172055521e4d28e81cee51ccf614466@52.231.195.25:30303",//Linux server bootnode
	"enode://b2490e8dae3dfd84c143bdb2193dc522303aa8d8cabce94de8fabfeb56c84dc52959ffd42cf133e1dea480a27bb1beceaf66d02af6ce6559112cdc574950a75e@52.185.141.160:30303",
	"enode://e19c546bd65634f1e5956f134b5e6ecd1ce89e09c69802dfb4d57b40e0fb84622fa28f4ceb3c103969e2b05a7c4281c64cb833a5d5db0c8424a0c29cd68f35b1@104.215.6.178:30303",//Linux backup bootnodes
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://688c1866a34ee43cec42d23aaa66b1868cc6d55fe58148831adc814512523c2412bc9ab481469dca07f354b8820985a92730a0accd5ee1209619a308d380c6f5@104.42.143.8:30303",//windows server bootnode
	"enode://113ac360330377da000d5fec4457b60b4bc9baa5801cfabc0aa28b722943c1fd974d05fe3717734e9c6c4e9da9cf81ecc172055521e4d28e81cee51ccf614466@52.231.195.25:30303",//Linux server bootnode
	"enode://b2490e8dae3dfd84c143bdb2193dc522303aa8d8cabce94de8fabfeb56c84dc52959ffd42cf133e1dea480a27bb1beceaf66d02af6ce6559112cdc574950a75e@52.185.141.160:30303",
	"enode://e19c546bd65634f1e5956f134b5e6ecd1ce89e09c69802dfb4d57b40e0fb84622fa28f4ceb3c103969e2b05a7c4281c64cb833a5d5db0c8424a0c29cd68f35b1@104.215.6.178:30303",//Linux backup bootnodes
}
