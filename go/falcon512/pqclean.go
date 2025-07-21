package falcon512

import (
	"crypto/rand"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"
)

const NONCELEN = 40

func PQCLEAN_FALCON512_CLEAN_crypto_sign_keypair(pk *uint8, sk *uint8) int {
	var (
		tmp struct {
			// union
			b         [14336]uint8
			dummy_u64 uint64
			dummy_fpr fpr
		}
		f    [512]int8
		g    [512]int8
		F    [512]int8
		h    [512]uint16
		seed [48]uint8
		rng  shake256incctx
		u    uint64
		v    uint64
	)
	PQCLEAN_randombytes(&seed[0], uint64(48))
	shake256_inc_init(&rng)
	shake256_inc_absorb(&rng, &seed[0], uint64(48))
	shake256_inc_finalize(&rng)
	PQCLEAN_FALCON512_CLEAN_keygen(&rng, &f[0], &g[0], &F[0], nil, &h[0], 9, &tmp.b[0])
	shake256_inc_ctx_release(&rng)
	*sk = 0x50 + 9
	u = 1
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_encode(unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u, &f[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_fg_bits[9]))
	if v == 0 {
		return -1
	}
	u += v
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_encode(unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u, &g[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_fg_bits[9]))
	if v == 0 {
		return -1
	}
	u += v
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_encode(unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u, &F[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_FG_bits[9]))
	if v == 0 {
		return -1
	}
	u += v
	if u != PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES {
		return -1
	}
	*pk = 0x0 + 9
	v = PQCLEAN_FALCON512_CLEAN_modq_encode(unsafe.Add(unsafe.Pointer(pk), 1), uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_PUBLICKEYBYTES-1)), &h[0], 9)
	if v != uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_PUBLICKEYBYTES-1)) {
		return -1
	}
	return 0
}
func do_sign(nonce *uint8, sigbuf *uint8, sigbuflen *uint64, m *uint8, mlen uint64, sk *uint8) int {
	var (
		tmp struct {
			// union
			b         [36864]uint8
			dummy_u64 uint64
			dummy_fpr fpr
		}
		f [512]int8
		g [512]int8
		F [512]int8
		G [512]int8
		r struct {
			sig [512]int16
			hm  [512]uint16
		}
		seed [48]uint8
		sc   shake256incctx
		u    uint64
		v    uint64
	)
	if int(*sk) != 0x50+9 {
		return -1
	}
	u = 1
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_decode(&f[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_fg_bits[9]), unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u)
	if v == 0 {
		return -1
	}
	u += v
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_decode(&g[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_fg_bits[9]), unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u)
	if v == 0 {
		return -1
	}
	u += v
	v = PQCLEAN_FALCON512_CLEAN_trim_i8_decode(&F[0], 9, uint(PQCLEAN_FALCON512_CLEAN_max_FG_bits[9]), unsafe.Add(unsafe.Pointer(sk), u), PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES-u)
	if v == 0 {
		return -1
	}
	u += v
	if u != PQCLEAN_FALCON512_CLEAN_CRYPTO_SECRETKEYBYTES {
		return -1
	}
	if PQCLEAN_FALCON512_CLEAN_complete_private(&G[0], &f[0], &g[0], &F[0], 9, &tmp.b[0]) == 0 {
		return -1
	}
	PQCLEAN_randombytes(nonce, NONCELEN)
	shake256_inc_init(&sc)
	shake256_inc_absorb(&sc, nonce, NONCELEN)
	shake256_inc_absorb(&sc, m, mlen)
	shake256_inc_finalize(&sc)
	PQCLEAN_FALCON512_CLEAN_hash_to_point_ct(&sc, &r.hm[0], 9, &tmp.b[0])
	shake256_inc_ctx_release(&sc)
	PQCLEAN_randombytes(&seed[0], uint64(48))
	shake256_inc_init(&sc)
	shake256_inc_absorb(&sc, &seed[0], uint64(48))
	shake256_inc_finalize(&sc)
	PQCLEAN_FALCON512_CLEAN_sign_dyn(&r.sig[0], &sc, &f[0], &g[0], &F[0], &G[0], &r.hm[0], 9, &tmp.b[0])
	v = PQCLEAN_FALCON512_CLEAN_comp_encode(unsafe.Pointer(sigbuf), *sigbuflen, &r.sig[0], 9)
	if v != 0 {
		shake256_inc_ctx_release(&sc)
		*sigbuflen = v
		return 0
	}
	return -1
}
func do_verify(nonce *uint8, sigbuf *uint8, sigbuflen uint64, m *uint8, mlen uint64, pk *uint8) int {
	var (
		tmp struct {
			// union
			b         [1024]uint8
			dummy_u64 uint64
			dummy_fpr fpr
		}
		h   [512]uint16
		hm  [512]uint16
		sig [512]int16
		sc  shake256incctx
		v   uint64
	)
	if int(*pk) != 0x0+9 {
		return -1
	}
	if PQCLEAN_FALCON512_CLEAN_modq_decode(&h[0], 9, unsafe.Add(unsafe.Pointer(pk), 1), uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_PUBLICKEYBYTES-1))) != uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_PUBLICKEYBYTES-1)) {
		return -1
	}
	PQCLEAN_FALCON512_CLEAN_to_ntt_monty(&h[0], 9)
	if sigbuflen == 0 {
		return -1
	}
	v = PQCLEAN_FALCON512_CLEAN_comp_decode(&sig[0], 9, unsafe.Pointer(sigbuf), sigbuflen)
	if v == 0 {
		return -1
	}
	if v != sigbuflen {
		if sigbuflen == uint64(int(PQCLEAN_FALCONPADDED512_CLEAN_CRYPTO_BYTES-NONCELEN)-1) {
			for v < sigbuflen {
				if int(*(*uint8)(unsafe.Add(unsafe.Pointer(sigbuf), func() uint64 {
					p_ := &v
					x := *p_
					*p_++
					return x
				}()))) != 0 {
					return -1
				}
			}
		} else {
			return -1
		}
	}
	shake256_inc_init(&sc)
	shake256_inc_absorb(&sc, nonce, NONCELEN)
	shake256_inc_absorb(&sc, m, mlen)
	shake256_inc_finalize(&sc)
	PQCLEAN_FALCON512_CLEAN_hash_to_point_ct(&sc, &hm[0], 9, &tmp.b[0])
	shake256_inc_ctx_release(&sc)
	if PQCLEAN_FALCON512_CLEAN_verify_raw(&hm[0], &sig[0], &h[0], 9, &tmp.b[0]) == 0 {
		return -1
	}
	return 0
}
func PQCLEAN_FALCON512_CLEAN_crypto_sign_signature(sig *uint8, siglen *uint64, m *uint8, mlen uint64, sk *uint8) int {
	var vlen uint64
	vlen = uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_BYTES-NONCELEN) - 1)
	if do_sign((*uint8)(unsafe.Add(unsafe.Pointer(sig), 1)), (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sig), 1))), NONCELEN)), &vlen, m, mlen, sk) < 0 {
		return -1
	}
	*sig = 0x30 + 9
	*siglen = uint64(int(NONCELEN+1) + int(vlen))
	return 0
}
func PQCLEAN_FALCON512_CLEAN_crypto_sign_verify(sig *uint8, siglen uint64, m *uint8, mlen uint64, pk *uint8) int {
	if siglen < uint64(int(NONCELEN+1)) {
		return -1
	}
	if int(*sig) != 0x30+9 {
		return -1
	}
	return do_verify((*uint8)(unsafe.Add(unsafe.Pointer(sig), 1)), (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sig), 1))), NONCELEN)), siglen-1-NONCELEN, m, mlen, pk)
}
func PQCLEAN_FALCON512_CLEAN_crypto_sign(sm *uint8, smlen *uint64, m *uint8, mlen uint64, sk *uint8) int {
	var (
		pm        *uint8
		sigbuf    *uint8
		sigbuflen uint64
	)
	libc.MemMove(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2))), NONCELEN), unsafe.Pointer(m), int(mlen))
	pm = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2))), NONCELEN))
	sigbuf = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(pm), 1))), mlen))
	sigbuflen = uint64(int(PQCLEAN_FALCON512_CLEAN_CRYPTO_BYTES-NONCELEN) - 3)
	if do_sign((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2)), sigbuf, &sigbuflen, pm, mlen, sk) < 0 {
		return -1
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer(pm), mlen)) = 0x20 + 9
	sigbuflen++
	*sm = uint8(sigbuflen >> 8)
	*(*uint8)(unsafe.Add(unsafe.Pointer(sm), 1)) = uint8(sigbuflen)
	*smlen = mlen + 2 + NONCELEN + sigbuflen
	return 0
}
func PQCLEAN_FALCON512_CLEAN_crypto_sign_open(m *uint8, mlen *uint64, sm *uint8, smlen uint64, pk *uint8) int {
	var (
		sigbuf    *uint8
		pmlen     uint64
		sigbuflen uint64
	)
	if smlen < uint64(int(NONCELEN+3)) {
		return -1
	}
	sigbuflen = (uint64(*sm) << 8) | uint64(*(*uint8)(unsafe.Add(unsafe.Pointer(sm), 1)))
	if sigbuflen < 2 || sigbuflen > (smlen-NONCELEN-2) {
		return -1
	}
	sigbuflen--
	pmlen = smlen - NONCELEN - 3 - sigbuflen
	if int(*(*uint8)(unsafe.Add(unsafe.Pointer(sm), int(NONCELEN+2)+int(pmlen)))) != 0x20+9 {
		return -1
	}
	sigbuf = (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2))), NONCELEN))), pmlen))), 1))
	if do_verify((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2)), sigbuf, sigbuflen, (*uint8)(unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2))), NONCELEN)), pmlen, pk) < 0 {
		return -1
	}
	libc.MemMove(unsafe.Pointer(m), unsafe.Add(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(sm), 2))), NONCELEN), int(pmlen))
	*mlen = pmlen
	return 0
}

func PQCLEAN_randombytes(output *uint8, n uint64) int {
	buf := unsafe.Slice(output, n)
	if _, err := rand.Read(buf); err != nil {
		return -1
	}
	return 0
}
