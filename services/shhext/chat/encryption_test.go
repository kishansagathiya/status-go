package chat

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

var cleartext = []byte("hello")

func TestEncryptionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EncryptionServiceTestSuite))
}

type EncryptionServiceTestSuite struct {
	suite.Suite
	alicedb *leveldb.DB
	bobdb   *leveldb.DB
	alice   *EncryptionService
	bob     *EncryptionService
}

func (s *EncryptionServiceTestSuite) SetupTest() {
	alicedb, err := leveldb.Open(storage.NewMemStorage(), nil)

	if err != nil {
		panic(err)
	}
	bobdb, err := leveldb.Open(storage.NewMemStorage(), nil)

	if err != nil {
		panic(err)
	}

	s.alicedb = alicedb
	s.bobdb = bobdb
	s.alice = NewEncryptionService(NewPersistenceService(alicedb))
	s.bob = NewEncryptionService(NewPersistenceService(bobdb))
}

func (s *EncryptionServiceTestSuite) TearDownTest() {
	s.NoError(s.alicedb.Close())
	s.NoError(s.bobdb.Close())
}

// Alice sends Bob an encrypted message with DH using an ephemeral key
// and Bob's identity key.
// Bob is able to decrypt it.
// Alice does not re-use the symmetric key
func (s *EncryptionServiceTestSuite) testEncryptPayloadNoBundle() {
	bobKey, err := crypto.GenerateKey()
	s.NoError(err)
	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)

	cyphertext1, ephemeralKey1, err := s.alice.EncryptPayload(&bobKey.PublicKey, aliceKey, cleartext)
	s.NotNil(ephemeralKey1, "It generates an ephemeral key for DH exchange")
	s.NotNil(cyphertext1, "It generates an encrypted payload")
	s.NotEqualf(cyphertext1, cleartext, "It encrypts the payload correctly")

	// On the receiver side, we should be able to decrypt using our private key and the ephemeral just sent
	decryptedPayload1, err := s.bob.DecryptWithDH(bobKey, ephemeralKey1, cyphertext1)
	s.NoError(err)
	s.Equalf(cleartext, decryptedPayload1, "It correctly decrypts the payload using DH")

	// The next message will not be re-using the same key
	cyphertext2, ephemeralKey2, err := s.alice.EncryptPayload(&bobKey.PublicKey, aliceKey, cleartext)
	s.NoError(err)

	s.NotEqual(cyphertext1, cyphertext2, "It does not re-use the symmetric key")
	s.NotEqual(ephemeralKey1, ephemeralKey2, "It does not re-use the ephemeral key")

	decryptedPayload2, err := s.bob.DecryptWithDH(bobKey, ephemeralKey2, cyphertext2)
	s.NoError(err)

	s.Equalf(cleartext, decryptedPayload2, "It correctly decrypts the payload using DH")
}

// Alice has Bob's bundle
// Alice sends Bob an encrypted message with X3DH using an ephemeral key
// and Bob's bundle.
func (s *EncryptionServiceTestSuite) TestEncryptPayloadBundle() {
	bobKey, err := crypto.GenerateKey()
	s.NoError(err)
	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)

	bobBundle, err := s.bob.CreateBundle(bobKey)
	s.NoError(err)

	// We add bob bundle
	err = s.alice.persistence.AddPublicBundle(bobBundle)
	s.NoError(err)

	// We send a message using the bundle
	cyphertext1, ephemeralKey1, err := s.alice.EncryptPayload(&bobKey.PublicKey, aliceKey, cleartext)
	s.NoError(err)
	s.NotNil(cyphertext1, "It generates an encrypted payload")
	s.NotEqualf(cyphertext1, cleartext, "It encrypts the payload correctly")
	s.NotNil(ephemeralKey1, "It generates an ephemeral key")

	// Bob is able to decrypt it using the bundle

	bundleId := bobBundle.GetSignedPreKey()

	decryptedPayload1, err := s.bob.DecryptWithX3DH(bobKey, &aliceKey.PublicKey, ephemeralKey1, bundleId, cyphertext1)
	s.NoError(err)
	s.Equalf(cleartext, decryptedPayload1, "It correctly decrypts the payload using X3DH")

	// Alice sends another message, this time she will use the same key as generated previously
	cyphertext2, ephemeralKey2, err := s.alice.EncryptPayload(&bobKey.PublicKey, aliceKey, cleartext)
	s.NoError(err)
	s.NotNil(cyphertext2, "It generates an encrypted payload")
	s.NotEqualf(cyphertext2, cleartext, "It encrypts the payload correctly")
	s.Equal(ephemeralKey1, ephemeralKey2, "It returns the same ephemeral key")

	// Bob this time should be able to decrypt it with a symmetric key
	decryptedPayload2, err := s.bob.DecryptSymmetricPayload(&aliceKey.PublicKey, bundleId, cyphertext2)
	s.NoError(err)
	s.Equalf(cleartext, decryptedPayload2, "It correctly decrypts the payload using X3DH")
}
