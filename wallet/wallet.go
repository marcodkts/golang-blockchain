package wallet

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/marcodkts/golang-blockchain/handle"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey []byte
	PublicKey  []byte
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)

	address := Base58Encode(fullHash)

	return address
}

func NewKeyPair() ([]byte, []byte) {
	curve := elliptic.P256()

	private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	handle.Handle(err)

	pub := append(x.Bytes(), y.Bytes()...)

	return private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := sha256.New()
	_, err := hasher.Write(pubHash[:])
	handle.Handle(err)

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
