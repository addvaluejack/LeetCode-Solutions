func detectCapitalUse(word string) bool {
    first_capital := false
    other_capital := false
    capital_count := 0
    
    for i := 0; i < len(word); i++ {
        if int(word[i]) >= int('A') && int(word[i]) <= int('Z') {
            capital_count++
            if i == 0 {
                first_capital = true
            } else {
                other_capital = true
            }
        }
    }
    
    if capital_count == len(word) {
        return true
    }
    if first_capital == false && other_capital == false {
        return true
    }
    if first_capital == true && other_capital == false {
        return true
    }
    return false
}
