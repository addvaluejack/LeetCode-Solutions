import (
    "strconv"
)

func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    
    for i := 0; i < len(tokens); i++ {
        if tokens[i] == "+" {
            stack[len(stack)-2] = stack[len(stack)-2]+stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else if tokens[i] == "-" {
            stack[len(stack)-2] = stack[len(stack)-2]-stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else if tokens[i] == "*" {
            stack[len(stack)-2] = stack[len(stack)-2]*stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else if tokens[i] == "/" {
            stack[len(stack)-2] = stack[len(stack)-2]/stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        } else {
            num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
        }
    } 
    return stack[0]
}
