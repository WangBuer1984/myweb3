package task2

func MyPoint(p *int) int {
	p1 := *p
	p1 += 10
	return p1
}

func MySlicePoint(slicePtr *[]int) {
	tmpslicePtr := *slicePtr
	for i := range tmpslicePtr {
		tmpslicePtr[i] = tmpslicePtr[i] * 2

	}

}

func MyGoroutine() {

}
