class Node133 {
    val: number
    neighbors: Node133[]
    constructor(val?: number, neighbors?: Node133[]) {
        this.val = (val === undefined ? 0 : val)
        this.neighbors = (neighbors === undefined ? [] : neighbors)
    }
}

/**
 * Definition for Node.
 * class Node {
 *     val: number
 *     neighbors: Node[]
 *     constructor(val?: number, neighbors?: Node[]) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.neighbors = (neighbors===undefined ? [] : neighbors)
 *     }
 * }
 */

function cloneGraph(node: Node133 | null): Node133 | null {
    return helper(node, new Map());
};

const helper = (node: Node133 | null, nodes: Map<number, Node133>): Node133 | null => {
    if (!node) return null;

    nodes.set(node.val, new Node133(node.val, []));
    for (let neighbor of node.neighbors) {
        if (!nodes.has(neighbor.val)) helper(neighbor, nodes);
        nodes.get(node.val)?.neighbors.push(nodes.get(neighbor.val)!);
    }

    return nodes.get(node.val)!;
}

test('133', () => {
    const node1 = new Node133(1);
    const node2 = new Node133(2);
    const node3 = new Node133(3);
    const node4 = new Node133(4);
    node1.neighbors = [node2, node4];
    node2.neighbors = [node1, node3];
    node3.neighbors = [node2, node4];
    node4.neighbors = [node1, node3];
    expect(cloneGraph(node1)).toStrictEqual(node1);
});