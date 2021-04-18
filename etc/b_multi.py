import sys

line = sys.stdin.readlines()


def main(line):
    for i in range(len(line)):
        l = line[i].rstrip('\r\n')
        es = ves = ies = None
        if not l.isalpha():
            continue
        es = checkES(l)
        ves = checkVES(l)
        ies = checkIES(l)
        if es:
            print(es)
        elif ves:
            print(ves)
        elif ies:
            print(ies)
        elif not ies and not es and not ves:
            print(l + "s")


def is_str(v):
    return type(v) is str


def checkES(s):
    add_str = "es"
    one = s[-1:]
    two = s[-2:]
    if one == "s" or one == "o" or one == "x" or two == "sh" or two == "ch":
        return s + add_str


def checkIES(s) -> None:
    one = s[-1:]
    if one == "y":
        two = s[-2:-1]
        if two == 'a' or two == "i" or two == "u" or two == "e" or two == "o":
            return s + "s"
        else:
            return s[:-1] + "ies"


def checkVES(s) -> None:
    add_str = "ves"
    one = s[-1:]
    two = s[-2:]
    if one == "f":
        return s[:-1] + add_str
    elif two == "fe":
        return s[:-2] + add_str


if __name__ == '__main__':
    main(line)
