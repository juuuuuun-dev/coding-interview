import sys


def main(lines):
    sort_num = []
    zero_count = 0
    zero_str = ""
    res = ""
    for i, v in enumerate(lines):
        # print("line[{0}]: {1}".format(i, v))
        sort_num = sorted(v)
        zero_count = sort_num.count("0")
    for i in range(zero_count):
        zero_str += "0"
    for i in range(len(sort_num)):
        if sort_num[i] == '0':
            continue
        else:
            res += str(sort_num[i])
        if i == zero_count:
            res += str(zero_str)
    print(res)


if __name__ == '__main__':
    lines = []
    for l in sys.stdin:
        lines.append(l.rstrip('\r\n'))
    main(lines)
