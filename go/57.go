/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
    var result []Interval
    
    if len(intervals) == 0 {
        return append(result, newInterval)
    }
    
    if newInterval.End < intervals[0].Start {
        result = append(result, newInterval)
        result = append(result, intervals...)
        return result
    }
    
    if newInterval.Start > intervals[len(intervals)-1].End {
        return append(intervals, newInterval)
    } 

    for i := 0; i < len(intervals)-1; i++ {
        if newInterval.Start > intervals[i].End && newInterval.End < intervals[i+1].Start {
            result = append(result, intervals[:i+1]...)
            result = append(result, newInterval)
            result = append(result, intervals[i+1:]...)
            return result
        }
    }
    
    x := 0
    
    for ; x < len(intervals); x++ {
        if newInterval.Start <= intervals[x].End {
            break
        }
    }
    
    if newInterval.Start > intervals[x].Start {
        newInterval.Start = intervals[x].Start
    }
    
    y := x
    
    for ; y < len(intervals); y++ {
        if newInterval.End < intervals[y].Start {
            break
        }
        if newInterval.End < intervals[y].End {
            newInterval.End = intervals[y].End
        }
    }
    
    result = append(result, intervals[:x]...)
    result = append(result, newInterval)
    result = append(result, intervals[y:]...)
    
    return result
}
