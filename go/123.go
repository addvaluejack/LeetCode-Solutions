import("fmt")

func subMax(values []int) int {
    tmp := 0
    result := 0
    
    for i := 0; i < len(values); i++ {
        if tmp+values[i] > 0 {
            tmp += values[i]
            if tmp > result {
                result = tmp
            }
        } else {
            tmp = 0
        }
    }
    
    return result
}

func maxProfit(prices []int) int {
    if len(prices) <= 1 {
        return 0
    }
    
    result := 0
    diff := make([]int, len(prices)-1)
    
    for i := 0; i < len(diff); i++ {
        diff[i] = prices[i+1]-prices[i]
        //fmt.Printf("%d\n", diff[i])
    }
    
    for i := 0; i < len(diff); i++ {
        if subMax(diff[:i])+subMax(diff[i:]) > result {
            fmt.Printf("%d\n", i)
            result = subMax(diff[:i])+subMax(diff[i:])
        }
    }
    
    return result
}
