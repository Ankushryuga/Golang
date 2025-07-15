package main

//question informations: https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

func reverse(nums []int, i, j int){
    for j>i{
        temp := nums[i]
        nums[i]=nums[j]
        nums[j]=temp
        i++
        j--
    }
}
func rotate(nums []int, k int)  {
    numsLen := len(nums)
    k = k%numsLen
    reverse(nums, 0, numsLen-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, numsLen-1)
}
