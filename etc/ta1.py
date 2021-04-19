from typing import List


def main(list) -> int:
    min = None
    for i in list:
        if not min and i > 0:
            min = i
        if i < min and i > 0:
            min = i
    return min


if __name__ == '__main__':
    nums = [2, -1, 0, 4, 5, 4, 3, 7, 99, 10]
    res = main(nums)
    print(res)
    print(0 > -1)
