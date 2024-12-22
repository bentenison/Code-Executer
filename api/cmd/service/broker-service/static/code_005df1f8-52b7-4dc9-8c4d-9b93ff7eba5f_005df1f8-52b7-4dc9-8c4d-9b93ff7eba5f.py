# User function logic goes here
def main(n):
    return sum(i**2 for i in range(1, n+1))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5, 55),
        (10, 385),
        (3, 14),
        (7, 140)
    ]
    for n, expected in test_cases:
        result = main(n)
        if result != expected:
            all_passed = False
    print(all_passed)