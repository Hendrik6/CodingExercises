"use strict";

function _toConsumableArray(arr) { return _arrayWithoutHoles(arr) || _iterableToArray(arr) || _nonIterableSpread(); }

function _nonIterableSpread() { throw new TypeError("Invalid attempt to spread non-iterable instance"); }

function _iterableToArray(iter) { if (Symbol.iterator in Object(iter) || Object.prototype.toString.call(iter) === "[object Arguments]") return Array.from(iter); }

function _arrayWithoutHoles(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = new Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } }

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function findMedianSortedArrays(nums1, nums2) {
  var merged = merge(nums1, nums2);
  var medianPosition = merged.length;
  var isEven = medianPosition % 2 == 0;

  if (isEven) {
    return (merged[medianPosition / 2 - 1] + merged[medianPosition / 2]) / 2;
  } else {
    return merged[Math.floor(medianPosition / 2)];
  }
};

var merge = function merge(a, b) {
  var result = [];

  while (a.length > 0 && b.length > 0) {
    if (a[0] < b[0]) {
      result.push(a.shift());
    } else {
      result.push(b.shift());
    }
  }

  if (a.length > 0) return [].concat(result, _toConsumableArray(a));
  return [].concat(result, _toConsumableArray(b));
};