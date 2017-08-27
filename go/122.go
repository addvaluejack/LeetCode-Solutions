func maxProfit(prices []int) int {
    result := 0
    
    if len(prices) <= 1 {
        return 0
    }
    
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            result += prices[i]-prices[i-1]
        }
    }
    
    return result
}
