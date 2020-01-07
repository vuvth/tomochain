package lendingstate

import (
	"fmt"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/core/state"
	"github.com/tomochain/tomochain/crypto"
	"math/big"
)

var (
	LendingRelayerListSlot    = uint64(0)
	CollateralMapSlot         = uint64(1)
	DefaultCollateralSlot     = uint64(2)
	SupportedBaseSlot         = uint64(3)
	SupportedTermSlot         = uint64(4)
	LendingRelayerStructSlots = map[string]*big.Int{
		"fee":         big.NewInt(0),
		"bases":       big.NewInt(1),
		"terms":       big.NewInt(2),
		"collaterals": big.NewInt(3),
	}
	CollateralStructSlots = map[string]*big.Int{
		"depositRate":     big.NewInt(0),
		"liquidationRate": big.NewInt(1),
		"price":           big.NewInt(2),
	}
)

// @function IsValidRelayer : return whether the given address is the coinbase of a valid relayer or not
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @return: true if it's a valid coinbase address of lending protocol, otherwise return false
func IsValidRelayer(statedb *state.StateDB, coinbase common.Address) bool {
	locRelayerState := GetLocMappingAtKey(coinbase.Hash(), LendingRelayerListSlot)

	if v := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), common.BytesToHash(locRelayerState.Bytes())); v != (common.Hash{}) {
		return true
	}
	return false
}

// @function GetFee
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @return: feeRate of lending
func GetFee(statedb *state.StateDB, coinbase common.Address) *big.Int {
	locRelayerState := state.GetLocMappingAtKey(coinbase.Hash(), LendingRelayerListSlot)
	locHash := common.BytesToHash(new(big.Int).Add(locRelayerState, LendingRelayerStructSlots["fee"]).Bytes())
	return statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locHash).Big()
}

// @function GetBaseList
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @return: list of base tokens
func GetBaseList(statedb *state.StateDB, coinbase common.Address) []common.Address {
	baseList := []common.Address{}
	locRelayerState := state.GetLocMappingAtKey(coinbase.Hash(), LendingRelayerListSlot)
	locBaseHash := state.GetLocOfStructElement(locRelayerState, LendingRelayerStructSlots["bases"])
	length := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locBaseHash).Big().Uint64()
	for i := uint64(0); i < length; i++ {
		loc := state.GetLocDynamicArrAtElement(locBaseHash, i, 1)
		addr := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Bytes())
		if addr != (common.Address{}) {
			baseList = append(baseList, addr)
		}
	}
	return baseList
}

// @function GetTerms
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @return: list of supported terms of the given relayer
func GetTerms(statedb *state.StateDB, coinbase common.Address) []uint64 {
	terms := []uint64{}
	locRelayerState := state.GetLocMappingAtKey(coinbase.Hash(), LendingRelayerListSlot)
	locTermHash := state.GetLocOfStructElement(locRelayerState, LendingRelayerStructSlots["terms"])
	length := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locTermHash).Big().Uint64()
	for i := uint64(0); i < length; i++ {
		loc := state.GetLocDynamicArrAtElement(locTermHash, i, 1)
		t := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Big().Uint64()
		if t != uint64(0) {
			terms = append(terms, t)
		}
	}
	return terms
}

// @function IsValidPair
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @param baseToken: address of baseToken
// @param terms: term
// @return: TRUE if the given baseToken, term organize a valid pair
func IsValidPair(statedb *state.StateDB, coinbase common.Address, baseToken common.Address, term uint64) (valid bool, pairIndex uint64) {
	baseTokenList := GetBaseList(statedb, coinbase)
	terms := GetTerms(statedb, coinbase)
	baseIndexes := []uint64{}
	for i := uint64(0); i < uint64(len(baseTokenList)); i++ {
		if baseTokenList[i] == baseToken {
			baseIndexes = append(baseIndexes, i)
		}
	}
	for _, index := range baseIndexes {
		if terms[index] == term {
			pairIndex = index
			return true, pairIndex
		}
	}
	return false, pairIndex
}

// @function GetCollaterals
// @param statedb : current state
// @param coinbase: coinbase address of relayer
// @param baseToken: address of baseToken
// @param terms: term
// @return:
//		- collaterals []common.Address  : list of addresses of collateral
//		- isSpecialCollateral			: TRUE if collateral is a token which is NOT available for trading in TomoX, otherwise FALSE
func GetCollaterals(statedb *state.StateDB, coinbase common.Address, baseToken common.Address, term uint64) (collaterals []common.Address, isSpecialCollateral bool) {
	validPair, pairIndex := IsValidPair(statedb, coinbase, baseToken, term)
	if !validPair {
		return []common.Address{}, false
	}

	locRelayerState := state.GetLocMappingAtKey(coinbase.Hash(), LendingRelayerListSlot)
	locCollateralHash := state.GetLocOfStructElement(locRelayerState, LendingRelayerStructSlots["collaterals"])
	length := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locCollateralHash).Big().Uint64()

	loc := state.GetLocDynamicArrAtElement(locCollateralHash, pairIndex, 1)
	collateralAddr := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Bytes())
	if collateralAddr != (common.Address{}) && collateralAddr != (common.HexToAddress("0x0")) {
		return []common.Address{collateralAddr}, true
	}

	// if collaterals is not defined for the relayer, return default collaterals
	locDefaultCollateralHash := state.GetLocSimpleVariable(DefaultCollateralSlot)
	length = statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locDefaultCollateralHash).Big().Uint64()
	for i := uint64(0); i < length; i++ {
		loc := state.GetLocDynamicArrAtElement(locDefaultCollateralHash, i, 1)
		addr := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Bytes())
		if addr != (common.Address{}) {
			collaterals = append(collaterals, addr)
		}
	}
	return collaterals, false
}

// @function GetCollateralDetail
// @param statedb : current state
// @param token: address of collateral token
// @return: depositRate, liquidationRate, price of collateral
func GetCollateralDetail(statedb *state.StateDB, token common.Address) (depositRate *big.Int, liquidationRate *big.Int, price *big.Int) {
	collateralState := GetLocMappingAtKey(token.Hash(), CollateralMapSlot)
	locDepositRate := state.GetLocOfStructElement(collateralState, CollateralStructSlots["depositRate"])
	locLiquidationRate := state.GetLocOfStructElement(collateralState, CollateralStructSlots["liquidationRate"])
	locCollateralPrice := state.GetLocOfStructElement(collateralState, CollateralStructSlots["price"])
	depositRate = statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locDepositRate).Big()
	liquidationRate = statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locLiquidationRate).Big()
	price = statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locCollateralPrice).Big()
	return depositRate, liquidationRate, price
}

// @function GetSupportedTerms
// @param statedb : current state
// @return: list of terms which tomoxlending supports
func GetSupportedTerms(statedb *state.StateDB) []uint64 {
	terms := []uint64{}
	locSupportedTerm := state.GetLocSimpleVariable(SupportedTermSlot)
	length := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locSupportedTerm).Big().Uint64()
	for i := uint64(0); i < length; i++ {
		loc := state.GetLocDynamicArrAtElement(locSupportedTerm, i, 1)
		t := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Big().Uint64()
		if t != 0 {
			terms = append(terms, t)
		}
	}
	return terms
}

// @function GetSupportedBaseToken
// @param statedb : current state
// @return: list of tokens which are available for lending
func GetSupportedBaseToken(statedb *state.StateDB) []common.Address {
	baseTokens := []common.Address{}
	locSupportedBaseToken := state.GetLocSimpleVariable(SupportedBaseSlot)
	length := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), locSupportedBaseToken).Big().Uint64()
	for i := uint64(0); i < length; i++ {
		loc := state.GetLocDynamicArrAtElement(locSupportedBaseToken, i, 1)
		addr := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), loc).Bytes())
		if addr != (common.Address{}) {
			baseTokens = append(baseTokens, addr)
		}
	}
	return baseTokens
}

func GetRelayerCount(statedb *state.StateDB) uint64 {
	slot := RelayerMappingSlot["RelayerCount"]
	slotHash := common.BigToHash(new(big.Int).SetUint64(slot))
	valueHash := statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), slotHash)
	return new(big.Int).SetBytes(valueHash.Bytes()).Uint64()
}

func GetAllCoinbases(statedb *state.StateDB) []common.Address {
	relayerCount := GetRelayerCount(statedb)
	slot := RelayerMappingSlot["RELAYER_COINBASES"]
	slotHash := common.BigToHash(new(big.Int).SetUint64(slot))
	coinbases := []common.Address{}
	for i := uint64(0); i < relayerCount; i++ {
		retByte := crypto.Keccak256(new(big.Int).SetUint64(i).Bytes(), slotHash.Bytes())
		valueHash := statedb.GetState(common.HexToAddress(common.LendingRegistrationSMC), common.BytesToHash(retByte))
		coinbases = append(coinbases, common.BytesToAddress(valueHash.Bytes()))
	}
	return coinbases
}
func GetAllLendingPairs(statedb *state.StateDB) ([]LendingPair, error) {
	coinbases := GetAllCoinbases(statedb)
	slot := RelayerMappingSlot["RELAYER_LIST"]
	allPairs := []LendingPair{}
	for _, coinbase := range coinbases {
		locBig := GetLocMappingAtKey(coinbase.Hash(), slot)
		fromTokenSlot := new(big.Int).Add(locBig, RelayerStructMappingSlot["_fromTokens"])
		fromTokenLength := statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(fromTokenSlot)).Big().Uint64()
		toTokenSlot := new(big.Int).Add(locBig, RelayerStructMappingSlot["_toTokens"])
		toTokenLength := statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(toTokenSlot)).Big().Uint64()
		if toTokenLength != fromTokenLength {
			return []LendingPair{}, fmt.Errorf("Invalid length from token & to toke : from :%d , to :%d ", fromTokenLength, toTokenLength)
		}
		fromTokens := []common.Address{}
		for i := uint64(0); i < fromTokenLength; i++ {
			slotKecBig := crypto.Keccak256Hash(fromTokenSlot.Bytes()).Big()
			arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(i))
			fromToken := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(arrBig)).Bytes())
			fromTokens = append(fromTokens, fromToken)
		}
		for i := uint64(0); i < toTokenLength; i++ {
			slotKecBig := crypto.Keccak256Hash(toTokenSlot.Bytes()).Big()
			arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(i))
			toToken := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(arrBig)).Bytes())
			allPairs = append(allPairs, LendingPair{LendingToken: fromTokens[i], CollateralToken: toToken})
		}
	}
	return allPairs, nil
}

func GetAllLendingBooks(statedb *state.StateDB) (map[common.Hash]bool, error) {
	coinbases := GetAllCoinbases(statedb)
	slot := RelayerMappingSlot["RELAYER_LIST"]
	allPairs := map[common.Hash]bool{}
	for _, coinbase := range coinbases {
		locBig := GetLocMappingAtKey(coinbase.Hash(), slot)
		fromTokenSlot := new(big.Int).Add(locBig, RelayerStructMappingSlot["_fromTokens"])
		fromTokenLength := statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(fromTokenSlot)).Big().Uint64()
		toTokenSlot := new(big.Int).Add(locBig, RelayerStructMappingSlot["_toTokens"])
		toTokenLength := statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(toTokenSlot)).Big().Uint64()
		if toTokenLength != fromTokenLength {
			return map[common.Hash]bool{}, fmt.Errorf("Invalid length from token & to toke : from :%d , to :%d ", fromTokenLength, toTokenLength)
		}
		fromTokens := []common.Address{}
		for i := uint64(0); i < fromTokenLength; i++ {
			slotKecBig := crypto.Keccak256Hash(fromTokenSlot.Bytes()).Big()
			arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(i))
			fromToken := common.BytesToAddress(statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(arrBig)).Bytes())
			fromTokens = append(fromTokens, fromToken)
		}
		for i := uint64(0); i < toTokenLength; i++ {
			slotKecBig := crypto.Keccak256Hash(toTokenSlot.Bytes()).Big()
			arrBig := slotKecBig.Add(slotKecBig, new(big.Int).SetUint64(i))
			toToken := common.BytesToHash(statedb.GetState(common.HexToAddress(common.RelayerRegistrationSMC), common.BigToHash(arrBig)).Bytes())
			allPairs[GetLendingOrderBookHash(fromTokens[i], toToken.Big().Uint64())] = true
		}
	}
	return allPairs, nil
}