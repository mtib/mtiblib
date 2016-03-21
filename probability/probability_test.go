package probability

import (
	"fmt"
	"math"
	"testing"
)

func TestHalffunc(t *testing.T) {
	Initialize()
	num := 100000
	allow := 0.05
	sum := 0
	for i := 0; i < num; i++ {
		if Half() {
			sum++
		}
	}
	deriv := num/2 - sum
	got := math.Abs(float64(deriv)) / float64(num)
	if got > allow {
		t.Fail()
	}
	fmt.Printf("expected 0.5, got: %.5f\n", math.Abs(float64(sum)/float64(num)))
}

func TestProb(t *testing.T) {
	Initialize()
	type test struct {
		size, prob, err float64
	}
	tests := []test{
		{300000, 1.0 / 3.0, .05},
		{400000, 0.25, .05},
		{300000, 2.0 / 3.0, .05},
		{400000, 0.75, 0.05},
		{1000, 1, 0},
		{1000, 0, 0},
		{1000, 0.99, 0.03},
	}
	for _, v := range tests {
		func(v test, t *testing.T) {
			if !tprob(v.size, v.prob, v.err) {
				t.Fail()
			}
		}(v, t)
	}
}

func tprob(size, prob, err float64) bool {
	sum := 0
	for i := 0; i < int(size); i++ {
		if Probability(prob) {
			sum++
		}
	}
	res := float64(sum) / size
	fmt.Printf("expected %.3f, got: %.5f\n", prob, res)
	if math.Abs(res-prob) > err {
		return false
	}
	return true
}
