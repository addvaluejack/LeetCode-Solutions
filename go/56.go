/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func QuickSort(slice []Interval) []Interval {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]Interval, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]Interval, 0, length)
	middle := make([]Interval, 0, length)
	more := make([]Interval, 0, length)

	for _, item := range slice {
		switch {
		case item.Start < m.Start:
			less = append(less, item)
		case item.Start == m.Start:
			middle = append(middle, item)
		case item.Start > m.Start:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}

func merge(intervals []Interval) []Interval {
    number_of_intervals := len(intervals)
    
    intervals = QuickSort(intervals)
    
    for i := 0; i < number_of_intervals-1; {
        if intervals[i].End >= intervals[i+1].Start {
            if intervals[i].End < intervals[i+1].End {
                intervals[i].End = intervals[i+1].End
            }
            intervals = append(intervals[:i+1], intervals[i+2:]...)
            number_of_intervals--
        } else {
            i++
        }
    }
    
    return intervals
}
