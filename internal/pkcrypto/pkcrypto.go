package pkcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

const (
	keySize   = 32
	envKeyVar = "PKEYS_KEY"
)

// Generate key for encryption
func Generate() (string, error) {
	// Create empty byte slice
	key := make([]byte, keySize)

	// Read bytes from the OS entropy source into the key byte slice
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", fmt.Errorf("failed to generate random key: %w", err)
	}

	// Encode key to hex and return
	return hex.EncodeToString(key), nil
}

// Encrypt data
func Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	// Define AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Wrap block for GCM mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Get nonce for encryption of size aesgcm.NonceSize() (12 bytes)
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Seal plaintext
	// nonce is passed first as the destintaion, then as the nonce for appending
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aesgcm.NonceSize() {
		return nil, fmt.Errorf("ciphertext too short: must be at least %d bytes, got %d", aesgcm.NonceSize(), len(ciphertext))
	}

	// Split nonce from ciphertext
	nonce, ciphertext := ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// Get key
func GetEnvKey() ([]byte, error) {
	keyString := os.Getenv(envKeyVar)
	if keyString == "" {
		return nil, fmt.Errorf("enviroment key not set - set `PKEYS_KEY` system variable")
	}

	// Decode key into proper format
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return nil, fmt.Errorf("key '%v' is not a valid hex string", err)
	}

	// Verify key length
	if len(key) != keySize {
		return nil, fmt.Errorf("invalid key size - must be %d bytes, but got %d", keySize, len(key))
	}

	return key, nil
}
