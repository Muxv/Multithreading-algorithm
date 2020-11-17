package mulalg

func MergeSort(a []int) {
	n := len(a)
	mergeMainLoop(a, 0, n-1)
}

func MergeSortParallel(a []int) {
	n := len(a)
	mergeMainLoopParallel(a, 0, n-1)
}

func mergeMainLoop(a []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeMainLoop(a, start, mid)
		mergeMainLoop(a, mid+1, end)
		merge(a, start, end, mid)
	} else {
		return
	}
}

func mergeMainLoopParallel(a []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		ch := make(chan int)
		go func(sync chan int) {
			mergeMainLoop(a, start, mid)
			ch <- 1
		}(ch)
		go func(sync chan int) {
			mergeMainLoop(a, mid+1, end)
			ch <- 1
		}(ch)
		<-ch
		<-ch
		merge(a, start, end, mid)
	} else {
		return
	}
}

func merge(a []int, start int, end int, mid int) {
	//fmt.Printf("start : %d, end : %d, mid : %d\n", start, end, mid)
	l := end - start + 1
	res := make([]int, l)
	leftFlag, rightFlag := start, mid+1
	i := 0
	//fmt.Printf("Lflag : %d, Rflag: %d\n", leftFlag, rightFlag)
	for (leftFlag != mid+1) && (rightFlag != end+1) {
		if a[leftFlag] <= a[rightFlag] {
			res[i] = a[leftFlag]
			leftFlag++
		} else {
			res[i] = a[rightFlag]
			rightFlag++
		}
		i++
	}
	if leftFlag == mid+1 {
		for ; i < l; i++ {
			res[i] = a[rightFlag]
			rightFlag++
		}
	} else {
		for ; i < l; i++ {
			res[i] = a[leftFlag]
			leftFlag++
		}
	}
	copy(a[start:end+1], res)
}
