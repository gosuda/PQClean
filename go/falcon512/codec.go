package falcon512

import "unsafe"

func PQCLEAN_FALCON512_CLEAN_modq_encode(out unsafe.Pointer, max_out_len uint64, x *uint16, logn uint) uint64 {
	var (
		n       uint64
		out_len uint64
		u       uint64
		buf     *uint8
		acc     uint32
		acc_len int
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		if int(*(*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(u)))) >= 12289 {
			return 0
		}
	}
	out_len = ((n * 14) + 7) >> 3
	if out == nil {
		return out_len
	}
	if out_len > max_out_len {
		return 0
	}
	buf = (*uint8)(out)
	acc = 0
	acc_len = 0
	for u = 0; u < n; u++ {
		acc = uint32(int32((int(acc) << 14) | int(*(*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(u))))))
		acc_len += 14
		for acc_len >= 8 {
			acc_len -= 8
			*func() *uint8 {
				p_ := &buf
				x := *p_
				*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
				return x
			}() = uint8(int8(int(acc) >> acc_len))
		}
	}
	if acc_len > 0 {
		*buf = uint8(int8(int(acc) << (8 - acc_len)))
	}
	return out_len
}
func PQCLEAN_FALCON512_CLEAN_modq_decode(x *uint16, logn uint, in unsafe.Pointer, max_in_len uint64) uint64 {
	var (
		n       uint64
		in_len  uint64
		u       uint64
		buf     *uint8
		acc     uint32
		acc_len int
	)
	n = uint64(1 << logn)
	in_len = ((n * 14) + 7) >> 3
	if in_len > max_in_len {
		return 0
	}
	buf = (*uint8)(in)
	acc = 0
	acc_len = 0
	u = 0
	for u < n {
		acc = uint32(int32((int(acc) << 8) | int(*func() *uint8 {
			p_ := &buf
			x := *p_
			*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
			return x
		}())))
		acc_len += 8
		if acc_len >= 14 {
			var w uint
			acc_len -= 14
			w = uint((int(acc) >> acc_len) & 0x3FFF)
			if w >= 12289 {
				return 0
			}
			*(*uint16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(uint16(0))*uintptr(func() uint64 {
				p_ := &u
				x := *p_
				*p_++
				return x
			}()))) = uint16(w)
		}
	}
	if (int(acc) & ((1 << acc_len) - 1)) != 0 {
		return 0
	}
	return in_len
}
func PQCLEAN_FALCON512_CLEAN_trim_i16_encode(out unsafe.Pointer, max_out_len uint64, x *int16, logn uint, bits uint) uint64 {
	var (
		n       uint64
		u       uint64
		out_len uint64
		minv    int
		maxv    int
		buf     *uint8
		acc     uint32
		mask    uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	maxv = int((1 << (bits - 1)) - 1)
	minv = -maxv
	for u = 0; u < n; u++ {
		if int(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u)))) < minv || int(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u)))) > maxv {
			return 0
		}
	}
	out_len = ((n * uint64(bits)) + 7) >> 3
	if out == nil {
		return out_len
	}
	if out_len > max_out_len {
		return 0
	}
	buf = (*uint8)(out)
	acc = 0
	acc_len = 0
	mask = uint32((1 << bits) - 1)
	for u = 0; u < n; u++ {
		acc = uint32((uint(acc) << bits) | uint(int(uint16(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u)))))&int(mask)))
		acc_len += bits
		for acc_len >= 8 {
			acc_len -= 8
			*func() *uint8 {
				p_ := &buf
				x := *p_
				*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
				return x
			}() = uint8(uint(acc) >> acc_len)
		}
	}
	if acc_len > 0 {
		*func() *uint8 {
			p_ := &buf
			x := *p_
			*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
			return x
		}() = uint8(uint(acc) << (8 - acc_len))
	}
	return out_len
}
func PQCLEAN_FALCON512_CLEAN_trim_i16_decode(x *int16, logn uint, bits uint, in unsafe.Pointer, max_in_len uint64) uint64 {
	var (
		n       uint64
		in_len  uint64
		buf     *uint8
		u       uint64
		acc     uint32
		mask1   uint32
		mask2   uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	in_len = ((n * uint64(bits)) + 7) >> 3
	if in_len > max_in_len {
		return 0
	}
	buf = (*uint8)(in)
	u = 0
	acc = 0
	acc_len = 0
	mask1 = uint32((1 << bits) - 1)
	mask2 = uint32(1 << (bits - 1))
	for u < n {
		acc = uint32(int32((int(acc) << 8) | int(*func() *uint8 {
			p_ := &buf
			x := *p_
			*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
			return x
		}())))
		acc_len += 8
		for acc_len >= bits && u < n {
			var w uint32
			acc_len -= bits
			w = uint32((uint(acc) >> acc_len) & uint(mask1))
			w |= uint32(int32(-(int(w) & int(mask2))))
			if int(w) == int(-mask2) {
				return 0
			}
			w |= uint32(int32(-(int(w) & int(mask2))))
			*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(func() uint64 {
				p_ := &u
				x := *p_
				*p_++
				return x
			}()))) = int16(*(*int32)(unsafe.Pointer(&w)))
		}
	}
	if (uint(acc) & ((1 << acc_len) - 1)) != 0 {
		return 0
	}
	return in_len
}
func PQCLEAN_FALCON512_CLEAN_trim_i8_encode(out unsafe.Pointer, max_out_len uint64, x *int8, logn uint, bits uint) uint64 {
	var (
		n       uint64
		u       uint64
		out_len uint64
		minv    int
		maxv    int
		buf     *uint8
		acc     uint32
		mask    uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	maxv = int((1 << (bits - 1)) - 1)
	minv = -maxv
	for u = 0; u < n; u++ {
		if int(*(*int8)(unsafe.Add(unsafe.Pointer(x), u))) < minv || int(*(*int8)(unsafe.Add(unsafe.Pointer(x), u))) > maxv {
			return 0
		}
	}
	out_len = ((n * uint64(bits)) + 7) >> 3
	if out == nil {
		return out_len
	}
	if out_len > max_out_len {
		return 0
	}
	buf = (*uint8)(out)
	acc = 0
	acc_len = 0
	mask = uint32((1 << bits) - 1)
	for u = 0; u < n; u++ {
		acc = uint32((uint(acc) << bits) | uint(int(uint8(*(*int8)(unsafe.Add(unsafe.Pointer(x), u))))&int(mask)))
		acc_len += bits
		for acc_len >= 8 {
			acc_len -= 8
			*func() *uint8 {
				p_ := &buf
				x := *p_
				*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
				return x
			}() = uint8(uint(acc) >> acc_len)
		}
	}
	if acc_len > 0 {
		*func() *uint8 {
			p_ := &buf
			x := *p_
			*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
			return x
		}() = uint8(uint(acc) << (8 - acc_len))
	}
	return out_len
}
func PQCLEAN_FALCON512_CLEAN_trim_i8_decode(x *int8, logn uint, bits uint, in unsafe.Pointer, max_in_len uint64) uint64 {
	var (
		n       uint64
		in_len  uint64
		buf     *uint8
		u       uint64
		acc     uint32
		mask1   uint32
		mask2   uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	in_len = ((n * uint64(bits)) + 7) >> 3
	if in_len > max_in_len {
		return 0
	}
	buf = (*uint8)(in)
	u = 0
	acc = 0
	acc_len = 0
	mask1 = uint32((1 << bits) - 1)
	mask2 = uint32(1 << (bits - 1))
	for u < n {
		acc = uint32(int32((int(acc) << 8) | int(*func() *uint8 {
			p_ := &buf
			x := *p_
			*p_ = (*uint8)(unsafe.Add(unsafe.Pointer(*p_), 1))
			return x
		}())))
		acc_len += 8
		for acc_len >= bits && u < n {
			var w uint32
			acc_len -= bits
			w = uint32((uint(acc) >> acc_len) & uint(mask1))
			w |= uint32(int32(-(int(w) & int(mask2))))
			if int(w) == int(-mask2) {
				return 0
			}
			*(*int8)(unsafe.Add(unsafe.Pointer(x), func() uint64 {
				p_ := &u
				x := *p_
				*p_++
				return x
			}())) = int8(*(*int32)(unsafe.Pointer(&w)))
		}
	}
	if (uint(acc) & ((1 << acc_len) - 1)) != 0 {
		return 0
	}
	return in_len
}
func PQCLEAN_FALCON512_CLEAN_comp_encode(out unsafe.Pointer, max_out_len uint64, x *int16, logn uint) uint64 {
	var (
		buf     *uint8
		n       uint64
		u       uint64
		v       uint64
		acc     uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	buf = (*uint8)(out)
	for u = 0; u < n; u++ {
		if int(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u)))) < -2047 || int(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u)))) > +2047 {
			return 0
		}
	}
	acc = 0
	acc_len = 0
	v = 0
	for u = 0; u < n; u++ {
		var (
			t int
			w uint
		)
		acc <<= 1
		t = int(*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u))))
		if t < 0 {
			t = -t
			acc |= 1
		}
		w = uint(t)
		acc <<= 7
		acc |= uint32(w & 127)
		w >>= 7
		acc_len += 8
		acc <<= uint32(w + 1)
		acc |= 1
		acc_len += w + 1
		for acc_len >= 8 {
			acc_len -= 8
			if buf != nil {
				if v >= max_out_len {
					return 0
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(buf), v)) = uint8(uint(acc) >> acc_len)
			}
			v++
		}
	}
	if acc_len > 0 {
		if buf != nil {
			if v >= max_out_len {
				return 0
			}
			*(*uint8)(unsafe.Add(unsafe.Pointer(buf), v)) = uint8(uint(acc) << (8 - acc_len))
		}
		v++
	}
	return v
}
func PQCLEAN_FALCON512_CLEAN_comp_decode(x *int16, logn uint, in unsafe.Pointer, max_in_len uint64) uint64 {
	var (
		buf     *uint8
		n       uint64
		u       uint64
		v       uint64
		acc     uint32
		acc_len uint
	)
	n = uint64(1 << logn)
	buf = (*uint8)(in)
	acc = 0
	acc_len = 0
	v = 0
	for u = 0; u < n; u++ {
		var (
			b uint
			s uint
			m uint
		)
		if v >= max_in_len {
			return 0
		}
		acc = uint32(int32((int(acc) << 8) | int(uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(buf), func() uint64 {
			p_ := &v
			x := *p_
			*p_++
			return x
		}()))))))
		b = uint(acc) >> acc_len
		s = b & 128
		m = b & 127
		for {
			if acc_len == 0 {
				if v >= max_in_len {
					return 0
				}
				acc = uint32(int32((int(acc) << 8) | int(uint32(*(*uint8)(unsafe.Add(unsafe.Pointer(buf), func() uint64 {
					p_ := &v
					x := *p_
					*p_++
					return x
				}()))))))
				acc_len = 8
			}
			acc_len--
			if ((uint(acc) >> acc_len) & 1) != 0 {
				break
			}
			m += 128
			if m > 2047 {
				return 0
			}
		}
		if s != 0 && m == 0 {
			return 0
		}
		if s != 0 {
			*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(uint16(-m))
		} else {
			*(*int16)(unsafe.Add(unsafe.Pointer(x), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(uint16(m))
		}
	}
	if (uint(acc) & ((1 << acc_len) - 1)) != 0 {
		return 0
	}
	return v
}

var PQCLEAN_FALCON512_CLEAN_max_fg_bits [11]uint8 = [11]uint8{0, 8, 8, 8, 8, 8, 7, 7, 6, 6, 5}
var PQCLEAN_FALCON512_CLEAN_max_FG_bits [11]uint8 = [11]uint8{0, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}
var PQCLEAN_FALCON512_CLEAN_max_sig_bits [11]uint8 = [11]uint8{0, 10, 11, 11, 12, 12, 12, 12, 12, 12, 12}
