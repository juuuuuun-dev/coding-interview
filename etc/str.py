input_line = input()


def leet(input_line):
    strs = list(input_line)
    leet_dict = {
        "A": 4,
        "E": 3,
        "G": 6,
        "I": 1,
        "O": 0,
        "S": 5,
        "Z": 2,
    }
    temp = []
    for i in strs:
        if i in leet_dict:
            temp.append(str(leet_dict[i]))
        else:
            temp.append(str(i))
    return ''.join(temp)


if __name__ == '__main__':
    print(leet(input_line))
