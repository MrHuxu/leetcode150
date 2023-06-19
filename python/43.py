import unittest


class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == '0' or num2 == '0':
            return '0'

        tmp = [0]
        for i in reversed(range(len(num2))):
            for j in reversed(range(len(num1))):
                curr_pos = (len(num2) - 1 - i) + (len(num1) - 1 - j)
                next_pos = (len(num2) - 1 - i) + (len(num1) - 1 - j) + 1

                product = int(num2[i]) * int(num1[j])

                if curr_pos > len(tmp) - 1:
                    tmp.append(product)
                else:
                    tmp[curr_pos] += product

                if tmp[curr_pos] > 9:
                    if next_pos > len(tmp) - 1:
                        tmp.append(tmp[curr_pos] // 10)
                    else:
                        tmp[next_pos] += tmp[curr_pos] // 10
                    tmp[curr_pos] %= 10

        ret = ''
        for i in reversed(range(len(tmp))):
            ret += str(tmp[i])
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual('6', solution.multiply('2', '3'))
