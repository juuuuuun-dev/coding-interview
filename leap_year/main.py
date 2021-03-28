
def main():
    years = (4, 1000, 1992, 2000, 2001)
    for i in years:
        if i == 4:
            continue
        if i % 400 == 0 or (i % 4 == 0 and i % 100 != 0):
            print(f"{i} is a leap year")
        else:
            print(f"{i} is not a leap year")
            
main()
