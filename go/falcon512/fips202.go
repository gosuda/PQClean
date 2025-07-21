package falcon512

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"os"
	"unsafe"
)

const SHAKE128_RATE = 168
const SHAKE256_RATE = 136
const SHA3_256_RATE = 136
const SHA3_384_RATE = 104
const SHA3_512_RATE = 72
const NROUNDS = 24

type shake128incctx struct {
	ctx *uint64
}
type shake128ctx struct {
	ctx *uint64
}
type shake256incctx struct {
	ctx *uint64
}
type shake256ctx struct {
	ctx *uint64
}
type sha3_256incctx struct {
	ctx *uint64
}
type sha3_384incctx struct {
	ctx *uint64
}
type sha3_512incctx struct {
	ctx *uint64
}

func load64(x *uint8) uint64 {
	var r uint64 = 0
	for i := uint64(0); i < 8; i++ {
		r |= uint64(*(*uint8)(unsafe.Add(unsafe.Pointer(x), i))) << (i * 8)
	}
	return r
}
func store64(x *uint8, u uint64) {
	for i := uint64(0); i < 8; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(x), i)) = uint8(u >> (i * 8))
	}
}

var KeccakF_RoundConstants [24]uint64 = [24]uint64{0x1, 0x8082, 0x800000000000808A, 0x8000000080008000, 0x808B, 0x80000001, 0x8000000080008081, 0x8000000000008009, 0x8A, 0x88, 0x80008009, 0x8000000A, 0x8000808B, 0x800000000000008B, 0x8000000000008089, 0x8000000000008003, 0x8000000000008002, 0x8000000000000080, 0x800A, 0x800000008000000A, 0x8000000080008081, 0x8000000000008080, 0x80000001, 0x8000000080008008}

func KeccakF1600_StatePermute(state *uint64) {
	var (
		round int
		Aba   uint64
		Abe   uint64
		Abi   uint64
		Abo   uint64
		Abu   uint64
		Aga   uint64
		Age   uint64
		Agi   uint64
		Ago   uint64
		Agu   uint64
		Aka   uint64
		Ake   uint64
		Aki   uint64
		Ako   uint64
		Aku   uint64
		Ama   uint64
		Ame   uint64
		Ami   uint64
		Amo   uint64
		Amu   uint64
		Asa   uint64
		Ase   uint64
		Asi   uint64
		Aso   uint64
		Asu   uint64
		BCa   uint64
		BCe   uint64
		BCi   uint64
		BCo   uint64
		BCu   uint64
		Da    uint64
		De    uint64
		Di    uint64
		Do    uint64
		Du    uint64
		Eba   uint64
		Ebe   uint64
		Ebi   uint64
		Ebo   uint64
		Ebu   uint64
		Ega   uint64
		Ege   uint64
		Egi   uint64
		Ego   uint64
		Egu   uint64
		Eka   uint64
		Eke   uint64
		Eki   uint64
		Eko   uint64
		Eku   uint64
		Ema   uint64
		Eme   uint64
		Emi   uint64
		Emo   uint64
		Emu   uint64
		Esa   uint64
		Ese   uint64
		Esi   uint64
		Eso   uint64
		Esu   uint64
	)
	Aba = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*0))
	Abe = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*1))
	Abi = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*2))
	Abo = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*3))
	Abu = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*4))
	Aga = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*5))
	Age = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*6))
	Agi = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*7))
	Ago = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*8))
	Agu = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*9))
	Aka = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*10))
	Ake = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*11))
	Aki = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*12))
	Ako = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*13))
	Aku = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*14))
	Ama = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*15))
	Ame = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*16))
	Ami = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*17))
	Amo = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*18))
	Amu = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*19))
	Asa = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*20))
	Ase = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*21))
	Asi = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*22))
	Aso = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*23))
	Asu = *(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*24))
	for round = 0; round < NROUNDS; round += 2 {
		BCa = Aba ^ Aga ^ Aka ^ Ama ^ Asa
		BCe = Abe ^ Age ^ Ake ^ Ame ^ Ase
		BCi = Abi ^ Agi ^ Aki ^ Ami ^ Asi
		BCo = Abo ^ Ago ^ Ako ^ Amo ^ Aso
		BCu = Abu ^ Agu ^ Aku ^ Amu ^ Asu
		Da = BCu ^ ((BCe << 1) ^ BCe>>(64-1))
		De = BCa ^ ((BCi << 1) ^ BCi>>(64-1))
		Di = BCe ^ ((BCo << 1) ^ BCo>>(64-1))
		Do = BCi ^ ((BCu << 1) ^ BCu>>(64-1))
		Du = BCo ^ ((BCa << 1) ^ BCa>>(64-1))
		Aba ^= Da
		BCa = Aba
		Age ^= De
		BCe = (Age << 44) ^ Age>>(64-44)
		Aki ^= Di
		BCi = (Aki << 43) ^ Aki>>(64-43)
		Amo ^= Do
		BCo = (Amo << 21) ^ Amo>>(64-21)
		Asu ^= Du
		BCu = (Asu << 14) ^ Asu>>(64-14)
		Eba = BCa ^ ((^BCe) & BCi)
		Eba ^= KeccakF_RoundConstants[round]
		Ebe = BCe ^ ((^BCi) & BCo)
		Ebi = BCi ^ ((^BCo) & BCu)
		Ebo = BCo ^ ((^BCu) & BCa)
		Ebu = BCu ^ ((^BCa) & BCe)
		Abo ^= Do
		BCa = (Abo << 28) ^ Abo>>(64-28)
		Agu ^= Du
		BCe = (Agu << 20) ^ Agu>>(64-20)
		Aka ^= Da
		BCi = (Aka << 3) ^ Aka>>(64-3)
		Ame ^= De
		BCo = (Ame << 45) ^ Ame>>(64-45)
		Asi ^= Di
		BCu = (Asi << 61) ^ Asi>>(64-61)
		Ega = BCa ^ ((^BCe) & BCi)
		Ege = BCe ^ ((^BCi) & BCo)
		Egi = BCi ^ ((^BCo) & BCu)
		Ego = BCo ^ ((^BCu) & BCa)
		Egu = BCu ^ ((^BCa) & BCe)
		Abe ^= De
		BCa = (Abe << 1) ^ Abe>>(64-1)
		Agi ^= Di
		BCe = (Agi << 6) ^ Agi>>(64-6)
		Ako ^= Do
		BCi = (Ako << 25) ^ Ako>>(64-25)
		Amu ^= Du
		BCo = (Amu << 8) ^ Amu>>(64-8)
		Asa ^= Da
		BCu = (Asa << 18) ^ Asa>>(64-18)
		Eka = BCa ^ ((^BCe) & BCi)
		Eke = BCe ^ ((^BCi) & BCo)
		Eki = BCi ^ ((^BCo) & BCu)
		Eko = BCo ^ ((^BCu) & BCa)
		Eku = BCu ^ ((^BCa) & BCe)
		Abu ^= Du
		BCa = (Abu << 27) ^ Abu>>(64-27)
		Aga ^= Da
		BCe = (Aga << 36) ^ Aga>>(64-36)
		Ake ^= De
		BCi = (Ake << 10) ^ Ake>>(64-10)
		Ami ^= Di
		BCo = (Ami << 15) ^ Ami>>(64-15)
		Aso ^= Do
		BCu = (Aso << 56) ^ Aso>>(64-56)
		Ema = BCa ^ ((^BCe) & BCi)
		Eme = BCe ^ ((^BCi) & BCo)
		Emi = BCi ^ ((^BCo) & BCu)
		Emo = BCo ^ ((^BCu) & BCa)
		Emu = BCu ^ ((^BCa) & BCe)
		Abi ^= Di
		BCa = (Abi << 62) ^ Abi>>(64-62)
		Ago ^= Do
		BCe = (Ago << 55) ^ Ago>>(64-55)
		Aku ^= Du
		BCi = (Aku << 39) ^ Aku>>(64-39)
		Ama ^= Da
		BCo = (Ama << 41) ^ Ama>>(64-41)
		Ase ^= De
		BCu = (Ase << 2) ^ Ase>>(64-2)
		Esa = BCa ^ ((^BCe) & BCi)
		Ese = BCe ^ ((^BCi) & BCo)
		Esi = BCi ^ ((^BCo) & BCu)
		Eso = BCo ^ ((^BCu) & BCa)
		Esu = BCu ^ ((^BCa) & BCe)
		BCa = Eba ^ Ega ^ Eka ^ Ema ^ Esa
		BCe = Ebe ^ Ege ^ Eke ^ Eme ^ Ese
		BCi = Ebi ^ Egi ^ Eki ^ Emi ^ Esi
		BCo = Ebo ^ Ego ^ Eko ^ Emo ^ Eso
		BCu = Ebu ^ Egu ^ Eku ^ Emu ^ Esu
		Da = BCu ^ ((BCe << 1) ^ BCe>>(64-1))
		De = BCa ^ ((BCi << 1) ^ BCi>>(64-1))
		Di = BCe ^ ((BCo << 1) ^ BCo>>(64-1))
		Do = BCi ^ ((BCu << 1) ^ BCu>>(64-1))
		Du = BCo ^ ((BCa << 1) ^ BCa>>(64-1))
		Eba ^= Da
		BCa = Eba
		Ege ^= De
		BCe = (Ege << 44) ^ Ege>>(64-44)
		Eki ^= Di
		BCi = (Eki << 43) ^ Eki>>(64-43)
		Emo ^= Do
		BCo = (Emo << 21) ^ Emo>>(64-21)
		Esu ^= Du
		BCu = (Esu << 14) ^ Esu>>(64-14)
		Aba = BCa ^ ((^BCe) & BCi)
		Aba ^= KeccakF_RoundConstants[round+1]
		Abe = BCe ^ ((^BCi) & BCo)
		Abi = BCi ^ ((^BCo) & BCu)
		Abo = BCo ^ ((^BCu) & BCa)
		Abu = BCu ^ ((^BCa) & BCe)
		Ebo ^= Do
		BCa = (Ebo << 28) ^ Ebo>>(64-28)
		Egu ^= Du
		BCe = (Egu << 20) ^ Egu>>(64-20)
		Eka ^= Da
		BCi = (Eka << 3) ^ Eka>>(64-3)
		Eme ^= De
		BCo = (Eme << 45) ^ Eme>>(64-45)
		Esi ^= Di
		BCu = (Esi << 61) ^ Esi>>(64-61)
		Aga = BCa ^ ((^BCe) & BCi)
		Age = BCe ^ ((^BCi) & BCo)
		Agi = BCi ^ ((^BCo) & BCu)
		Ago = BCo ^ ((^BCu) & BCa)
		Agu = BCu ^ ((^BCa) & BCe)
		Ebe ^= De
		BCa = (Ebe << 1) ^ Ebe>>(64-1)
		Egi ^= Di
		BCe = (Egi << 6) ^ Egi>>(64-6)
		Eko ^= Do
		BCi = (Eko << 25) ^ Eko>>(64-25)
		Emu ^= Du
		BCo = (Emu << 8) ^ Emu>>(64-8)
		Esa ^= Da
		BCu = (Esa << 18) ^ Esa>>(64-18)
		Aka = BCa ^ ((^BCe) & BCi)
		Ake = BCe ^ ((^BCi) & BCo)
		Aki = BCi ^ ((^BCo) & BCu)
		Ako = BCo ^ ((^BCu) & BCa)
		Aku = BCu ^ ((^BCa) & BCe)
		Ebu ^= Du
		BCa = (Ebu << 27) ^ Ebu>>(64-27)
		Ega ^= Da
		BCe = (Ega << 36) ^ Ega>>(64-36)
		Eke ^= De
		BCi = (Eke << 10) ^ Eke>>(64-10)
		Emi ^= Di
		BCo = (Emi << 15) ^ Emi>>(64-15)
		Eso ^= Do
		BCu = (Eso << 56) ^ Eso>>(64-56)
		Ama = BCa ^ ((^BCe) & BCi)
		Ame = BCe ^ ((^BCi) & BCo)
		Ami = BCi ^ ((^BCo) & BCu)
		Amo = BCo ^ ((^BCu) & BCa)
		Amu = BCu ^ ((^BCa) & BCe)
		Ebi ^= Di
		BCa = (Ebi << 62) ^ Ebi>>(64-62)
		Ego ^= Do
		BCe = (Ego << 55) ^ Ego>>(64-55)
		Eku ^= Du
		BCi = (Eku << 39) ^ Eku>>(64-39)
		Ema ^= Da
		BCo = (Ema << 41) ^ Ema>>(64-41)
		Ese ^= De
		BCu = (Ese << 2) ^ Ese>>(64-2)
		Asa = BCa ^ ((^BCe) & BCi)
		Ase = BCe ^ ((^BCi) & BCo)
		Asi = BCi ^ ((^BCo) & BCu)
		Aso = BCo ^ ((^BCu) & BCa)
		Asu = BCu ^ ((^BCa) & BCe)
	}
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*0)) = Aba
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*1)) = Abe
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*2)) = Abi
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*3)) = Abo
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*4)) = Abu
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*5)) = Aga
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*6)) = Age
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*7)) = Agi
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*8)) = Ago
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*9)) = Agu
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*10)) = Aka
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*11)) = Ake
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*12)) = Aki
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*13)) = Ako
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*14)) = Aku
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*15)) = Ama
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*16)) = Ame
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*17)) = Ami
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*18)) = Amo
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*19)) = Amu
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*20)) = Asa
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*21)) = Ase
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*22)) = Asi
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*23)) = Aso
	*(*uint64)(unsafe.Add(unsafe.Pointer(state), unsafe.Sizeof(uint64(0))*24)) = Asu
}
func keccak_absorb(s *uint64, r uint32, m *uint8, mlen uint64, p uint8) {
	var (
		i uint64
		t [200]uint8
	)
	for i = 0; i < 25; i++ {
		*(*uint64)(unsafe.Add(unsafe.Pointer(s), unsafe.Sizeof(uint64(0))*uintptr(i))) = 0
	}
	for mlen >= uint64(r) {
		for i = 0; i < uint64(int(r)/8); i++ {
			*(*uint64)(unsafe.Add(unsafe.Pointer(s), unsafe.Sizeof(uint64(0))*uintptr(i))) ^= load64((*uint8)(unsafe.Add(unsafe.Pointer(m), i*8)))
		}
		KeccakF1600_StatePermute(s)
		mlen -= uint64(r)
		m = (*uint8)(unsafe.Add(unsafe.Pointer(m), r))
	}
	for i = 0; i < uint64(r); i++ {
		t[i] = 0
	}
	for i = 0; i < mlen; i++ {
		t[i] = *(*uint8)(unsafe.Add(unsafe.Pointer(m), i))
	}
	t[i] = p
	t[int(r)-1] |= 128
	for i = 0; i < uint64(int(r)/8); i++ {
		*(*uint64)(unsafe.Add(unsafe.Pointer(s), unsafe.Sizeof(uint64(0))*uintptr(i))) ^= load64(&t[i*8])
	}
}
func keccak_squeezeblocks(h *uint8, nblocks uint64, s *uint64, r uint32) {
	for nblocks > 0 {
		KeccakF1600_StatePermute(s)
		for i := uint64(0); i < uint64(int(r)>>3); i++ {
			store64((*uint8)(unsafe.Add(unsafe.Pointer(h), i*8)), *(*uint64)(unsafe.Add(unsafe.Pointer(s), unsafe.Sizeof(uint64(0))*uintptr(i))))
		}
		h = (*uint8)(unsafe.Add(unsafe.Pointer(h), r))
		nblocks--
	}
}
func keccak_inc_init(s_inc *uint64) {
	var i uint64
	for i = 0; i < 25; i++ {
		*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr(i))) = 0
	}
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) = 0
}
func keccak_inc_absorb(s_inc *uint64, r uint32, m *uint8, mlen uint64) {
	var i uint64
	for mlen+*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) >= uint64(r) {
		for i = 0; i < uint64(int(r)-int(uint32(*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))))); i++ {
			*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr((*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))+i)>>3))) ^= uint64(*(*uint8)(unsafe.Add(unsafe.Pointer(m), i))) << (((*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) + i) & 0x7) * 8)
		}
		mlen -= uint64(r) - *(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))
		m = (*uint8)(unsafe.Add(unsafe.Pointer(m), uint64(r)-*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))))
		*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) = 0
		KeccakF1600_StatePermute(s_inc)
	}
	for i = 0; i < mlen; i++ {
		*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr((*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))+i)>>3))) ^= uint64(*(*uint8)(unsafe.Add(unsafe.Pointer(m), i))) << (((*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) + i) & 0x7) * 8)
	}
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) += mlen
}
func keccak_inc_finalize(s_inc *uint64, r uint32, p uint8) {
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr(*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))>>3))) ^= uint64(p) << ((*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) & 0x7) * 8)
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr((int(r)-1)>>3))) ^= uint64(128 << (((int(r) - 1) & 0x7) * 8))
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) = 0
}
func keccak_inc_squeeze(h *uint8, outlen uint64, s_inc *uint64, r uint32) {
	var i uint64
	for i = 0; i < outlen && i < *(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)); i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(h), i)) = uint8(*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr((uint64(r)-*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25))+i)>>3))) >> (((uint64(r) - *(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) + i) & 0x7) * 8))
	}
	h = (*uint8)(unsafe.Add(unsafe.Pointer(h), i))
	outlen -= i
	*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) -= i
	for outlen > 0 {
		KeccakF1600_StatePermute(s_inc)
		for i = 0; i < outlen && i < uint64(r); i++ {
			*(*uint8)(unsafe.Add(unsafe.Pointer(h), i)) = uint8(*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*uintptr(i>>3))) >> ((i & 0x7) * 8))
		}
		h = (*uint8)(unsafe.Add(unsafe.Pointer(h), i))
		outlen -= i
		*(*uint64)(unsafe.Add(unsafe.Pointer(s_inc), unsafe.Sizeof(uint64(0))*25)) = uint64(r) - i
	}
}
func shake128_inc_init(state *shake128incctx) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_inc_init(state.ctx)
}
func shake128_inc_absorb(state *shake128incctx, input *uint8, inlen uint64) {
	keccak_inc_absorb(state.ctx, SHAKE128_RATE, input, inlen)
}
func shake128_inc_finalize(state *shake128incctx) {
	keccak_inc_finalize(state.ctx, SHAKE128_RATE, 0x1F)
}
func shake128_inc_squeeze(output *uint8, outlen uint64, state *shake128incctx) {
	keccak_inc_squeeze(output, outlen, state.ctx, SHAKE128_RATE)
}
func shake128_inc_ctx_clone(dest *shake128incctx, src *shake128incctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*26))
}
func shake128_inc_ctx_release(state *shake128incctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func shake256_inc_init(state *shake256incctx) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_inc_init(state.ctx)
}
func shake256_inc_absorb(state *shake256incctx, input *uint8, inlen uint64) {
	keccak_inc_absorb(state.ctx, SHAKE256_RATE, input, inlen)
}
func shake256_inc_finalize(state *shake256incctx) {
	keccak_inc_finalize(state.ctx, SHAKE256_RATE, 0x1F)
}
func shake256_inc_squeeze(output *uint8, outlen uint64, state *shake256incctx) {
	keccak_inc_squeeze(output, outlen, state.ctx, SHAKE256_RATE)
}
func shake256_inc_ctx_clone(dest *shake256incctx, src *shake256incctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*26))
}
func shake256_inc_ctx_release(state *shake256incctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func shake128_absorb(state *shake128ctx, input *uint8, inlen uint64) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 25)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_absorb(state.ctx, SHAKE128_RATE, input, inlen, 0x1F)
}
func shake128_squeezeblocks(output *uint8, nblocks uint64, state *shake128ctx) {
	keccak_squeezeblocks(output, nblocks, state.ctx, SHAKE128_RATE)
}
func shake128_ctx_clone(dest *shake128ctx, src *shake128ctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 25)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*25))
}
func shake128_ctx_release(state *shake128ctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func shake256_absorb(state *shake256ctx, input *uint8, inlen uint64) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 25)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_absorb(state.ctx, SHAKE256_RATE, input, inlen, 0x1F)
}
func shake256_squeezeblocks(output *uint8, nblocks uint64, state *shake256ctx) {
	keccak_squeezeblocks(output, nblocks, state.ctx, SHAKE256_RATE)
}
func shake256_ctx_clone(dest *shake256ctx, src *shake256ctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 25)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*25))
}
func shake256_ctx_release(state *shake256ctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func shake128(output *uint8, outlen uint64, input *uint8, inlen uint64) {
	var (
		nblocks uint64 = outlen / SHAKE128_RATE
		t       [168]uint8
		s       shake128ctx
	)
	shake128_absorb(&s, input, inlen)
	shake128_squeezeblocks(output, nblocks, &s)
	output = (*uint8)(unsafe.Add(unsafe.Pointer(output), nblocks*SHAKE128_RATE))
	outlen -= nblocks * SHAKE128_RATE
	if outlen != 0 {
		shake128_squeezeblocks(&t[0], 1, &s)
		for i := uint64(0); i < outlen; i++ {
			*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
		}
	}
	shake128_ctx_release(&s)
}
func shake256(output *uint8, outlen uint64, input *uint8, inlen uint64) {
	var (
		nblocks uint64 = outlen / SHAKE256_RATE
		t       [136]uint8
		s       shake256ctx
	)
	shake256_absorb(&s, input, inlen)
	shake256_squeezeblocks(output, nblocks, &s)
	output = (*uint8)(unsafe.Add(unsafe.Pointer(output), nblocks*SHAKE256_RATE))
	outlen -= nblocks * SHAKE256_RATE
	if outlen != 0 {
		shake256_squeezeblocks(&t[0], 1, &s)
		for i := uint64(0); i < outlen; i++ {
			*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
		}
	}
	shake256_ctx_release(&s)
}
func sha3_256_inc_init(state *sha3_256incctx) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_inc_init(state.ctx)
}
func sha3_256_inc_ctx_clone(dest *sha3_256incctx, src *sha3_256incctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*26))
}
func sha3_256_inc_ctx_release(state *sha3_256incctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func sha3_256_inc_absorb(state *sha3_256incctx, input *uint8, inlen uint64) {
	keccak_inc_absorb(state.ctx, SHA3_256_RATE, input, inlen)
}
func sha3_256_inc_finalize(output *uint8, state *sha3_256incctx) {
	var t [136]uint8
	keccak_inc_finalize(state.ctx, SHA3_256_RATE, 0x6)
	keccak_squeezeblocks(&t[0], 1, state.ctx, SHA3_256_RATE)
	sha3_256_inc_ctx_release(state)
	for i := uint64(0); i < 32; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
func sha3_256(output *uint8, input *uint8, inlen uint64) {
	var (
		s [25]uint64
		t [136]uint8
	)
	keccak_absorb(&s[0], SHA3_256_RATE, input, inlen, 0x6)
	keccak_squeezeblocks(&t[0], 1, &s[0], SHA3_256_RATE)
	for i := uint64(0); i < 32; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
func sha3_384_inc_init(state *sha3_384incctx) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_inc_init(state.ctx)
}
func sha3_384_inc_ctx_clone(dest *sha3_384incctx, src *sha3_384incctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*26))
}
func sha3_384_inc_absorb(state *sha3_384incctx, input *uint8, inlen uint64) {
	keccak_inc_absorb(state.ctx, SHA3_384_RATE, input, inlen)
}
func sha3_384_inc_ctx_release(state *sha3_384incctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func sha3_384_inc_finalize(output *uint8, state *sha3_384incctx) {
	var t [104]uint8
	keccak_inc_finalize(state.ctx, SHA3_384_RATE, 0x6)
	keccak_squeezeblocks(&t[0], 1, state.ctx, SHA3_384_RATE)
	sha3_384_inc_ctx_release(state)
	for i := uint64(0); i < 48; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
func sha3_384(output *uint8, input *uint8, inlen uint64) {
	var (
		s [25]uint64
		t [104]uint8
	)
	keccak_absorb(&s[0], SHA3_384_RATE, input, inlen, 0x6)
	keccak_squeezeblocks(&t[0], 1, &s[0], SHA3_384_RATE)
	for i := uint64(0); i < 48; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
func sha3_512_inc_init(state *sha3_512incctx) {
	state.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if state.ctx == nil {
		os.Exit(111)
	}
	keccak_inc_init(state.ctx)
}
func sha3_512_inc_ctx_clone(dest *sha3_512incctx, src *sha3_512incctx) {
	dest.ctx = (*uint64)(libc.Malloc(int(unsafe.Sizeof(uint64(0)) * 26)))
	if dest.ctx == nil {
		os.Exit(111)
	}
	libc.MemCpy(unsafe.Pointer(dest.ctx), unsafe.Pointer(src.ctx), int(unsafe.Sizeof(uint64(0))*26))
}
func sha3_512_inc_absorb(state *sha3_512incctx, input *uint8, inlen uint64) {
	keccak_inc_absorb(state.ctx, SHA3_512_RATE, input, inlen)
}
func sha3_512_inc_ctx_release(state *sha3_512incctx) {
	libc.Free(unsafe.Pointer(state.ctx))
}
func sha3_512_inc_finalize(output *uint8, state *sha3_512incctx) {
	var t [72]uint8
	keccak_inc_finalize(state.ctx, SHA3_512_RATE, 0x6)
	keccak_squeezeblocks(&t[0], 1, state.ctx, SHA3_512_RATE)
	sha3_512_inc_ctx_release(state)
	for i := uint64(0); i < 64; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
func sha3_512(output *uint8, input *uint8, inlen uint64) {
	var (
		s [25]uint64
		t [72]uint8
	)
	keccak_absorb(&s[0], SHA3_512_RATE, input, inlen, 0x6)
	keccak_squeezeblocks(&t[0], 1, &s[0], SHA3_512_RATE)
	for i := uint64(0); i < 64; i++ {
		*(*uint8)(unsafe.Add(unsafe.Pointer(output), i)) = t[i]
	}
}
