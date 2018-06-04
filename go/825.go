import "math/rand"

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range slice {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func check(A int, B int) int {
    checkA := 1
    checkB := 1
    if float64(B) <= 0.5*float64(A)+7 {
        checkA = 0
    }
    if B > A {
        checkA = 0
    }
    if B > 100 && A < 100 {
        checkA = 0
    }
    if float64(A) <= 0.5*float64(B)+7 {
        checkB = 0
    }
    if A > B {
        checkB = 0
    }
    if A > 100 && B < 100 {
        checkB = 0
    }
    return checkA+checkB
}

func sort(ages []int) ([]int, []int) {
    new_ages := QuickSort(ages)
    index := 0
    a := make([]int, 0)
    b := make([]int, 0)
    for i := 0; i < len(new_ages); i++ {
        a = append(a, new_ages[i])
        b = append(b, 1)
        for j := i+1; j < len(new_ages) && new_ages[j] == new_ages[i]; j++ {
            i++
            b[index]++
        }
        index++
    }
    return a, b
}

func numFriendRequests(ages []int) int {
    count := 0
    a, b := sort(ages)
    for i := 0; i < len(a); i++ {
        for j := i; j < len(a); j++ {
            if i == j {
                count += check(a[i], a[j])*((1+(b[i]-1))*(b[i]-1)/2)
            } else {
                count += check(a[i], a[j])*b[i]*b[j]
            }
        }
    }
    return count
}
