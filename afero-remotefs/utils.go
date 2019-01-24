package remoteAfero

func min(x int, y int64) (a int) {
	a = int(y)
	// when gzip, the content-length=-1
	if x < a || a < 0 {
		a = x
	}
	return
}
