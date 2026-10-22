func maxArea(height []int) int {
    lo, hi := 0, len(height) - 1
    var area int

    for lo < hi {
        area = max(area, (hi-lo)*min(height[lo], height[hi]))

        if height[lo] <= height[hi] {
            lo++
        } else {
            hi--
        }
    }

    return area
}