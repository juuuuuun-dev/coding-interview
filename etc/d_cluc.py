input_line = input()


# def cluc(input_line):
#     default_price = 100
#     distance_price = 10
#     price = int(input_line) * distance_price
#     return default_price + price
    
    
# if __name__ == '__main__':
#     print(cluc(input_line))

def cluc(input_line):
    n, m = input_line.split()
    return (int(n) * 6000) + (int(m) * 4000)


if __name__ == '__main__':
    print(cluc(input_line))
