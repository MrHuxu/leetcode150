package question_20;

import java.util.Stack;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (Character ch : s.toCharArray()) {
            if (stack.isEmpty()) {
                stack.push(ch);
                continue;
            }

            switch (ch) {
                case ')':
                    if (stack.peek() == '(')
                        stack.pop();
                    else
                        stack.push(ch);
                    break;

                case ']':
                    if (stack.peek() == '[')
                        stack.pop();
                    else
                        stack.push(ch);
                    break;

                case '}':
                    if (stack.peek() == '{')
                        stack.pop();
                    else
                        stack.push(ch);
                    break;

                default:
                    stack.push(ch);
            }
        }
        return stack.size() == 0;
    }
}