// code
/**
 * // Definition for a Node.
 * function Node(val,neighbors) {
 *    this.val = val;
 *    this.neighbors = neighbors;
 * };
 */
/**
 * @param {Node} node
 * @return {Node}
 */
let cloneGraph = function(node) {
  const visited = {};
  
  const visit = node => {
    if (visited[node.val]) return visited[node.val];

    const newNode = new Node(node.val, []);
    visited[node.val] = newNode;
    for (let neighbor of node.neighbors) {
      newNode.neighbors.push(visit(neighbor));
    }
    return newNode;
  };

  return visit(node);
};
