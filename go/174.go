func calculateMinimumHP(dungeon [][]int) int {
    hp := make([][]int, len(dungeon))
    for i := 0; i < len(dungeon); i++ {
        hp[i] = make([]int, len(dungeon[0]))
    }
    if dungeon[len(dungeon)-1][len(dungeon[0])-1] < 0 {
        hp[len(dungeon)-1][len(dungeon[0])-1] = -dungeon[len(dungeon)-1][len(dungeon[0])-1]+1
    } else {
        hp[len(dungeon)-1][len(dungeon[0])-1] = 1 
    }
    for i := len(dungeon)-1; i >= 0; i-- {
        for j := len(dungeon[0])-1; j >= 0; j-- {
            if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
                continue
            }
            var t int
            if i == len(dungeon)-1 {
                t = hp[i][j+1]
            } else if j == len(dungeon[0])-1 {
                t = hp[i+1][j]
            } else {
                if hp[i][j+1] > hp[i+1][j] {
                    t = hp[i+1][j]
                } else {
                    t = hp[i][j+1]
                }
            }
            if dungeon[i][j] <= 0 {
                hp[i][j] = t-dungeon[i][j]
            } else {
                if dungeon[i][j] >= t {
                    hp[i][j] = 1
                } else {
                    hp[i][j] = t-dungeon[i][j]
                }
            }
        }
    }
    return hp[0][0]
}
