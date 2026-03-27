package core

import (
	"github.com/TerraDharitri/drt-go-chain-core/core/check"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
	"github.com/TerraDharitri/drt-go-sdk/blockchain/cryptoProvider"
	"github.com/TerraDharitri/drt-go-sdk/core"
)

// CryptoComponentsHolderFactory is the implementation of the CryptoComponentsHolderFactory interface
type CryptoComponentsHolderFactory struct {
	keyGen crypto.KeyGenerator
}

// NewCryptoComponentsHolderFactory creates a new instance of CryptoComponentsHolderFactory
func NewCryptoComponentsHolderFactory(keyGen crypto.KeyGenerator) (*CryptoComponentsHolderFactory, error) {
	if check.IfNil(keyGen) {
		return nil, ErrNilKeyGenerator
	}

	return &CryptoComponentsHolderFactory{
		keyGen: keyGen,
	}, nil
}

// Create creates a new instance of CryptoComponentsHolder
func (f *CryptoComponentsHolderFactory) Create(skBytes []byte) (core.CryptoComponentsHolder, error) {
	return cryptoProvider.NewCryptoComponentsHolder(f.keyGen, skBytes)
}

// IsInterfaceNil returns true if there is no value under the interface
func (f *CryptoComponentsHolderFactory) IsInterfaceNil() bool {
	return f == nil
}
