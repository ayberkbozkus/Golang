func isValid(s string) bool {
    stack := make([]rune,0)
    values := map[rune]rune {
        '(': ')', 
        '[': ']',
        '{': '}',
    }
    for _, char := range s {
        if _, ok := values[char]; ok { 
            stack = append(stack,values[char])
        } else {
            if len(stack) > 0 && (char == stack[len(stack)-1]){
                stack = stack[:len(stack)-1]
            } else{
                return false
            }
        }
    } 
    if len(stack) > 0 {
        return false
    }
    return true
}