package factory

import (
	"fmt"

	"github.com/TerraDharitri/drt-go-chain-core/core/pubkeyConverter"
	crypto "github.com/TerraDharitri/drt-go-chain-crypto"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing"
	"github.com/TerraDharitri/drt-go-chain-crypto/signing/ed25519"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/config"
	"github.com/TerraDharitri/drt-go-multi-factor-auth-service/core"
	"github.com/TerraDharitri/drt-go-sdk/blockchain/cryptoProvider"
	"github.com/TerraDharitri/drt-go-sdk/builders"
)

const bech32Format = "bech32"

// cryptoComponentsHolder will hold core crypto components
type cryptoComponentsHolder struct {
	keyGenerator    crypto.KeyGenerator
	signer          builders.Signer
	pubKeyConverter core.PubkeyConverter
}

// CreateCoreCryptoComponents will create core crypto components
func CreateCoreCryptoComponents(conf config.PubkeyConfig) (*cryptoComponentsHolder, error) {
	pkConv, err := createPubkeyConverter(conf)
	if err != nil {
		return nil, err
	}

	keyGen := signing.NewKeyGenerator(ed25519.NewEd25519())
	signer := cryptoProvider.NewSigner()

	return &cryptoComponentsHolder{
		keyGenerator:    keyGen,
		signer:          signer,
		pubKeyConverter: pkConv,
	}, nil
}

func createPubkeyConverter(config config.PubkeyConfig) (core.PubkeyConverter, error) {
	switch config.Type {
	case bech32Format:
		return pubkeyConverter.NewBech32PubkeyConverter(config.Length, config.Hrp)
	default:
		return nil, fmt.Errorf("%w unrecognized type %s", core.ErrInvalidPubkeyConverterType, config.Type)
	}
}

// KeyGenerator returns key generator component
func (cch *cryptoComponentsHolder) KeyGenerator() crypto.KeyGenerator {
	return cch.keyGenerator
}

// Signer returns signer component
func (cch *cryptoComponentsHolder) Signer() builders.Signer {
	return cch.signer
}

// PubkeyConverter returns pubkey converter component
func (cch *cryptoComponentsHolder) PubkeyConverter() core.PubkeyConverter {
	return cch.pubKeyConverter
}
