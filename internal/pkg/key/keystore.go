package key

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"

	"github.com/mr-tron/base58"
)

//const DefaultKey = "cc20530eb755388b7678ca721ce456b0dbe5418abf58953bd817d2c2b89c4593"

type Store struct {
	key *ecies.PrivateKey
}

func NewStore(key string) *Store {
	priv, _ := crypto.HexToECDSA(key)
	return &Store{
		key: ecies.ImportECDSA(priv),
	}
}

func (s *Store) Decode(encodePriv string) (*ecdsa.PrivateKey, error) {
	m, err := base58.Decode(encodePriv)
	if err != nil {
		return nil, err
	}
	decrypt, err := s.key.Decrypt(m, nil, nil)
	if err != nil {
		return nil, err
	}
	m2 := hexutil.Encode(decrypt)
	if priv, err := crypto.HexToECDSA(m2[2:]); err == nil {
		return priv, nil
	} else {
		return nil, err
	}
}

func (s *Store) Encode(privateKey *ecdsa.PrivateKey) (string, error) {
	m := crypto.FromECDSA(privateKey)
	encrypt, err := ecies.Encrypt(rand.Reader, &s.key.PublicKey, m, nil, nil)
	if err != nil {
		return "", err
	}
	return base58.Encode(encrypt), nil
}

func (s *Store) EncodePrivKey(rawPriv string) (string, error) {
	m, err := hexutil.Decode(rawPriv)
	if err != nil {
		return "", nil
	}
	encrypt, err := ecies.Encrypt(rand.Reader, &s.key.PublicKey, m, nil, nil)
	if err != nil {
		return "", err
	}
	return base58.Encode(encrypt), nil
}
