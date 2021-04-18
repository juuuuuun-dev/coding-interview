import sys

line = sys.stdin.readlines()


def main(line):
    point = 0
    for i in range(len(line)):
        if i == 0:
            start, m = line[i].split()
        else:
            n = int(line[i])
            if n < point:
                point -= n
            else:
                start = int(start) - n
                point += int(n * 0.1)
            print(f"{start} {point}")


if __name__ == '__main__':
    main(line)
