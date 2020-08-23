/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */

var findMedianSortedArrays = function (nums1, nums2) {
  const merged = merge(nums1, nums2);
  const medianPosition = merged.length;
  const isEven = medianPosition % 2 == 0;
  if (isEven) {
    return (merged[(medianPosition / 2) - 1] + merged[(medianPosition / 2)]) /
      2;
  } else {
    return merged[Math.floor(medianPosition / 2)];
  }
};

const merge = (a, b) => {
  let result = [];
  while (a.length > 0 && b.length > 0) {
    if (a[0] < b[0]) {
      result.push(a.shift());
    } else {
      result.push(b.shift());
    }
  }
  if (a.length > 0) return [...result, ...a];
  return [...result, ...b];
};
