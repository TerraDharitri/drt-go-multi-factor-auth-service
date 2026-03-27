package testscommon

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
	"github.com/TerraDharitri/drt-go-sdk/core"
)

// GuardedTxBuilderStub -
type GuardedTxBuilderStub struct {
	ApplyGuardianSignatureCalled func(cryptoHolderGuardian core.CryptoComponentsHolder, tx *transaction.FrontendTransaction) error
}

// ApplyGuardianSignature -
func (stub *GuardedTxBuilderStub) ApplyGuardianSignature(cryptoHolderGuardian core.CryptoComponentsHolder, tx *transaction.FrontendTransaction) error {
	if stub.ApplyGuardianSignatureCalled != nil {
		return stub.ApplyGuardianSignatureCalled(cryptoHolderGuardian, tx)
	}
	return nil
}

// IsInterfaceNil -
func (stub *GuardedTxBuilderStub) IsInterfaceNil() bool {
	return stub == nil
}
