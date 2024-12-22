# User function logic goes here
import math

def main(radius):
    return round(math.pi * radius ** 2, 2)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5, 78.54),
        (10, 314.16),
        (7, 153.94),
        (2, 12.57)
    ]
    for radius, expected in test_cases:
        result = main(radius)
        if round(result, 2) != expected:
            all_passed = False
    print(all_passed)