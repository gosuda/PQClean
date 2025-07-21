package falcon512

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func ffLDL_treesize(logn uint) uint {
	return (logn + 1) << logn
}
func ffLDL_fft_inner(tree *fpr, g0 *fpr, g1 *fpr, logn uint, tmp *fpr) {
	var (
		n  uint64
		hn uint64
	)
	n = uint64(1 << logn)
	if n == 1 {
		*(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0)) = *(*fpr)(unsafe.Add(unsafe.Pointer(g0), unsafe.Sizeof(fpr(0))*0))
		return
	}
	hn = n >> 1
	PQCLEAN_FALCON512_CLEAN_poly_LDLmv_fft(tmp, tree, g0, g1, g0, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(g1, (*fpr)(unsafe.Add(unsafe.Pointer(g1), unsafe.Sizeof(fpr(0))*uintptr(hn))), g0, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(g0, (*fpr)(unsafe.Add(unsafe.Pointer(g0), unsafe.Sizeof(fpr(0))*uintptr(hn))), tmp, logn)
	ffLDL_fft_inner((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n))), g1, (*fpr)(unsafe.Add(unsafe.Pointer(g1), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, tmp)
	ffLDL_fft_inner((*fpr)(unsafe.Add(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n)))), unsafe.Sizeof(fpr(0))*uintptr(ffLDL_treesize(logn-1)))), g0, (*fpr)(unsafe.Add(unsafe.Pointer(g0), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, tmp)
}
func ffLDL_fft(tree *fpr, g00 *fpr, g01 *fpr, g11 *fpr, logn uint, tmp *fpr) {
	var (
		n   uint64
		hn  uint64
		d00 *fpr
		d11 *fpr
	)
	n = uint64(1 << logn)
	if n == 1 {
		*(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0)) = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*0))
		return
	}
	hn = n >> 1
	d00 = tmp
	d11 = (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n)))
	tmp = (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n<<1)))
	libc.MemCpy(unsafe.Pointer(d00), unsafe.Pointer(g00), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_LDLmv_fft(d11, tree, g00, g01, g11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), d00, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(d00, (*fpr)(unsafe.Add(unsafe.Pointer(d00), unsafe.Sizeof(fpr(0))*uintptr(hn))), d11, logn)
	libc.MemCpy(unsafe.Pointer(d11), unsafe.Pointer(tmp), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	ffLDL_fft_inner((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n))), d11, (*fpr)(unsafe.Add(unsafe.Pointer(d11), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, tmp)
	ffLDL_fft_inner((*fpr)(unsafe.Add(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n)))), unsafe.Sizeof(fpr(0))*uintptr(ffLDL_treesize(logn-1)))), d00, (*fpr)(unsafe.Add(unsafe.Pointer(d00), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, tmp)
}
func ffLDL_binary_normalize(tree *fpr, orig_logn uint, logn uint) {
	var n uint64
	n = uint64(1 << logn)
	if n == 1 {
		*(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0)) = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_sqrt(*(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0))), fpr_inv_sigma[orig_logn])
	} else {
		ffLDL_binary_normalize((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n))), orig_logn, logn-1)
		ffLDL_binary_normalize((*fpr)(unsafe.Add(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n)))), unsafe.Sizeof(fpr(0))*uintptr(ffLDL_treesize(logn-1)))), orig_logn, logn-1)
	}
}
func smallints_to_fpr(r *fpr, t *int8, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(r), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_of(int64(*(*int8)(unsafe.Add(unsafe.Pointer(t), u))))
	}
}
func skoff_b00(logn uint) uint64 {
	_ = logn
	return 0
}
func skoff_b01(logn uint) uint64 {
	return uint64(1 << logn)
}
func skoff_b10(logn uint) uint64 {
	return uint64((1 << logn) * 2)
}
func skoff_b11(logn uint) uint64 {
	return uint64((1 << logn) * 3)
}
func skoff_tree(logn uint) uint64 {
	return uint64((1 << logn) * 4)
}
func PQCLEAN_FALCON512_CLEAN_expand_privkey(expanded_key *fpr, f *int8, g *int8, F *int8, G *int8, logn uint, tmp *uint8) {
	var (
		n    uint64
		rf   *fpr
		rg   *fpr
		rF   *fpr
		rG   *fpr
		b00  *fpr
		b01  *fpr
		b10  *fpr
		b11  *fpr
		g00  *fpr
		g01  *fpr
		g11  *fpr
		gxx  *fpr
		tree *fpr
	)
	n = uint64(1 << logn)
	b00 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b00(logn))))
	b01 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b01(logn))))
	b10 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b10(logn))))
	b11 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b11(logn))))
	tree = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_tree(logn))))
	rf = b01
	rg = b00
	rF = b11
	rG = b10
	smallints_to_fpr(rf, f, logn)
	smallints_to_fpr(rg, g, logn)
	smallints_to_fpr(rF, F, logn)
	smallints_to_fpr(rG, G, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(rf, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(rg, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(rF, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(rG, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(rf, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(rF, logn)
	g00 = (*fpr)(unsafe.Pointer(tmp))
	g01 = (*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(n)))
	g11 = (*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(n)))
	gxx = (*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(n)))
	libc.MemCpy(unsafe.Pointer(g00), unsafe.Pointer(b00), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(g00, logn)
	libc.MemCpy(unsafe.Pointer(gxx), unsafe.Pointer(b01), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(gxx, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(g00, gxx, logn)
	libc.MemCpy(unsafe.Pointer(g01), unsafe.Pointer(b00), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_muladj_fft(g01, b10, logn)
	libc.MemCpy(unsafe.Pointer(gxx), unsafe.Pointer(b01), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_muladj_fft(gxx, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(g01, gxx, logn)
	libc.MemCpy(unsafe.Pointer(g11), unsafe.Pointer(b10), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(g11, logn)
	libc.MemCpy(unsafe.Pointer(gxx), unsafe.Pointer(b11), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(gxx, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(g11, gxx, logn)
	ffLDL_fft(tree, g00, g01, g11, logn, gxx)
	ffLDL_binary_normalize(tree, logn, logn)
}

type samplerZ func(ctx unsafe.Pointer, mu fpr, sigma fpr) int

func ffSampling_fft_dyntree(samp samplerZ, samp_ctx unsafe.Pointer, t0 *fpr, t1 *fpr, g00 *fpr, g01 *fpr, g11 *fpr, orig_logn uint, logn uint, tmp *fpr) {
	var (
		n  uint64
		hn uint64
		z0 *fpr
		z1 *fpr
	)
	if logn == 0 {
		var leaf fpr
		leaf = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*0))
		leaf = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_sqrt(leaf), fpr_inv_sigma[orig_logn])
		*(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*0)) = fpr_of(int64(samp(samp_ctx, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*0)), leaf)))
		*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*0)) = fpr_of(int64(samp(samp_ctx, *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*0)), leaf)))
		return
	}
	n = uint64(1 << logn)
	hn = n >> 1
	PQCLEAN_FALCON512_CLEAN_poly_LDL_fft(g00, g01, g11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), g00, logn)
	libc.MemCpy(unsafe.Pointer(g00), unsafe.Pointer(tmp), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), g11, logn)
	libc.MemCpy(unsafe.Pointer(g11), unsafe.Pointer(tmp), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	libc.MemCpy(unsafe.Pointer(tmp), unsafe.Pointer(g01), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	libc.MemCpy(unsafe.Pointer(g01), unsafe.Pointer(g00), int(hn*uint64(unsafe.Sizeof(fpr(0)))))
	libc.MemCpy(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(hn)))), unsafe.Pointer(g11), int(hn*uint64(unsafe.Sizeof(fpr(0)))))
	z1 = (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n)))
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(z1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(hn))), t1, logn)
	ffSampling_fft_dyntree(samp, samp_ctx, z1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(hn))), g11, (*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(hn))), (*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(hn))), orig_logn, logn-1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(n))))
	PQCLEAN_FALCON512_CLEAN_poly_merge_fft((*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n<<1))), z1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn)
	libc.MemCpy(unsafe.Pointer(z1), unsafe.Pointer(t1), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_sub(z1, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n<<1))), logn)
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n<<1)))), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(tmp, z1, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(t0, tmp, logn)
	z0 = tmp
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(z0, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(hn))), t0, logn)
	ffSampling_fft_dyntree(samp, samp_ctx, z0, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(hn))), g00, (*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(hn))), g01, orig_logn, logn-1, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(n))))
	PQCLEAN_FALCON512_CLEAN_poly_merge_fft(t0, z0, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn)
}
func ffSampling_fft(samp samplerZ, samp_ctx unsafe.Pointer, z0 *fpr, z1 *fpr, tree *fpr, t0 *fpr, t1 *fpr, logn uint, tmp *fpr) {
	var (
		n     uint64
		hn    uint64
		tree0 *fpr
		tree1 *fpr
	)
	if logn == 2 {
		var (
			x0    fpr
			x1    fpr
			y0    fpr
			y1    fpr
			w0    fpr
			w1    fpr
			w2    fpr
			w3    fpr
			sigma fpr
			a_re  fpr
			a_im  fpr
			b_re  fpr
			b_im  fpr
			c_re  fpr
			c_im  fpr
		)
		tree0 = (*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*4))
		tree1 = (*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*8))
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*0))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*2))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*1))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*3))
		c_re = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, b_re)
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, b_im)
		w0 = fpr_half(c_re)
		w1 = fpr_half(c_im)
		c_re = fpr_sub(a_re, b_re)
		c_im = fpr_sub(a_im, b_im)
		w2 = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_add(c_re, c_im), fpr_invsqrt8)
		w3 = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_sub(c_im, c_re), fpr_invsqrt8)
		x0 = w2
		x1 = w3
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree1), unsafe.Sizeof(fpr(0))*3))
		w2 = fpr_of(int64(samp(samp_ctx, x0, sigma)))
		w3 = fpr_of(int64(samp(samp_ctx, x1, sigma)))
		a_re = fpr_sub(x0, w2)
		a_im = fpr_sub(x1, w3)
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(tree1), unsafe.Sizeof(fpr(0))*0))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(tree1), unsafe.Sizeof(fpr(0))*1))
		c_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_im))
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_re))
		x0 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_re, w0)
		x1 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_im, w1)
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree1), unsafe.Sizeof(fpr(0))*2))
		w0 = fpr_of(int64(samp(samp_ctx, x0, sigma)))
		w1 = fpr_of(int64(samp(samp_ctx, x1, sigma)))
		a_re = w0
		a_im = w1
		b_re = w2
		b_im = w3
		c_re = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_sub(b_re, b_im), fpr_invsqrt2)
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_add(b_re, b_im), fpr_invsqrt2)
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*0)) = func() fpr {
			w0 = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, c_re)
			return w0
		}()
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*2)) = func() fpr {
			w2 = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, c_im)
			return w2
		}()
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*1)) = func() fpr {
			w1 = fpr_sub(a_re, c_re)
			return w1
		}()
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*3)) = func() fpr {
			w3 = fpr_sub(a_im, c_im)
			return w3
		}()
		w0 = fpr_sub(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*0)), w0)
		w1 = fpr_sub(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*1)), w1)
		w2 = fpr_sub(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*2)), w2)
		w3 = fpr_sub(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*3)), w3)
		a_re = w0
		a_im = w2
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*2))
		w0 = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_im))
		w2 = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_re))
		a_re = w1
		a_im = w3
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*1))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*3))
		w1 = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_im))
		w3 = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_re))
		w0 = PQCLEAN_FALCON512_CLEAN_fpr_add(w0, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*0)))
		w1 = PQCLEAN_FALCON512_CLEAN_fpr_add(w1, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*1)))
		w2 = PQCLEAN_FALCON512_CLEAN_fpr_add(w2, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*2)))
		w3 = PQCLEAN_FALCON512_CLEAN_fpr_add(w3, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*3)))
		a_re = w0
		a_im = w2
		b_re = w1
		b_im = w3
		c_re = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, b_re)
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, b_im)
		w0 = fpr_half(c_re)
		w1 = fpr_half(c_im)
		c_re = fpr_sub(a_re, b_re)
		c_im = fpr_sub(a_im, b_im)
		w2 = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_add(c_re, c_im), fpr_invsqrt8)
		w3 = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_sub(c_im, c_re), fpr_invsqrt8)
		x0 = w2
		x1 = w3
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree0), unsafe.Sizeof(fpr(0))*3))
		w2 = func() fpr {
			y0 = fpr_of(int64(samp(samp_ctx, x0, sigma)))
			return y0
		}()
		w3 = func() fpr {
			y1 = fpr_of(int64(samp(samp_ctx, x1, sigma)))
			return y1
		}()
		a_re = fpr_sub(x0, y0)
		a_im = fpr_sub(x1, y1)
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(tree0), unsafe.Sizeof(fpr(0))*0))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(tree0), unsafe.Sizeof(fpr(0))*1))
		c_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_im))
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_re))
		x0 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_re, w0)
		x1 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_im, w1)
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree0), unsafe.Sizeof(fpr(0))*2))
		w0 = fpr_of(int64(samp(samp_ctx, x0, sigma)))
		w1 = fpr_of(int64(samp(samp_ctx, x1, sigma)))
		a_re = w0
		a_im = w1
		b_re = w2
		b_im = w3
		c_re = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_sub(b_re, b_im), fpr_invsqrt2)
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_mul(PQCLEAN_FALCON512_CLEAN_fpr_add(b_re, b_im), fpr_invsqrt2)
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*0)) = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, c_re)
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*2)) = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, c_im)
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*1)) = fpr_sub(a_re, c_re)
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*3)) = fpr_sub(a_im, c_im)
		return
	}
	if logn == 1 {
		var (
			x0    fpr
			x1    fpr
			y0    fpr
			y1    fpr
			sigma fpr
			a_re  fpr
			a_im  fpr
			b_re  fpr
			b_im  fpr
			c_re  fpr
			c_im  fpr
		)
		x0 = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*0))
		x1 = *(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*1))
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*3))
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*0)) = func() fpr {
			y0 = fpr_of(int64(samp(samp_ctx, x0, sigma)))
			return y0
		}()
		*(*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*1)) = func() fpr {
			y1 = fpr_of(int64(samp(samp_ctx, x1, sigma)))
			return y1
		}()
		a_re = fpr_sub(x0, y0)
		a_im = fpr_sub(x1, y1)
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*0))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*1))
		c_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_im))
		c_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(a_re, b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(a_im, b_re))
		x0 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_re, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*0)))
		x1 = PQCLEAN_FALCON512_CLEAN_fpr_add(c_im, *(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*1)))
		sigma = *(*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*2))
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*0)) = fpr_of(int64(samp(samp_ctx, x0, sigma)))
		*(*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*1)) = fpr_of(int64(samp(samp_ctx, x1, sigma)))
		return
	}
	n = uint64(1 << logn)
	hn = n >> 1
	tree0 = (*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n)))
	tree1 = (*fpr)(unsafe.Add(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(tree), unsafe.Sizeof(fpr(0))*uintptr(n)))), unsafe.Sizeof(fpr(0))*uintptr(ffLDL_treesize(logn-1))))
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(z1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(hn))), t1, logn)
	ffSampling_fft(samp, samp_ctx, tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), tree1, z1, (*fpr)(unsafe.Add(unsafe.Pointer(z1), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n))))
	PQCLEAN_FALCON512_CLEAN_poly_merge_fft(z1, tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn)
	libc.MemCpy(unsafe.Pointer(tmp), unsafe.Pointer(t1), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_sub(tmp, z1, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(tmp, tree, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(tmp, t0, logn)
	PQCLEAN_FALCON512_CLEAN_poly_split_fft(z0, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(hn))), tmp, logn)
	ffSampling_fft(samp, samp_ctx, tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), tree0, z0, (*fpr)(unsafe.Add(unsafe.Pointer(z0), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn-1, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(n))))
	PQCLEAN_FALCON512_CLEAN_poly_merge_fft(z0, tmp, (*fpr)(unsafe.Add(unsafe.Pointer(tmp), unsafe.Sizeof(fpr(0))*uintptr(hn))), logn)
}
func do_sign_tree(samp samplerZ, samp_ctx unsafe.Pointer, s2 *int16, expanded_key *fpr, hm *uint16, logn uint, tmp *fpr) int {
	var (
		n     uint64
		u     uint64
		t0    *fpr
		t1    *fpr
		tx    *fpr
		ty    *fpr
		b00   *fpr
		b01   *fpr
		b10   *fpr
		b11   *fpr
		tree  *fpr
		ni    fpr
		sqn   uint32
		ng    uint32
		s1tmp *int16
		s2tmp *int16
	)
	n = uint64(1 << logn)
	t0 = tmp
	t1 = (*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(n)))
	b00 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b00(logn))))
	b01 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b01(logn))))
	b10 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b10(logn))))
	b11 = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_b11(logn))))
	tree = (*fpr)(unsafe.Add(unsafe.Pointer(expanded_key), unsafe.Sizeof(fpr(0))*uintptr(skoff_tree(logn))))
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_of(int64(*(*uint16)(unsafe.Add(unsafe.Pointer(hm), unsafe.Sizeof(uint16(0))*uintptr(u)))))
	}
	PQCLEAN_FALCON512_CLEAN_FFT(t0, logn)
	ni = fpr_inverse_of_q
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer(t0), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t1, b01, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulconst(t1, fpr_neg(ni), logn)
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t0, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulconst(t0, ni, logn)
	tx = (*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*uintptr(n)))
	ty = (*fpr)(unsafe.Add(unsafe.Pointer(tx), unsafe.Sizeof(fpr(0))*uintptr(n)))
	ffSampling_fft(samp, samp_ctx, tx, ty, tree, t0, t1, logn, (*fpr)(unsafe.Add(unsafe.Pointer(ty), unsafe.Sizeof(fpr(0))*uintptr(n))))
	libc.MemCpy(unsafe.Pointer(t0), unsafe.Pointer(tx), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer(ty), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(tx, b00, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(ty, b10, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(tx, ty, logn)
	libc.MemCpy(unsafe.Pointer(ty), unsafe.Pointer(t0), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(ty, b01, logn)
	libc.MemCpy(unsafe.Pointer(t0), unsafe.Pointer(tx), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t1, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(t1, ty, logn)
	PQCLEAN_FALCON512_CLEAN_iFFT(t0, logn)
	PQCLEAN_FALCON512_CLEAN_iFFT(t1, logn)
	s1tmp = (*int16)(unsafe.Pointer(tx))
	sqn = 0
	ng = 0
	for u = 0; u < n; u++ {
		var z int32
		z = int32(int(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(hm), unsafe.Sizeof(uint16(0))*uintptr(u))))) - int(int32(fpr_rint(*(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(u)))))))
		sqn += uint32(int32(int(z) * int(z)))
		ng |= sqn
		*(*int16)(unsafe.Add(unsafe.Pointer(s1tmp), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(z)
	}
	sqn |= uint32(int32(-(int(ng) >> 31)))
	s2tmp = (*int16)(unsafe.Pointer(tmp))
	for u = 0; u < n; u++ {
		*(*int16)(unsafe.Add(unsafe.Pointer(s2tmp), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(-fpr_rint(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*uintptr(u)))))
	}
	if PQCLEAN_FALCON512_CLEAN_is_short_half(sqn, s2tmp, logn) != 0 {
		libc.MemCpy(unsafe.Pointer(s2), unsafe.Pointer(s2tmp), int(n*uint64(unsafe.Sizeof(int16(0)))))
		libc.MemCpy(unsafe.Pointer(tmp), unsafe.Pointer(s1tmp), int(n*uint64(unsafe.Sizeof(int16(0)))))
		return 1
	}
	return 0
}
func do_sign_dyn(samp samplerZ, samp_ctx unsafe.Pointer, s2 *int16, f *int8, g *int8, F *int8, G *int8, hm *uint16, logn uint, tmp *fpr) int {
	var (
		n     uint64
		u     uint64
		t0    *fpr
		t1    *fpr
		tx    *fpr
		ty    *fpr
		b00   *fpr
		b01   *fpr
		b10   *fpr
		b11   *fpr
		g00   *fpr
		g01   *fpr
		g11   *fpr
		ni    fpr
		sqn   uint32
		ng    uint32
		s1tmp *int16
		s2tmp *int16
	)
	n = uint64(1 << logn)
	b00 = tmp
	b01 = (*fpr)(unsafe.Add(unsafe.Pointer(b00), unsafe.Sizeof(fpr(0))*uintptr(n)))
	b10 = (*fpr)(unsafe.Add(unsafe.Pointer(b01), unsafe.Sizeof(fpr(0))*uintptr(n)))
	b11 = (*fpr)(unsafe.Add(unsafe.Pointer(b10), unsafe.Sizeof(fpr(0))*uintptr(n)))
	smallints_to_fpr(b01, f, logn)
	smallints_to_fpr(b00, g, logn)
	smallints_to_fpr(b11, F, logn)
	smallints_to_fpr(b10, G, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b01, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b00, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b11, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b10, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(b01, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(b11, logn)
	t0 = (*fpr)(unsafe.Add(unsafe.Pointer(b11), unsafe.Sizeof(fpr(0))*uintptr(n)))
	t1 = (*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(n)))
	libc.MemCpy(unsafe.Pointer(t0), unsafe.Pointer(b01), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(t0, logn)
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer(b00), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_muladj_fft(t1, b10, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(b00, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(b00, t0, logn)
	libc.MemCpy(unsafe.Pointer(t0), unsafe.Pointer(b01), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_muladj_fft(b01, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(b01, t1, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(b10, logn)
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer(b11), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(t1, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(b10, t1, logn)
	g00 = b00
	g01 = b01
	g11 = b10
	b01 = t0
	t0 = (*fpr)(unsafe.Add(unsafe.Pointer(b01), unsafe.Sizeof(fpr(0))*uintptr(n)))
	t1 = (*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(n)))
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_of(int64(*(*uint16)(unsafe.Add(unsafe.Pointer(hm), unsafe.Sizeof(uint16(0))*uintptr(u)))))
	}
	PQCLEAN_FALCON512_CLEAN_FFT(t0, logn)
	ni = fpr_inverse_of_q
	libc.MemCpy(unsafe.Pointer(t1), unsafe.Pointer(t0), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t1, b01, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulconst(t1, fpr_neg(ni), logn)
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t0, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mulconst(t0, ni, logn)
	libc.MemCpy(unsafe.Pointer(b11), unsafe.Pointer(t0), int(n*2*uint64(unsafe.Sizeof(fpr(0)))))
	t0 = (*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(n)))
	t1 = (*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(n)))
	ffSampling_fft_dyntree(samp, samp_ctx, t0, t1, g00, g01, g11, logn, logn, (*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*uintptr(n))))
	b00 = tmp
	b01 = (*fpr)(unsafe.Add(unsafe.Pointer(b00), unsafe.Sizeof(fpr(0))*uintptr(n)))
	b10 = (*fpr)(unsafe.Add(unsafe.Pointer(b01), unsafe.Sizeof(fpr(0))*uintptr(n)))
	b11 = (*fpr)(unsafe.Add(unsafe.Pointer(b10), unsafe.Sizeof(fpr(0))*uintptr(n)))
	libc.MemMove(unsafe.Pointer((*fpr)(unsafe.Add(unsafe.Pointer(b11), unsafe.Sizeof(fpr(0))*uintptr(n)))), unsafe.Pointer(t0), int(n*2*uint64(unsafe.Sizeof(fpr(0)))))
	t0 = (*fpr)(unsafe.Add(unsafe.Pointer(b11), unsafe.Sizeof(fpr(0))*uintptr(n)))
	t1 = (*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(n)))
	smallints_to_fpr(b01, f, logn)
	smallints_to_fpr(b00, g, logn)
	smallints_to_fpr(b11, F, logn)
	smallints_to_fpr(b10, G, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b01, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b00, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b11, logn)
	PQCLEAN_FALCON512_CLEAN_FFT(b10, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(b01, logn)
	PQCLEAN_FALCON512_CLEAN_poly_neg(b11, logn)
	tx = (*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*uintptr(n)))
	ty = (*fpr)(unsafe.Add(unsafe.Pointer(tx), unsafe.Sizeof(fpr(0))*uintptr(n)))
	libc.MemCpy(unsafe.Pointer(tx), unsafe.Pointer(t0), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	libc.MemCpy(unsafe.Pointer(ty), unsafe.Pointer(t1), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(tx, b00, logn)
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(ty, b10, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(tx, ty, logn)
	libc.MemCpy(unsafe.Pointer(ty), unsafe.Pointer(t0), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(ty, b01, logn)
	libc.MemCpy(unsafe.Pointer(t0), unsafe.Pointer(tx), int(n*uint64(unsafe.Sizeof(fpr(0)))))
	PQCLEAN_FALCON512_CLEAN_poly_mul_fft(t1, b11, logn)
	PQCLEAN_FALCON512_CLEAN_poly_add(t1, ty, logn)
	PQCLEAN_FALCON512_CLEAN_iFFT(t0, logn)
	PQCLEAN_FALCON512_CLEAN_iFFT(t1, logn)
	s1tmp = (*int16)(unsafe.Pointer(tx))
	sqn = 0
	ng = 0
	for u = 0; u < n; u++ {
		var z int32
		z = int32(int(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(hm), unsafe.Sizeof(uint16(0))*uintptr(u))))) - int(int32(fpr_rint(*(*fpr)(unsafe.Add(unsafe.Pointer(t0), unsafe.Sizeof(fpr(0))*uintptr(u)))))))
		sqn += uint32(int32(int(z) * int(z)))
		ng |= sqn
		*(*int16)(unsafe.Add(unsafe.Pointer(s1tmp), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(z)
	}
	sqn |= uint32(int32(-(int(ng) >> 31)))
	s2tmp = (*int16)(unsafe.Pointer(tmp))
	for u = 0; u < n; u++ {
		*(*int16)(unsafe.Add(unsafe.Pointer(s2tmp), unsafe.Sizeof(int16(0))*uintptr(u))) = int16(-fpr_rint(*(*fpr)(unsafe.Add(unsafe.Pointer(t1), unsafe.Sizeof(fpr(0))*uintptr(u)))))
	}
	if PQCLEAN_FALCON512_CLEAN_is_short_half(sqn, s2tmp, logn) != 0 {
		libc.MemCpy(unsafe.Pointer(s2), unsafe.Pointer(s2tmp), int(n*uint64(unsafe.Sizeof(int16(0)))))
		libc.MemCpy(unsafe.Pointer(tmp), unsafe.Pointer(s1tmp), int(n*uint64(unsafe.Sizeof(int16(0)))))
		return 1
	}
	return 0
}
func PQCLEAN_FALCON512_CLEAN_gaussian0_sampler(p *prng) int {
	var (
		dist [54]uint32 = [54]uint32{10745844, 3068844, 3741698, 5559083, 1580863, 8248194, 2260429, 13669192, 2736639, 708981, 4421575, 10046180, 169348, 7122675, 4136815, 30538, 13063405, 7650655, 4132, 14505003, 7826148, 417, 16768101, 11363290, 31, 8444042, 8086568, 1, 12844466, 265321, 0, 1232676, 13644283, 0, 38047, 9111839, 0, 870, 6138264, 0, 14, 12545723, 0, 0, 3104126, 0, 0, 28824, 0, 0, 198, 0, 0, 1}
		v0   uint32
		v1   uint32
		v2   uint32
		hi   uint32
		lo   uint64
		u    uint64
		z    int
	)
	lo = prng_get_u64(p)
	hi = uint32(prng_get_u8(p))
	v0 = uint32(int32(int(uint32(lo)) & 0xFFFFFF))
	v1 = uint32(int32(int(uint32(lo>>24)) & 0xFFFFFF))
	v2 = uint32(int32(int(uint32(lo>>48)) | int(hi)<<16))
	z = 0
	for u = 0; u < uint64((unsafe.Sizeof([54]uint32{}))/unsafe.Sizeof(uint32(0))); u += 3 {
		var (
			w0 uint32
			w1 uint32
			w2 uint32
			cc uint32
		)
		w0 = dist[u+2]
		w1 = dist[u+1]
		w2 = dist[u+0]
		cc = uint32(int32((int(v0) - int(w0)) >> 31))
		cc = uint32(int32((int(v1) - int(w1) - int(cc)) >> 31))
		cc = uint32(int32((int(v2) - int(w2) - int(cc)) >> 31))
		z += int(cc)
	}
	return z
}
func BerExp(p *prng, x fpr, ccs fpr) int {
	var (
		s  int
		i  int
		r  fpr
		sw uint32
		w  uint32
		z  uint64
	)
	s = int(fpr_trunc(PQCLEAN_FALCON512_CLEAN_fpr_mul(x, fpr_inv_log2)))
	r = fpr_sub(x, PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_of(int64(s)), fpr_log2))
	sw = uint32(int32(s))
	sw ^= uint32(int32((int(sw) ^ 63) & (-((63 - int(sw)) >> 31))))
	s = int(sw)
	z = ((PQCLEAN_FALCON512_CLEAN_fpr_expm_p63(r, ccs) << 1) - 1) >> uint64(s)
	i = 64
	for {
		i -= 8
		w = uint32(prng_get_u8(p) - uint(int(uint32(z>>uint64(i)))&0xFF))
		if int(w) != 0 || i <= 0 {
			break
		}
	}
	return int(w) >> 31
}
func PQCLEAN_FALCON512_CLEAN_sampler(ctx unsafe.Pointer, mu fpr, isigma fpr) int {
	var (
		spc *sampler_context
		s   int
		r   fpr
		dss fpr
		ccs fpr
	)
	spc = (*sampler_context)(ctx)
	s = int(fpr_floor(mu))
	r = fpr_sub(mu, fpr_of(int64(s)))
	dss = fpr_half(fpr_sqr(isigma))
	ccs = PQCLEAN_FALCON512_CLEAN_fpr_mul(isigma, spc.sigma_min)
	for {
		var (
			z0 int
			z  int
			b  int
			x  fpr
		)
		z0 = PQCLEAN_FALCON512_CLEAN_gaussian0_sampler(&spc.p)
		b = int(prng_get_u8(&spc.p)) & 1
		z = b + ((b<<1)-1)*z0
		x = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_sqr(fpr_sub(fpr_of(int64(z)), r)), dss)
		x = fpr_sub(x, PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_of(int64(z0*z0)), fpr_inv_2sqrsigma0))
		if BerExp(&spc.p, x, ccs) != 0 {
			return s + z
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_sign_tree(sig *int16, rng *shake256incctx, expanded_key *fpr, hm *uint16, logn uint, tmp *uint8) {
	var ftmp *fpr
	ftmp = (*fpr)(unsafe.Pointer(tmp))
	for {
		var (
			spc      sampler_context
			samp     samplerZ
			samp_ctx unsafe.Pointer
		)
		spc.sigma_min = fpr_sigma_min[logn]
		PQCLEAN_FALCON512_CLEAN_prng_init(&spc.p, rng)
		samp = func(ctx unsafe.Pointer, mu fpr, sigma fpr) int {
			return func(ctx unsafe.Pointer, mu fpr, sigma fpr) int {
				return PQCLEAN_FALCON512_CLEAN_sampler(ctx, mu, sigma)
			}(ctx, mu, sigma)
		}
		samp_ctx = unsafe.Pointer(&spc)
		if do_sign_tree(samp, samp_ctx, sig, expanded_key, hm, logn, ftmp) != 0 {
			break
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_sign_dyn(sig *int16, rng *shake256incctx, f *int8, g *int8, F *int8, G *int8, hm *uint16, logn uint, tmp *uint8) {
	var ftmp *fpr
	ftmp = (*fpr)(unsafe.Pointer(tmp))
	for {
		var (
			spc      sampler_context
			samp     samplerZ
			samp_ctx unsafe.Pointer
		)
		spc.sigma_min = fpr_sigma_min[logn]
		PQCLEAN_FALCON512_CLEAN_prng_init(&spc.p, rng)
		samp = func(ctx unsafe.Pointer, mu fpr, sigma fpr) int {
			return func(ctx unsafe.Pointer, mu fpr, sigma fpr) int {
				return PQCLEAN_FALCON512_CLEAN_sampler(ctx, mu, sigma)
			}(ctx, mu, sigma)
		}
		samp_ctx = unsafe.Pointer(&spc)
		if do_sign_dyn(samp, samp_ctx, sig, f, g, F, G, hm, logn, ftmp) != 0 {
			break
		}
	}
}
