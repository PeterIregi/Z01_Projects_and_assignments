package level_4

func WeAreUnique(str1, str2 string) int {
    // Edge case
    if str1 == "" && str2 == "" {
        return -1
    }

    set1 := make(map[rune]bool)
    set2 := make(map[rune]bool)

    // Fill sets (unique characters only)
    for _, ch := range str1 {
        set1[ch] = true
    }

    for _, ch := range str2 {
        set2[ch] = true
    }

    count := 0

    // Characters in str1 but not in str2
    for ch := range set1 {
        if !set2[ch] {
            count++
        }
    }

    // Characters in str2 but not in str1
    for ch := range set2 {
        if !set1[ch] {
            count++
        }
    }

    return count
}