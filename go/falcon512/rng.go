package falcon512

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func PQCLEAN_FALCON512_CLEAN_prng_init(p *prng, src *shake256incctx) {
	var (
		tmp [56]uint8
		th  uint64
		tl  uint64
		i   int
		d32 *uint32 = (*uint32)(unsafe.Pointer(&p.state.d[0]))
		d64 *uint64 = (*uint64)(unsafe.Pointer(&p.state.d[0]))
	)
	shake256_inc_squeeze(&tmp[0], 56, src)
	for i = 0; i < 14; i++ {
		var w uint32
		w = uint32(int32(int(uint32(tmp[(i<<2)+0])) | int(uint32(tmp[(i<<2)+1]))<<8 | int(uint32(tmp[(i<<2)+2]))<<16 | int(uint32(tmp[(i<<2)+3]))<<24))
		*(*uint32)(unsafe.Add(unsafe.Pointer(d32), unsafe.Sizeof(uint32(0))*uintptr(i))) = w
	}
	tl = uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(d32), unsafe.Sizeof(uint32(0))*(48/unsafe.Sizeof(uint32(0))))))
	th = uint64(*(*uint32)(unsafe.Add(unsafe.Pointer(d32), unsafe.Sizeof(uint32(0))*(52/unsafe.Sizeof(uint32(0))))))
	*(*uint64)(unsafe.Add(unsafe.Pointer(d64), unsafe.Sizeof(uint64(0))*(48/unsafe.Sizeof(uint64(0))))) = tl + (th << 32)
	PQCLEAN_FALCON512_CLEAN_prng_refill(p)
}
func PQCLEAN_FALCON512_CLEAN_prng_refill(p *prng) {
	var (
		CW [4]uint32 = [4]uint32{0x61707865, 0x3320646E, 0x79622D32, 0x6B206574}
		cc uint64
		u  uint64
	)
	cc = *(*uint64)(unsafe.Pointer(&p.state.d[48]))
	for u = 0; u < 8; u++ {
		var (
			state [16]uint32
			v     uint64
			i     int
		)
		libc.MemCpy(unsafe.Pointer(&state[0]), unsafe.Pointer(&CW[0]), int(unsafe.Sizeof([4]uint32{})))
		libc.MemCpy(unsafe.Pointer(&state[4]), unsafe.Pointer(&p.state.d[0]), 48)
		state[14] ^= uint32(cc)
		state[15] ^= uint32(cc >> 32)
		for i = 0; i < 10; i++ {
			for {
				state[0] += state[4]
				state[12] ^= state[0]
				state[12] = uint32(int32((int(state[12]) << 16) | int(state[12])>>16))
				state[8] += state[12]
				state[4] ^= state[8]
				state[4] = uint32(int32((int(state[4]) << 12) | int(state[4])>>20))
				state[0] += state[4]
				state[12] ^= state[0]
				state[12] = uint32(int32((int(state[12]) << 8) | int(state[12])>>24))
				state[8] += state[12]
				state[4] ^= state[8]
				state[4] = uint32(int32((int(state[4]) << 7) | int(state[4])>>25))
				if true {
					break
				}
			}
			for {
				state[1] += state[5]
				state[13] ^= state[1]
				state[13] = uint32(int32((int(state[13]) << 16) | int(state[13])>>16))
				state[9] += state[13]
				state[5] ^= state[9]
				state[5] = uint32(int32((int(state[5]) << 12) | int(state[5])>>20))
				state[1] += state[5]
				state[13] ^= state[1]
				state[13] = uint32(int32((int(state[13]) << 8) | int(state[13])>>24))
				state[9] += state[13]
				state[5] ^= state[9]
				state[5] = uint32(int32((int(state[5]) << 7) | int(state[5])>>25))
				if true {
					break
				}
			}
			for {
				state[2] += state[6]
				state[14] ^= state[2]
				state[14] = uint32(int32((int(state[14]) << 16) | int(state[14])>>16))
				state[10] += state[14]
				state[6] ^= state[10]
				state[6] = uint32(int32((int(state[6]) << 12) | int(state[6])>>20))
				state[2] += state[6]
				state[14] ^= state[2]
				state[14] = uint32(int32((int(state[14]) << 8) | int(state[14])>>24))
				state[10] += state[14]
				state[6] ^= state[10]
				state[6] = uint32(int32((int(state[6]) << 7) | int(state[6])>>25))
				if true {
					break
				}
			}
			for {
				state[3] += state[7]
				state[15] ^= state[3]
				state[15] = uint32(int32((int(state[15]) << 16) | int(state[15])>>16))
				state[11] += state[15]
				state[7] ^= state[11]
				state[7] = uint32(int32((int(state[7]) << 12) | int(state[7])>>20))
				state[3] += state[7]
				state[15] ^= state[3]
				state[15] = uint32(int32((int(state[15]) << 8) | int(state[15])>>24))
				state[11] += state[15]
				state[7] ^= state[11]
				state[7] = uint32(int32((int(state[7]) << 7) | int(state[7])>>25))
				if true {
					break
				}
			}
			for {
				state[0] += state[5]
				state[15] ^= state[0]
				state[15] = uint32(int32((int(state[15]) << 16) | int(state[15])>>16))
				state[10] += state[15]
				state[5] ^= state[10]
				state[5] = uint32(int32((int(state[5]) << 12) | int(state[5])>>20))
				state[0] += state[5]
				state[15] ^= state[0]
				state[15] = uint32(int32((int(state[15]) << 8) | int(state[15])>>24))
				state[10] += state[15]
				state[5] ^= state[10]
				state[5] = uint32(int32((int(state[5]) << 7) | int(state[5])>>25))
				if true {
					break
				}
			}
			for {
				state[1] += state[6]
				state[12] ^= state[1]
				state[12] = uint32(int32((int(state[12]) << 16) | int(state[12])>>16))
				state[11] += state[12]
				state[6] ^= state[11]
				state[6] = uint32(int32((int(state[6]) << 12) | int(state[6])>>20))
				state[1] += state[6]
				state[12] ^= state[1]
				state[12] = uint32(int32((int(state[12]) << 8) | int(state[12])>>24))
				state[11] += state[12]
				state[6] ^= state[11]
				state[6] = uint32(int32((int(state[6]) << 7) | int(state[6])>>25))
				if true {
					break
				}
			}
			for {
				state[2] += state[7]
				state[13] ^= state[2]
				state[13] = uint32(int32((int(state[13]) << 16) | int(state[13])>>16))
				state[8] += state[13]
				state[7] ^= state[8]
				state[7] = uint32(int32((int(state[7]) << 12) | int(state[7])>>20))
				state[2] += state[7]
				state[13] ^= state[2]
				state[13] = uint32(int32((int(state[13]) << 8) | int(state[13])>>24))
				state[8] += state[13]
				state[7] ^= state[8]
				state[7] = uint32(int32((int(state[7]) << 7) | int(state[7])>>25))
				if true {
					break
				}
			}
			for {
				state[3] += state[4]
				state[14] ^= state[3]
				state[14] = uint32(int32((int(state[14]) << 16) | int(state[14])>>16))
				state[9] += state[14]
				state[4] ^= state[9]
				state[4] = uint32(int32((int(state[4]) << 12) | int(state[4])>>20))
				state[3] += state[4]
				state[14] ^= state[3]
				state[14] = uint32(int32((int(state[14]) << 8) | int(state[14])>>24))
				state[9] += state[14]
				state[4] ^= state[9]
				state[4] = uint32(int32((int(state[4]) << 7) | int(state[4])>>25))
				if true {
					break
				}
			}
		}
		for v = 0; v < 4; v++ {
			state[v] += CW[v]
		}
		for v = 4; v < 14; v++ {
			state[v] += *(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&p.state.d[0]))), unsafe.Sizeof(uint32(0))*uintptr(v-4)))
		}
		state[14] += uint32(int32(int(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&p.state.d[0]))), unsafe.Sizeof(uint32(0))*10))) ^ int(uint32(cc))))
		state[15] += uint32(int32(int(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&p.state.d[0]))), unsafe.Sizeof(uint32(0))*11))) ^ int(uint32(cc>>32))))
		cc++
		for v = 0; v < 16; v++ {
			p.buf.d[(u<<2)+(v<<5)+0] = uint8(state[v])
			p.buf.d[(u<<2)+(v<<5)+1] = uint8(int8(int(state[v]) >> 8))
			p.buf.d[(u<<2)+(v<<5)+2] = uint8(int8(int(state[v]) >> 16))
			p.buf.d[(u<<2)+(v<<5)+3] = uint8(int8(int(state[v]) >> 24))
		}
	}
	*(*uint64)(unsafe.Pointer(&p.state.d[48])) = cc
	p.ptr = 0
}
func PQCLEAN_FALCON512_CLEAN_prng_get_bytes(p *prng, dst unsafe.Pointer, len_ uint64) {
	var buf *uint8
	buf = (*uint8)(dst)
	for len_ > 0 {
		var clen uint64
		clen = uint64((512) - uintptr(p.ptr))
		if clen > len_ {
			clen = len_
		}
		libc.MemCpy(unsafe.Pointer(buf), unsafe.Pointer(&p.buf.d[0]), int(clen))
		buf = (*uint8)(unsafe.Add(unsafe.Pointer(buf), clen))
		len_ -= clen
		p.ptr += clen
		if p.ptr == uint64(512) {
			PQCLEAN_FALCON512_CLEAN_prng_refill(p)
		}
	}
}
