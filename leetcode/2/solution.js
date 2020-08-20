/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */

var addTwoNumbers = function (x, y) {
  const sum = new ListNode();

  let node = sum;
  for (let ds = 0, carry = 0; _canRepeat(x, y, carry); node = node.next) {
    let dx = x !== null ? x.val : 0;
    let dy = y !== null ? y.val : 0;

    ds = dx + dy + carry;
    carry = ds > 9 ? 1 : 0;

    node.val = ds % 10;

    if (x !== null) x = x.next;
    if (y !== null) y = y.next;
    if (_canRepeat(x, y, carry)) node.next = new ListNode();
  }

  return sum;
};

function _canRepeat(x, y, carry) {
  return x !== null || y !== null || carry !== 0;
}
