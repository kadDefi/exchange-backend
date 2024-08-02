package key

import (
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/mr-tron/base58"
)

func TestKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	k := crypto.FromECDSA(privateKey)
	t.Log(hexutil.Encode(k))

	eciePriv := ecies.ImportECDSA(privateKey)
	priv2, _ := crypto.GenerateKey()
	m := crypto.FromECDSA(priv2)
	t.Log(hexutil.Encode(m))
	encrypt, err := ecies.Encrypt(rand.Reader, &eciePriv.PublicKey, m, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	m1 := base58.Encode(encrypt)
	t.Log(m1)

	decode, err := base58.Decode(m1)
	if err != nil {
		t.Fatal(err)
	}
	decrypt, err := eciePriv.Decrypt(decode, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	m2 := hexutil.Encode(decrypt)
	t.Log(m2)

	priv3, _ := crypto.HexToECDSA(m2[2:])
	t.Log(priv3)
}

func TestEncrypt(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	k := crypto.FromECDSA(privateKey)
	strKey := hexutil.Encode(k)
	t.Log(strKey)

	store := NewStore("0c20e7bca884cabbf7e0c5484866d96e45384671e9bc1613027d63026465fbbc")
	key, err := store.EncodePrivKey(strKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(key)

	decode, err := store.Decode(key)
	if err != nil {
		t.Fatal(err)
	}
	k2 := crypto.FromECDSA(decode)
	strKey2 := hexutil.Encode(k2)
	if strKey2 != strKey {
		t.Fatal("invalid key")
	}

	encode, err := store.Encode(privateKey)

	decode2, err := store.Decode(encode)
	if err != nil {
		t.Fatal(err)
	}
	k3 := crypto.FromECDSA(decode2)
	strKey3 := hexutil.Encode(k3)
	if strKey3 != strKey {
		t.Fatal("invalid key")
	}
}

func TestStore_Factor(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	k := crypto.FromECDSA(privateKey)
	s:=hexutil.Encode(k)
	s1:=s[2:]
	t.Log(s1)
}
