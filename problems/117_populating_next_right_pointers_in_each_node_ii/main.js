function Node(val,left,right,next) {
  this.val = val;
  this.left = left;
  this.right = right;
  this.next = next;
};

// code
/**
 * // Definition for a Node.
 * function Node(val,left,right,next) {
 *    this.val = val;
 *    this.left = left;
 *    this.right = right;
 *    this.next = next;
 * };
 */
/**
 * @param {Node} root
 * @return {Node}
 */
let connect = function(root) {
  if (null == root) return root;

  let head = new Node();
  head.next = root;
  while (head.next) {
    const nextHead = new Node();
    let tmp = nextHead;

    curr = head.next;
    while (curr) {
      if (curr.left) {
        tmp.next = curr.left;
        tmp = tmp.next;
      }
      if (curr.right) {
        tmp.next = curr.right;
        tmp = tmp.next;
      }
      curr = curr.next;
    }

    head = nextHead;
  }
  
  return root; 
};
