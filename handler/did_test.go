// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	hash2 "github.com/iotexproject/go-pkgs/hash"

	"github.com/lzxm160/testswagger/contract"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

func TestDidDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(contract.IoTeXDIDBin[2:])
	require.NoError(err)

	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidProxyDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(contract.IoTeXDIDProxyBin[2:])
	require.NoError(err)
	DIDcontract, err := address.FromString(IoTeXDID_address)
	require.NoError(err)
	abi, err := abi.JSON(strings.NewReader(contract.IoTeXDIDProxyABI))
	require.NoError(err)
	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).SetArgs(abi, DIDcontract).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidProxyDeployUpdateAddress(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(contract.IoTeXDIDProxyABI))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	IoTeXDIDAddress, err := address.FromString(IoTeXDID_address)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(IoTeXDIDAddress.Bytes()))
	hash, err := c.Contract(contract, abi).Execute("upgradeTo", "1.1.1", ethAddress).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidProxyReadVersion(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(contract.IoTeXDIDProxyABI))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	version := "0.0.1"
	//version := "1.1.1"
	ret, err := c.ReadOnlyContract(contract, abi).Read("getImplFromVersion", version).Call(context.Background())
	require.NoError(err)
	//fmt.Println(ret.method)
	fmt.Println(hex.EncodeToString(ret.Raw))
	var addr common.Address
	require.NoError(ret.Unmarshal(&addr))
	fmt.Println(addr.String())

	iotexAddr, err := address.FromBytes(addr.Bytes())
	require.NoError(err)
	fmt.Println(iotexAddr.String())
}

func TestDidProxyReadLatestVersion(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(contract.IoTeXDIDProxyABI))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	ret, err := c.ReadOnlyContract(contract, abi).Read("implementation").Call(context.Background())
	require.NoError(err)
	//fmt.Println(ret.method)
	fmt.Println(hex.EncodeToString(ret.Raw))
	var addr common.Address
	require.NoError(ret.Unmarshal(&addr))
	fmt.Println(addr.String())

	iotexAddr, err := address.FromBytes(addr.Bytes())
	require.NoError(err)
	fmt.Println(iotexAddr.String())
}

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	h, err := d.CreateDID("", "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed", "uri")
	require.NoError(err)
	fmt.Println(h)

	time.Sleep(20 * time.Second)
	decode, err := hex.DecodeString(h)
	hash := hash2.Hash256b(decode)
	receiptResponse, err := d.GetCli().GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidDelete(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.DeleteDID(didString)
	require.NoError(err)
	fmt.Println(h)

	time.Sleep(20 * time.Second)
	decode, err := hex.DecodeString(h)
	hash := hash2.Hash256b(decode)
	receiptResponse, err := d.GetCli().GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidUpdateHash(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.UpdateHash(didString, "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bxx")
	require.NoError(err)
	fmt.Println(h)

	time.Sleep(20 * time.Second)
	decode, err := hex.DecodeString(h)
	hash := hash2.Hash256b(decode)
	receiptResponse, err := d.GetCli().GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidUpdateUri(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.UpdateUri(didString, "uri2")
	require.NoError(err)
	fmt.Println(h)

	time.Sleep(20 * time.Second)
	decode, err := hex.DecodeString(h)
	hash := hash2.Hash256b(decode)
	receiptResponse, err := d.GetCli().GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidReadHashContract(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.GetHash(didString)
	require.NoError(err)
	fmt.Println(h)
}

func TestDidReadUriContract(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	uri, err := d.GetUri(didString)
	require.NoError(err)
	fmt.Println(uri)
}
