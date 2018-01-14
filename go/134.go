func canCompleteCircuit(gas []int, cost []int) int {
    result := -1
    
    for start_index := 0; start_index < len(gas); start_index++ {
        remaining_gas := 0
        flag := true
        for i := start_index; i < len(gas); i++ {
            remaining_gas += gas[i]
            if cost[i] > remaining_gas {
                flag = false
                break
            } else {
                remaining_gas -= cost[i]
            }
        }
        if flag == false {
            continue
        }
        for i := 0; i < start_index; i++ {
            remaining_gas += gas[i]
            if cost[i] > remaining_gas {
                flag = false
                break
            } else {
                remaining_gas -= cost[i]
            }
        }
        if flag == false {
            continue
        } else {
            return start_index
        }
    }
    
    return result
}
