//question informations: https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150


func removeElement(nums []int, value int) int {
  res := 0
  for i, _ := range nums{
    if nums[i]!=value{
    nums[res]=nums[i]
    res++
      }
  }
  return res
}
