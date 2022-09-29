package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// LProviderRecordKeyPrefix is the prefix to retrieve all LProviderRecord
	LProviderRecordKeyPrefix = "LProviderRecord/value/"
	RefKeyPrefix             = "LProviderRecordRef/"
	// A separator inserted between keys.
)

// LProviderRecordKey returns the store key to retrieve a LProviderRecord
func LProviderRecordKey(
	poolName string,
	provider string,
) []byte {

	poolBytes := []byte(poolName)
	addrBytes := []byte(provider)

	return CombineKeys(poolBytes, addrBytes)
}

// Takes LProviderRecord struct to generate store key.
// Key format is: {poolName}{provider}
func GetProviderKey(record LProviderRecord) []byte {
	return LProviderRecordKey(record.PoolName, record.Provider)
}

// Takes LProviderRecord struct to generate reference key.
// Key format is: {provider}{provider}
func GetProviderRefKey(record LProviderRecord) []byte {

	poolBytes := []byte(record.PoolName)
	addrBytes := []byte(record.Provider)

	return CombineKeys(addrBytes, poolBytes)
}
