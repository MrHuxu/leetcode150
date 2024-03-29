export class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

export const buildTree = (vals: (number | null)[]): TreeNode | null => {
    if (!vals.length) return null;

    const root = new TreeNode(vals[0]!);
    let nodes: TreeNode[] = [root];
    vals = vals.slice(1);
    while (vals.length) {
        const node = nodes.shift()!;

        if (vals[0] !== null && vals[0] !== undefined) {
            node.left = new TreeNode(vals[0]!);
            nodes.push(node.left);
        }
        if (vals[1] !== null && vals[1] !== undefined) {
            node.right = new TreeNode(vals[1]!);
            nodes.push(node.right);
        }
        vals = vals.slice(2);
    }

    return root;
};
