package numbers

func Map(num, min0, max0, min1, max1 float64) float64 {
	return num/(max0-min0)*(max1-min1) + min1
}
