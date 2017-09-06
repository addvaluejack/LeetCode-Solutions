func twoSum(numbers []int, target int) []int {
    var result []int
    index1 := 0
    index2 := 1
    
    for ; index1 < len(numbers); index1++ {
        index2 = index1+1
        if numbers[index1] > target || index1 == len(numbers)-1 {
            return result;
        }
        for ; index2 < len(numbers) && numbers[index1]+numbers[index2] < target; index2++ {
        }
        if index2 < len(numbers) && numbers[index1]+numbers[index2] == target {
           result = append(result, index1+1);
           result = append(result, index2+1);
           break;
        }
    }
    
    return result
}
