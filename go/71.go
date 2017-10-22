import (
    "bytes"
)

func simplifyPath(path string) string {
    var result bytes.Buffer
    path_stack := make([]string, 100)
    path_stack_size := 0
    
    for i := 0; i < len(path); i++ {
        if path[i] == '/' {
            continue
        } else if path[i] == '.' && i+1 == len(path) {
            continue
        } else if path[i] == '.' && i+1 < len(path) && path[i+1] == '/' {
            continue
        } else if path[i] == '.' && i+1 < len(path) && path[i+1] == '.' && i+2 == len(path) {
            if path_stack_size > 0 {
                path_stack_size--            
            }
            i++
            continue
        } else if path[i] == '.' && i+1 < len(path) && path[i+1] == '.' && i+2 < len(path) && path[i+2] == '/' {
            if path_stack_size > 0 {
                path_stack_size--            
            }
            i++
            continue
        } else {
            j := i
            for ; j < len(path); j++ {
                if path[j] == '/' {
                    break
                }
            }
            path_stack[path_stack_size] = path[i:j]
            path_stack_size++
            i = j
        }
    }
    
    if path_stack_size == 0 {
        return "/"
    }
    
    for i := 0; i < path_stack_size; i++ {
        result.WriteString("/")
        result.WriteString(path_stack[i])
    }
    
    return result.String()
}
