package util

import (
	"fmt"
)

func main() {
	var N, m int
	fmt.Scan(&N)
	fmt.Scan(&m)
	item := make([][]int, m)
	var v, p, q int
	for i := 0; i < m; i++ {
		fmt.Scan(&v)
		fmt.Scan(&p)
		fmt.Scan(&q)
		item[i] = append(item[i], v, p, q)
	}
	for i := 1; i < m; i++ {
		if item[i][2] != 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if item[j][2] == 0 {
				break
			}
			t1 := item[j][0]
			t2 := item[j][1]
			t3 := item[j][2]
			item[j][0] = item[j+1][0]
			item[j][1] = item[j+1][1]
			item[j][2] = item[j+1][2]
			item[j+1][0] = t1
			item[j+1][1] = t2
			item[j+1][2] = t3

		}
	}
	//fmt.Println(item)
	var best int = 0
	solve(item, N, m, &best, 0, 0)
	fmt.Println(best)
}
func solve(item [][]int, N int, m int, best *int, now int, k int) {
	if k == m {
		if now > *best {
			// fmt.Println(now)
			*best = now
		}
		return
	}
	if N < item[k][0] {
		solve(item, N, m, best, now, k+1)
	} else if item[k][2] == 0 || item[item[k][2]-1][2] < 0 {
		if item[k][2] == 0 {
			item[k][2] = -1
		}
		solve(item, N-item[k][0], m, best, now+item[k][0]*item[k][1], k+1)
		if item[k][2] == -1 {
			item[k][2] = 0
		}
		solve(item, N, m, best, now, k+1)
	} else {
		solve(item, N, m, best, now, k+1)
	}

}
