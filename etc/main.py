import sys


def main(lines):
    for i, v in enumerate(lines):
        print(f"Hello {v}!")


if __name__ == '__main__':
    lines = []
    for line in sys.stdin:
        lines.append(line.rstrip('\r\n'))
    main(lines)
