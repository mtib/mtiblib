package numbers

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	ans := []([6]float64){
		// num, n0, m0, n1, m1, res
		{0, 0, 100, 100, 200, 100},
		{100, 0, 100, 100, 200, 200},
		{50, 0, 100, 0, 1000, 500},
		{0, 0, 100, 0, -100, 0},
		{0, 0, 100, -100, 0, -100},
		{20, 0, 100, -100, 0, -80},
		{20, 0, 100, 0, -100, -20},
	}
	for _, v := range ans {
		res := Map(v[0], v[1], v[2], v[3], v[4])
		if res != v[5] {
			fmt.Printf("%8.2f != %8.2f\n", res, v[5])
			t.Fail()
		} else {
			fmt.Printf("%8.2f == %8.2f\n", res, v[5])
		}
	}
}
