import sys

line = sys.stdin.readlines()


def main(line):
    numStudents, passing_score = line[0].rstrip('\r\n').split()
    numStudents = int(numStudents)
    passing_score = int(passing_score)
    line.pop(0)
    for i in range(len(line)):
        if passing_score == 0:
            print(i+1)
        else:
            score, absence = line[i].rstrip('\r\n').split()
            score = int(score)
            absence = int(absence)
            sub_score = 0
            if absence > 0:
                sub_score = absence * 5
            if score - sub_score >= passing_score:
                print(i+1)


if __name__ == '__main__':
    main(line)
