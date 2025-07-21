package falcon512

import "unsafe"

func PQCLEAN_FALCON512_CLEAN_FFT(f *fpr, logn uint) {
	var (
		u  uint
		t  uint64
		n  uint64
		hn uint64
		m  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	t = hn
	for func() uint64 {
		u = 1
		return func() uint64 {
			m = 2
			return m
		}()
	}(); u < logn; func() uint64 {
		u++
		return func() uint64 {
			m <<= 1
			return m
		}()
	}() {
		var (
			ht uint64
			hm uint64
			i1 uint64
			j1 uint64
		)
		ht = t >> 1
		hm = m >> 1
		for func() uint64 {
			i1 = 0
			return func() uint64 {
				j1 = 0
				return j1
			}()
		}(); i1 < hm; func() uint64 {
			i1++
			return func() uint64 {
				j1 += t
				return j1
			}()
		}() {
			var (
				j  uint64
				j2 uint64
			)
			j2 = j1 + ht
			var s_re fpr
			var s_im fpr
			s_re = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((m+i1)<<1)+0]
			s_im = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((m+i1)<<1)+1]
			for j = j1; j < j2; j++ {
				var (
					x_re fpr
					x_im fpr
					y_re fpr
					y_im fpr
				)
				x_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j)))
				x_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+hn)))
				y_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+ht)))
				y_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+ht+hn)))
				for {
					{
						var (
							fpct_a_re fpr
							fpct_a_im fpr
							fpct_b_re fpr
							fpct_b_im fpr
							fpct_d_re fpr
							fpct_d_im fpr
						)
						fpct_a_re = y_re
						fpct_a_im = y_im
						fpct_b_re = s_re
						fpct_b_im = s_im
						fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
						fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
						y_re = fpct_d_re
						y_im = fpct_d_im
					}
					if true {
						break
					}
				}
				for {
					{
						var (
							fpct_re fpr
							fpct_im fpr
						)
						fpct_re = PQCLEAN_FALCON512_CLEAN_fpr_add(x_re, y_re)
						fpct_im = PQCLEAN_FALCON512_CLEAN_fpr_add(x_im, y_im)
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j))) = fpct_re
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+hn))) = fpct_im
					}
					if true {
						break
					}
				}
				for {
					{
						var (
							fpct_re fpr
							fpct_im fpr
						)
						fpct_re = fpr_sub(x_re, y_re)
						fpct_im = fpr_sub(x_im, y_im)
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+ht))) = fpct_re
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+ht+hn))) = fpct_im
					}
					if true {
						break
					}
				}
			}
		}
		t = ht
	}
}
func PQCLEAN_FALCON512_CLEAN_iFFT(f *fpr, logn uint) {
	var (
		u  uint64
		n  uint64
		hn uint64
		t  uint64
		m  uint64
	)
	n = uint64(1 << logn)
	t = 1
	m = n
	hn = n >> 1
	for u = uint64(logn); u > 1; u-- {
		var (
			hm uint64
			dt uint64
			i1 uint64
			j1 uint64
		)
		hm = m >> 1
		dt = t << 1
		for func() uint64 {
			i1 = 0
			return func() uint64 {
				j1 = 0
				return j1
			}()
		}(); j1 < hn; func() uint64 {
			i1++
			return func() uint64 {
				j1 += dt
				return j1
			}()
		}() {
			var (
				j  uint64
				j2 uint64
			)
			j2 = j1 + t
			var s_re fpr
			var s_im fpr
			s_re = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((hm+i1)<<1)+0]
			s_im = fpr_neg(PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((hm+i1)<<1)+1])
			for j = j1; j < j2; j++ {
				var (
					x_re fpr
					x_im fpr
					y_re fpr
					y_im fpr
				)
				x_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j)))
				x_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+hn)))
				y_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+t)))
				y_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+t+hn)))
				for {
					{
						var (
							fpct_re fpr
							fpct_im fpr
						)
						fpct_re = PQCLEAN_FALCON512_CLEAN_fpr_add(x_re, y_re)
						fpct_im = PQCLEAN_FALCON512_CLEAN_fpr_add(x_im, y_im)
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j))) = fpct_re
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+hn))) = fpct_im
					}
					if true {
						break
					}
				}
				for {
					{
						var (
							fpct_re fpr
							fpct_im fpr
						)
						fpct_re = fpr_sub(x_re, y_re)
						fpct_im = fpr_sub(x_im, y_im)
						x_re = fpct_re
						x_im = fpct_im
					}
					if true {
						break
					}
				}
				for {
					{
						var (
							fpct_a_re fpr
							fpct_a_im fpr
							fpct_b_re fpr
							fpct_b_im fpr
							fpct_d_re fpr
							fpct_d_im fpr
						)
						fpct_a_re = x_re
						fpct_a_im = x_im
						fpct_b_re = s_re
						fpct_b_im = s_im
						fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
						fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+t))) = fpct_d_re
						*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(j+t+hn))) = fpct_d_im
					}
					if true {
						break
					}
				}
			}
		}
		t = dt
		m = hm
	}
	if logn > 0 {
		var ni fpr
		ni = PQCLEAN_FALCON512_CLEAN_fpr_p2_tab[logn]
		for u = 0; u < n; u++ {
			*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(u))), ni)
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_add(a *fpr, b *fpr, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_add(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))), *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_sub(a *fpr, b *fpr, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_sub(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))), *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_neg(a *fpr, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_neg(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_adj_fft(a *fpr, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = n >> 1; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_neg(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_mul_fft(a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u)))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = a_re
				fpct_a_im = a_im
				fpct_b_re = b_re
				fpct_b_im = b_im
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpct_d_re
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpct_d_im
			}
			if true {
				break
			}
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_muladj_fft(a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u)))
		b_im = fpr_neg(*(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u+hn))))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = a_re
				fpct_a_im = a_im
				fpct_b_re = b_re
				fpct_b_im = b_im
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpct_d_re
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpct_d_im
			}
			if true {
				break
			}
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_mulselfadj_fft(a *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			a_re fpr
			a_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(a_re), fpr_sqr(a_im))
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpr_zero
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_mulconst(a *fpr, x fpr, logn uint) {
	var (
		n uint64
		u uint64
	)
	n = uint64(1 << logn)
	for u = 0; u < n; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))), x)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_div_fft(a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u)))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
					fpct_m    fpr
				)
				fpct_a_re = a_re
				fpct_a_im = a_im
				fpct_b_re = b_re
				fpct_b_im = b_im
				fpct_m = PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(fpct_b_re), fpr_sqr(fpct_b_im))
				fpct_m = fpr_inv(fpct_m)
				fpct_b_re = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_b_re, fpct_m)
				fpct_b_im = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_neg(fpct_b_im), fpct_m)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpct_d_re
				*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpct_d_im
			}
			if true {
				break
			}
		}
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_invnorm2_fft(d *fpr, a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u)))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		*(*fpr)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_inv(PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(a_re), fpr_sqr(a_im)), PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(b_re), fpr_sqr(b_im))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_add_muladj_fft(d *fpr, F *fpr, G *fpr, f *fpr, g *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			F_re fpr
			F_im fpr
			G_re fpr
			G_im fpr
			f_re fpr
			f_im fpr
			g_re fpr
			g_im fpr
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
		)
		F_re = *(*fpr)(unsafe.Add(unsafe.Pointer(F), unsafe.Sizeof(fpr(0))*uintptr(u)))
		F_im = *(*fpr)(unsafe.Add(unsafe.Pointer(F), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		G_re = *(*fpr)(unsafe.Add(unsafe.Pointer(G), unsafe.Sizeof(fpr(0))*uintptr(u)))
		G_im = *(*fpr)(unsafe.Add(unsafe.Pointer(G), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		f_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(u)))
		f_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		g_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = F_re
				fpct_a_im = F_im
				fpct_b_re = f_re
				fpct_b_im = fpr_neg(f_im)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				a_re = fpct_d_re
				a_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = G_re
				fpct_a_im = G_im
				fpct_b_re = g_re
				fpct_b_im = fpr_neg(g_im)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				b_re = fpct_d_re
				b_im = fpct_d_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, b_re)
		*(*fpr)(unsafe.Add(unsafe.Pointer(d), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, b_im)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_mul_autoadj_fft(a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))), *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u))))
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))), *(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u))))
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_div_autoadj_fft(a *fpr, b *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var ib fpr
		ib = fpr_inv(*(*fpr)(unsafe.Add(unsafe.Pointer(b), unsafe.Sizeof(fpr(0))*uintptr(u))))
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u))), ib)
		*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = PQCLEAN_FALCON512_CLEAN_fpr_mul(*(*fpr)(unsafe.Add(unsafe.Pointer(a), unsafe.Sizeof(fpr(0))*uintptr(u+hn))), ib)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_LDL_fft(g00 *fpr, g01 *fpr, g11 *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			g00_re fpr
			g00_im fpr
			g01_re fpr
			g01_im fpr
			g11_re fpr
			g11_im fpr
			mu_re  fpr
			mu_im  fpr
		)
		g00_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g00_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		g01_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g01_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		g11_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g11_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
					fpct_m    fpr
				)
				fpct_a_re = g01_re
				fpct_a_im = g01_im
				fpct_b_re = g00_re
				fpct_b_im = g00_im
				fpct_m = PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(fpct_b_re), fpr_sqr(fpct_b_im))
				fpct_m = fpr_inv(fpct_m)
				fpct_b_re = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_b_re, fpct_m)
				fpct_b_im = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_neg(fpct_b_im), fpct_m)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				mu_re = fpct_d_re
				mu_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = mu_re
				fpct_a_im = mu_im
				fpct_b_re = g01_re
				fpct_b_im = fpr_neg(g01_im)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				g01_re = fpct_d_re
				g01_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = fpr_sub(g11_re, g01_re)
				fpct_im = fpr_sub(g11_im, g01_im)
				*(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpct_re
				*(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpct_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u))) = mu_re
		*(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpr_neg(mu_im)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_LDLmv_fft(d11 *fpr, l10 *fpr, g00 *fpr, g01 *fpr, g11 *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	for u = 0; u < hn; u++ {
		var (
			g00_re fpr
			g00_im fpr
			g01_re fpr
			g01_im fpr
			g11_re fpr
			g11_im fpr
			mu_re  fpr
			mu_im  fpr
		)
		g00_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g00_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g00), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		g01_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g01_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g01), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		g11_re = *(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u)))
		g11_im = *(*fpr)(unsafe.Add(unsafe.Pointer(g11), unsafe.Sizeof(fpr(0))*uintptr(u+hn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
					fpct_m    fpr
				)
				fpct_a_re = g01_re
				fpct_a_im = g01_im
				fpct_b_re = g00_re
				fpct_b_im = g00_im
				fpct_m = PQCLEAN_FALCON512_CLEAN_fpr_add(fpr_sqr(fpct_b_re), fpr_sqr(fpct_b_im))
				fpct_m = fpr_inv(fpct_m)
				fpct_b_re = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_b_re, fpct_m)
				fpct_b_im = PQCLEAN_FALCON512_CLEAN_fpr_mul(fpr_neg(fpct_b_im), fpct_m)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				mu_re = fpct_d_re
				mu_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = mu_re
				fpct_a_im = mu_im
				fpct_b_re = g01_re
				fpct_b_im = fpr_neg(g01_im)
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				g01_re = fpct_d_re
				g01_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = fpr_sub(g11_re, g01_re)
				fpct_im = fpr_sub(g11_im, g01_im)
				*(*fpr)(unsafe.Add(unsafe.Pointer(d11), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpct_re
				*(*fpr)(unsafe.Add(unsafe.Pointer(d11), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpct_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(l10), unsafe.Sizeof(fpr(0))*uintptr(u))) = mu_re
		*(*fpr)(unsafe.Add(unsafe.Pointer(l10), unsafe.Sizeof(fpr(0))*uintptr(u+hn))) = fpr_neg(mu_im)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_split_fft(f0 *fpr, f1 *fpr, f *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		qn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	qn = hn >> 1
	*(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*0)) = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*0))
	*(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*0)) = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(hn)))
	for u = 0; u < qn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
			t_re fpr
			t_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+0)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+0+hn)))
		b_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+1)))
		b_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+1+hn)))
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, b_re)
				fpct_im = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, b_im)
				t_re = fpct_re
				t_im = fpct_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_half(t_re)
		*(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*uintptr(u+qn))) = fpr_half(t_im)
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = fpr_sub(a_re, b_re)
				fpct_im = fpr_sub(a_im, b_im)
				t_re = fpct_re
				t_im = fpct_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = t_re
				fpct_a_im = t_im
				fpct_b_re = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((u+hn)<<1)+0]
				fpct_b_im = fpr_neg(PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((u+hn)<<1)+1])
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				t_re = fpct_d_re
				t_im = fpct_d_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*uintptr(u))) = fpr_half(t_re)
		*(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*uintptr(u+qn))) = fpr_half(t_im)
	}
}
func PQCLEAN_FALCON512_CLEAN_poly_merge_fft(f *fpr, f0 *fpr, f1 *fpr, logn uint) {
	var (
		n  uint64
		hn uint64
		qn uint64
		u  uint64
	)
	n = uint64(1 << logn)
	hn = n >> 1
	qn = hn >> 1
	*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*0)) = *(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*0))
	*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr(hn))) = *(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*0))
	for u = 0; u < qn; u++ {
		var (
			a_re fpr
			a_im fpr
			b_re fpr
			b_im fpr
			t_re fpr
			t_im fpr
		)
		a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*uintptr(u)))
		a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f0), unsafe.Sizeof(fpr(0))*uintptr(u+qn)))
		for {
			{
				var (
					fpct_a_re fpr
					fpct_a_im fpr
					fpct_b_re fpr
					fpct_b_im fpr
					fpct_d_re fpr
					fpct_d_im fpr
				)
				fpct_a_re = *(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*uintptr(u)))
				fpct_a_im = *(*fpr)(unsafe.Add(unsafe.Pointer(f1), unsafe.Sizeof(fpr(0))*uintptr(u+qn)))
				fpct_b_re = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((u+hn)<<1)+0]
				fpct_b_im = PQCLEAN_FALCON512_CLEAN_fpr_gm_tab[((u+hn)<<1)+1]
				fpct_d_re = fpr_sub(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_re), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_im))
				fpct_d_im = PQCLEAN_FALCON512_CLEAN_fpr_add(PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_re, fpct_b_im), PQCLEAN_FALCON512_CLEAN_fpr_mul(fpct_a_im, fpct_b_re))
				b_re = fpct_d_re
				b_im = fpct_d_im
			}
			if true {
				break
			}
		}
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = PQCLEAN_FALCON512_CLEAN_fpr_add(a_re, b_re)
				fpct_im = PQCLEAN_FALCON512_CLEAN_fpr_add(a_im, b_im)
				t_re = fpct_re
				t_im = fpct_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+0))) = t_re
		*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+0+hn))) = t_im
		for {
			{
				var (
					fpct_re fpr
					fpct_im fpr
				)
				fpct_re = fpr_sub(a_re, b_re)
				fpct_im = fpr_sub(a_im, b_im)
				t_re = fpct_re
				t_im = fpct_im
			}
			if true {
				break
			}
		}
		*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+1))) = t_re
		*(*fpr)(unsafe.Add(unsafe.Pointer(f), unsafe.Sizeof(fpr(0))*uintptr((u<<1)+1+hn))) = t_im
	}
}
