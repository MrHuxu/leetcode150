// code
/**
 * // Definition for a Node.
 * function Node(val,next,random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */
/**
 * @param {Node} head
 * @return {Node}
 */
let copyRandomList = function(head) {
  const visited = {};
    
  const visit = node => {
    if (visited[node.val]) return visited[node.val];
  
    const newNode = new Node(node.val);
    visited[node.val] = newNode;
    newNode.next = visit(node.next);
    newNode.random = visit(node.random);
    return newNode;
  };
  
  return visit(head);
};