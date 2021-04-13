package main

import (
	"fmt"
	"strconv"
)

type SpeSkillTest struct{}

func main() {
	var test *SpeSkillTest
	//NARASIST
	c := test.narcissistic(153)
	fmt.Println(c)

	//PARITY
	n := []int{160, 3, 1719, 19, 11, 13, -21}
	x := test.ParityOutlier(n)
	fmt.Println(x)

	//FIND
	d := test.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue")
	fmt.Println(d)

	//BLUEOCEAN
	v := test.BlueOcean([]int{1, 2, 3, 4, 6, 10}, 1)
	fmt.Println(v)
}

func (c *SpeSkillTest) narcissistic(n int) bool {
	totall := 0
	for i := n; i > 0; i /= 10 {
		totall++
	}
	temp := 0
	for x := n; x > 0; x /= 10 {
		angka := x % 10
		kuadrat := 1
		for i := 1; i <= totall; i++ {
			kuadrat *= angka
		}
		temp += kuadrat
	}
	if n == temp {
		return true
	} else {
		return false
	}
}

func (c *SpeSkillTest) ParityOutlier(n []int) string {
	var genap, ganjil []int
	for _, a := range n {
		// fmt.Println(a)
		if a%2 == 0 {
			genap = append(genap, a)
		} else {
			ganjil = append(ganjil, a)
		}
	}
	if len(ganjil) == len(genap) {
		return "FALSE"
	} else {
		if len(genap) < len(ganjil) {
			return strconv.Itoa(genap[0]) + ",Satu satunya genap"
		} else {
			return strconv.Itoa(ganjil[0]) + ",Satu satunya ganjil"
		}
	}
}

func (c *SpeSkillTest) FindNeedle(n []string, search string) (x int) {
	for i, q := range n {
		if q == search {
			x = i
			break
		}
	}
	return
}

func (c *SpeSkillTest) BlueOcean(data []int, hapus int) []int {
	newdata := []int{}

	for _, row := range data {
		if row != hapus {
			newdata = append(newdata, row)
		}
	}

	return newdata
}
