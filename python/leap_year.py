import sys

def is_leap_year(year: int):
    if year % 4 != 0:
        return False
    if year % 100 == 0:
        return year % 400 == 0
    return True

year = input('Please enter the year: ').strip()

try:
    num_year = int(year)
except ValueError as e:
    print(f'\'{year}\' is not integer')
    sys.exit(1)

if is_leap_year(num_year):
    print(f'\'{year}\' is leap year')
else:
    print(f'\'{year}\' is not leap year')
