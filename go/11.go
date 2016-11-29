import "math"

func maxArea(height []int) int {
    max := 0;
    
    for i := 0; i < len(height); i++ {
        current_max := 0;
        for j := 0; j < len(height); j++ {
            if(i != j && height[i] <= height[j]) {
                if(i > j) {
                    if(height[i]*(i-j) > current_max) {
                        current_max = height[i]*(i-j);
                    }
                } else {
                    if(height[i]*(j-i) > current_max) {
                        current_max = height[i]*(j-i);
                    }
                }
            }
        }
        if(current_max > max) {
            max = current_max;
        }
    }
    
    return max;
}
