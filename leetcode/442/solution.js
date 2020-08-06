// 442. Find All Duplicates in an Array
// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?

var findDuplicates = function (nums) {
  let res = [];
  for (let i = 0; i < nums.length; i++) {
    while (nums[i] != nums[nums[i] - 1]) {
      let temp = nums[i];
      nums[i] = nums[nums[i] - 1];
      nums[temp - 1] = temp;
    }
  }

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] != i + 1) {
      res.push(nums[i]);
    }
  }

  return res;
};
