package falcon512

const FALCON_KEYGEN_TEMP_1 = 136
const FALCON_KEYGEN_TEMP_2 = 272
const FALCON_KEYGEN_TEMP_3 = 224
const FALCON_KEYGEN_TEMP_4 = 448
const FALCON_KEYGEN_TEMP_5 = 896
const FALCON_KEYGEN_TEMP_6 = 1792
const FALCON_KEYGEN_TEMP_7 = 3584
const FALCON_KEYGEN_TEMP_8 = 7168
const FALCON_KEYGEN_TEMP_9 = 14336
const FALCON_KEYGEN_TEMP_10 = 28672

func set_fpu_cw(x uint) uint {
	return x
}

type prng struct {
	buf struct {
		// union
		d         [512]uint8
		dummy_u64 uint64
	}
	ptr   uint64
	state struct {
		// union
		d         [256]uint8
		dummy_u64 uint64
	}
	type_ int
}

func prng_get_u64(p *prng) uint64 {
	var u uint64
	u = p.ptr
	if u >= uint64((512)-9) {
		PQCLEAN_FALCON512_CLEAN_prng_refill(p)
		u = 0
	}
	p.ptr = u + 8
	return uint64(p.buf.d[u+0]) | uint64(p.buf.d[u+1])<<8 | uint64(p.buf.d[u+2])<<16 | uint64(p.buf.d[u+3])<<24 | uint64(p.buf.d[u+4])<<32 | uint64(p.buf.d[u+5])<<40 | uint64(p.buf.d[u+6])<<48 | uint64(p.buf.d[u+7])<<56
}
func prng_get_u8(p *prng) uint {
	var v uint
	v = uint(p.buf.d[func() uint64 {
		p_ := &p.ptr
		x := *p_
		*p_++
		return x
	}()])
	if p.ptr == uint64(512) {
		PQCLEAN_FALCON512_CLEAN_prng_refill(p)
	}
	return v
}

type sampler_context struct {
	p         prng
	sigma_min fpr
}
