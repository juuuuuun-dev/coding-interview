import sys


def main(line):
    n, r = (int(x) for x in line[0].rstrip("\r\n").split())
    boxs = make_arr(line, n)
    filterd_boxs = filterd_box(boxs, r)
    for i in filterd_boxs:
        print(i)
    

def make_arr(line, n):
    tmp = []
    for i in range(n):
        tmp.append([int(x) for x in line[i+1].split()])
    return tmp


def filterd_box(boxs, r):
    tmp = []
    for i in range(len(boxs)):
        is_in = True
        for n in boxs[i]:
            if n < r * 2:
                is_in = False
        if is_in:
            tmp.append(i + 1)
    return tmp


if __name__ == '__main__':
    line = sys.stdin.readlines()
    main(line)
