package falcon512

import (
	"bytes"
	"crypto/rand"
	"testing"
)

const (
	CRYPTO_PUBLICKEYBYTES = 897
	CRYPTO_SECRETKEYBYTES = 1281
	CRYPTO_BYTES          = 752
)

func TestFalcon512(t *testing.T) {
	var (
		pk [CRYPTO_PUBLICKEYBYTES]byte
		sk [CRYPTO_SECRETKEYBYTES]byte
		m  [33]byte
		sm [33 + CRYPTO_BYTES]byte
	)

	if _, err := rand.Read(m[:]); err != nil {
		t.Fatalf("rand.Read failed: %v", err)
	}

	if ret := PQCLEAN_FALCON512_CLEAN_crypto_sign_keypair(&pk[0], &sk[0]); ret != 0 {
		t.Fatalf("crypto_sign_keypair returned %d", ret)
	}

	var smlen uint64
	if ret := PQCLEAN_FALCON512_CLEAN_crypto_sign(&sm[0], &smlen, &m[0], uint64(len(m)), &sk[0]); ret != 0 {
		t.Fatalf("crypto_sign returned %d", ret)
	}

	var mlen1 uint64
	if ret := PQCLEAN_FALCON512_CLEAN_crypto_sign_open(&m[0], &mlen1, &sm[0], smlen, &pk[0]); ret != 0 {
		t.Fatalf("crypto_sign_open returned %d", ret)
	}

	if mlen1 != uint64(len(m)) {
		t.Fatalf("crypto_sign_open returned bad 'mlen': got %d, expected %d", mlen1, len(m))
	}

	if !bytes.Equal(m[:], sm[:len(m)]) {
		t.Fatal("crypto_sign_open returned bad 'm' value")
	}
}
