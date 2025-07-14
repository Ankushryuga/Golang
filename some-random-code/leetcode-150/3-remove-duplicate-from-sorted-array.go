package main
//question information: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150


func removeDuplicate(nums []int) int {
  res := 0
  numsLen := len(nums)
  for i:=0;i<numsLen-1;i++{
    if nums[i]!=nums[i+1]{
      nums[res]=nums[i];
      res++
    }
  }
  nums[res]=nums[numsLen-1]
  res++
  return res
}
