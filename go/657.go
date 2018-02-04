func judgeCircle(moves string) bool {
    vertical := 0
    horizontal := 0
    
    for i := 0; i < len(moves); i++ {
        if moves[i] == 'U' {
            vertical++
        } else if moves[i] == 'D' {
            vertical--
        } else if moves[i] == 'L' {
            horizontal++
        } else {
            horizontal--
        }
    }
    
    return vertical == 0 && horizontal == 0
}
