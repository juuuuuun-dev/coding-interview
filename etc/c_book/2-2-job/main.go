package main

/*
n の仕事数
つづく行 最初が開始時間 空白のあとに終了時間で一行一つの仕事
もっとも多くの仕事を組み合わせる

考え方
選べる仕事の中で、終了時間が最も早いものを選ぶ事を繰り返す
終了:開始のmapを作成しsort
*/
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func split(s string) []string {
	return strings.Fields(s)
}

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func splitLine(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
}

type jobType map[int]int

func buildJobs(n int, lines [][]string) (jobType, []int) {
	jobs := make(map[int]int)
	// mapのforはランダムだからkeyを作成
	var jobKeys []int
	for i := 0; i < n; i++ {
		key := strToInt(lines[i][1])
		jobs[key] = strToInt(lines[i][0])
		jobKeys = append(jobKeys, key)
	}
	sort.Sort(sort.IntSlice(jobKeys))
	return jobs, jobKeys
}

func calc(n int, jobs jobType, jobKeys []int) {
	ans, lastTime := 0, 0
	for _, v := range jobKeys {
		// 開始時間が前回のジョブの終了時間より後なら
		if lastTime < jobs[v] {
			// 最後のジョブの終了時間を入れてカウント
			lastTime = v
			ans++
		}
	}
	fmt.Println(ans)
}

func main() {
	n := strToInt(readLine())
	lines := splitLine(n)
	jobs, jobKeys := buildJobs(n, lines)
	calc(n, jobs, jobKeys)
}
