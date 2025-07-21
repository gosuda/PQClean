package falcon512

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func PQCLEAN_FALCON512_CLEAN_hash_to_point_vartime(sc *shake256incctx, x *uint16, logn uint) {
	var n uint64
	n = uint64(1 << logn)
	for n > 0 {
		var (
			buf [2]uint8
			w   uint32
		)
		shake256_inc_squeeze((*uint8)(unsafe.Pointer(&buf[0])), uint64(2), sc)
		w = uint32((uint(buf[0]) << 8) | uint(buf[1]))
		if int(w) < 61445 {
			for int(w) >= 12289 {
				w -= 12289
			}
			*func() *uint16 {
				p_ := &x
				x := *p_
				*p_ = (*uint16)(unsafe.Add(unsafe.Pointer(*p_), unsafe.Sizeof(uint16(0))*1))
				return x
			}() = uint16(w)
			n--
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_hash_to_point_ct(sc *shake256incctx, x *uint16, logn uint, tmp *uint8) {
	var (
		overtab [11]uint16 = [11]uint16{0, 65, 67, 71, 77, 86, 100, 122, 154, 205, 287}
		n       uint
		n2      uint
		u       uint
		m       uint
		p       uint
		over    uint
		tt1     *uint16
		tt2     [63]uint16
	)
	n = 1 << logn
	n2 = n << 1
	over = uint(overtab[logn])
	m = n + over
	tt1 = (*uint16)(unsafe.Pointer(tmp))
	for u = 0; u < m; u++ {
		var (
			buf [2]uint8
			w   uint32
			wr  uint32
		)
		shake256_inc_squeeze(&buf[0], uint64(2), sc)
		w = uint32(int32((int(uint32(buf[0])) << 8) | int(uint32(buf[1]))))
		wr = uint32(int32(int(w) - ((((int(w) - 24578) >> 31) - 1) & 24578)))
		wr = uint32(int32(int(wr) - ((((int(wr) - 24578) >> 31) - 1) & 24578)))
		wr = uint32(int32(int(wr) - ((((int(wr) - 12289) >> 31) - 1) & 12289)))
		wr |= uint32(int32(((int(w) - 61445) >> 31) - 1))
		if u < n {
			*(*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(u))) = uint16(wr)
		} else if u < n2 {
			*(*uint16)(unsafe.Add(unsafe.Pointer(tt1), unsafe.Sizeof(uint16(0))*uintptr(u-n))) = uint16(wr)
		} else {
			tt2[u-n2] = uint16(wr)
		}
	}
	for p = 1; p <= over; p <<= 1 {
		var v uint
		v = 0
		for u = 0; u < m; u++ {
			var (
				s  *uint16
				d  *uint16
				j  uint
				sv uint
				dv uint
				mk uint
			)
			if u < n {
				s = (*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(u)))
			} else if u < n2 {
				s = (*uint16)(unsafe.Add(unsafe.Pointer(tt1), unsafe.Sizeof(uint16(0))*uintptr(u-n)))
			} else {
				s = &tt2[u-n2]
			}
			sv = uint(*s)
			j = u - v
			mk = (sv >> 15) - 1
			v -= mk
			if u < p {
				continue
			}
			if (u - p) < n {
				d = (*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(u-p)))
			} else if (u - p) < n2 {
				d = (*uint16)(unsafe.Add(unsafe.Pointer(tt1), unsafe.Sizeof(uint16(0))*uintptr((u-p)-n)))
			} else {
				d = &tt2[(u-p)-n2]
			}
			dv = uint(*d)
			mk &= -(((j & p) + 0x1FF) >> 9)
			*s = uint16(sv ^ (mk & (sv ^ dv)))
			*d = uint16(dv ^ (mk & (sv ^ dv)))
		}
	}
}

var l2bound [11]uint32 = [11]uint32{0, 101498, 208714, 428865, 892039, 1852696, 3842630, 7959734, 16468416, 34034726, 70265242}

func PQCLEAN_FALCON512_CLEAN_is_short(s1 *int16, s2 *int16, logn uint) int {
	var (
		n  uint64
		u  uint64
		s  uint32
		ng uint32
	)
	n = uint64(1 << logn)
	s = 0
	ng = 0
	for u = 0; u < n; u++ {
		var z int32
		z = int32(*(*int16)(unsafe.Add(unsafe.Pointer(s1), unsafe.Sizeof(int16(0))*uintptr(u))))
		s += uint32(int32(int(z) * int(z)))
		ng |= s
		z = int32(*(*int16)(unsafe.Add(unsafe.Pointer(s2), unsafe.Sizeof(int16(0))*uintptr(u))))
		s += uint32(int32(int(z) * int(z)))
		ng |= s
	}
	s |= uint32(int32(-(int(ng) >> 31)))
	return int(libc.BoolToInt(int(s) <= int(l2bound[logn])))
}
func PQCLEAN_FALCON512_CLEAN_is_short_half(sqn uint32, s2 *int16, logn uint) int {
	var (
		n  uint64
		u  uint64
		ng uint32
	)
	n = uint64(1 << logn)
	ng = uint32(int32(-(int(sqn) >> 31)))
	for u = 0; u < n; u++ {
		var z int32
		z = int32(*(*int16)(unsafe.Add(unsafe.Pointer(s2), unsafe.Sizeof(int16(0))*uintptr(u))))
		sqn += uint32(int32(int(z) * int(z)))
		ng |= sqn
	}
	sqn |= uint32(int32(-(int(ng) >> 31)))
	return int(libc.BoolToInt(int(sqn) <= int(l2bound[logn])))
}
