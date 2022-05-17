package question_117;

import static org.assertj.core.api.Assertions.assertThat;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void connect() {
        Node node1 = new Node(1);
        Node node2 = new Node(2);
        Node node3 = new Node(3);
        Node node4 = new Node(4);
        Node node5 = new Node(5);
        Node node7 = new Node(7);
        node1.left = node2;
        node1.right = node3;
        node2.left = node4;
        node2.right = node5;
        node3.right = node7;

        Node node11 = new Node(1);
        Node node12 = new Node(2);
        Node node13 = new Node(3);
        Node node14 = new Node(4);
        Node node15 = new Node(5);
        Node node17 = new Node(7);
        node11.left = node12;
        node11.right = node13;
        node12.left = node14;
        node12.right = node15;
        node12.next = node13;
        node13.right = node17;
        node14.next = node15;
        node15.next = node17;

        new Solution().connect(node1);
        assertThat(node1).usingRecursiveComparison().isEqualTo(node11);
    }
}
