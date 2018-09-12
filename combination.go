package util

func combination(n, r int, work []int, result *[][]int) {

	if r == 0 {
		var workCopy []int
		if len(work) == 0 {
			workCopy = append(workCopy, 0)
		} else {
			for _, w := range work {
				workCopy = append(workCopy, w)
			}
		}
		*result = append(*result, workCopy)
		return
	}
	for i := 0; i < n; i += 1 {
		work = append(work, i)
		combination(i, r-1, work, result)
		work = (work)[:len(work)-1]
	}

}
