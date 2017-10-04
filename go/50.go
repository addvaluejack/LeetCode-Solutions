func foo(x float64, n int) float64 {
    if n == 1 {
        return x
    }
    if n%2 == 1 {
        result := foo(x, n/2)
        return result*result*x
    } else {
        result := foo(x, n/2)
        return result*result
    }
}

func myPow(x float64, n int) float64 {
    var result float64
    sign := false
    
    if n == 0 {
        result = 1
        return result
    }
    
    if n < 0 {
        n = -n
        sign = true
    }
    
    result = foo(x, n)
    
    if sign {
        return 1/result
    } else {
        return result
    }
}
