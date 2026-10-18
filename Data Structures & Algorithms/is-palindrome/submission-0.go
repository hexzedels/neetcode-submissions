func isPalindrome(s string) bool {
    lo, hi := 0, len(s)-1

    for lo < hi {
        // isAlphanumeric
        if (s[lo] >= 'a' && s[lo] <= 'z') || (s[lo] >= 'A' && s[lo] <= 'Z') || (s[lo] >= '0' && s[lo] <= '9') {
             if (s[hi] >= 'a' && s[hi] <= 'z') || (s[hi] >= 'A' && s[hi] <= 'Z') || (s[hi] >= '0' && s[hi] <= '9') {
                if unicode.ToLower(rune(s[lo])) == unicode.ToLower(rune(s[hi])) {
                    lo++
                    hi--
                } else {
                    return false
                }
             } else {
                hi--
             }
        } else {
            lo++
        }
    }

    return true
}