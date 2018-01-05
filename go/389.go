func findTheDifference(ransomNote string, magazine string) byte {
    var result byte
    ransomNote_count := make([]int, 256)
    magazine_count := make([]int, 256)
    
    for i := 0; i < len(ransomNote); i++ {
        ransomNote_count[int(ransomNote[i])]++
    }
    
    for i := 0; i < len(magazine); i++ {
        magazine_count[int(magazine[i])]++
    }
    
    for i := 0; i < 256; i++ {
        if ransomNote_count[i] < magazine_count[i] {
            return byte(i)
        }
    }
    
    return result
}
