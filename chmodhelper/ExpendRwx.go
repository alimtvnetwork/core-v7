package chmodhelper

func ExpendRwx(rwx string) (r, w, x byte) {
	r = rwx[0]
	w = rwx[1]
	x = rwx[2]

	return r, w, x
}
