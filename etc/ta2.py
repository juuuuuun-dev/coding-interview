import itertools
import sys


def main(lines):
    l = int(lines[0])
    lines.pop(0)
    if l >= 3:
        print("YES")
        for i, v in enumerate(lines):
            nums = (int(x) for x in v.split())
            combi = list(itertools.combinations(nums, 3))
            total = sum(nums) / 2
            # for n in nums:
        for i in combi:
            print(i)
    else:
        print("NO")


if __name__ == '__main__':
    lines = []
    for l in sys.stdin:
        lines.append(l.rstrip('\r\n'))
    main(lines)
