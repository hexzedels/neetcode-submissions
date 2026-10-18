func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))
    prefix := make([]int, len(nums))
    suffix := make([]int, len(nums))

    curProduct := 1
    for i, n := range nums {
        prefix[i] = curProduct
        curProduct *= n
    }

    curProduct = 1
    for i := len(nums) - 1; i >= 0; i -- {
        suffix[i] = curProduct
        curProduct *=nums[i]
    }

    for i := range nums {
        answer[i] = prefix[i]*suffix[i]
    }

    return answer
}