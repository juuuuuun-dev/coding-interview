from typing import List


class Solution:
    def repeatedNTimes(self, A: List[int]):
        count = dict()
        for i in A:
            if i in count:
                return i
            count[i] = 1


if __name__ == '__main__':
    list = [1, 2, 3, 3]
    solution = Solution()
    res = solution.repeatedNTimes(list)
    print(res)
