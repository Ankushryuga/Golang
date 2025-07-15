package main
//question infromations:https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicateII(nums []int) int{
  res := 0
  for i, _ := range nums{
    if res<2 || num[i]>nums[res-2]{
      nums[res]=nums[i]
      res++
    }
  }
  return res
}
