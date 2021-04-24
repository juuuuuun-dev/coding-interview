import sys


def main(line, decoration):
    txt = line[0].rstrip('\r\n')
    txt_len = len(txt)
    border_len = 2 + txt_len
    decoration_str = ""
    for i in range(border_len):
        decoration_str += decoration
    print(decoration_str)
    print(f"{decoration}{txt}{decoration}")
    print(decoration_str)
    
if __name__ == '__main__':
    line = sys.stdin.readlines()
    decoration = "+"
    main(line, decoration)
