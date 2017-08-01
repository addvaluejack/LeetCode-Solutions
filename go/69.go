func mySqrt(x int) int {
    a := x
    for ; a*a > x; {
        a = (a + x/a)/2
    }
    return a
}
