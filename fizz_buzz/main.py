def main():
    numbers = range(1, 100)
    for i in numbers:
        if i % 3 == 0 and i % 5 == 0:
            print(f'Fizz Buzz {i}')
        if i % 3 == 0:
            print(f'Fizz {i}')
        if i % 5 == 0:
            print(f'Buzz {i}')
        else:
            print(i)
main()
