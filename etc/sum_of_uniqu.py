from typing import List


class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        return sum([x for x in nums if nums.count(x) == 1])


if __name__ == '__main__':
    list = [1, 2, 3, 5, 1]
    solution = Solution()
    res = solution.sumOfUnique(list)
    print(res)
