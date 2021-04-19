
def main():
    n = 13
    ran = range(2, n)
    print(ran)
    for p in range(2, n):
        if n % p == 0:
            print(str(n) + ' is composite.')
            return
    print(str(n) + ' is PRIME!!')


if __name__ == '__main__':
    main()
