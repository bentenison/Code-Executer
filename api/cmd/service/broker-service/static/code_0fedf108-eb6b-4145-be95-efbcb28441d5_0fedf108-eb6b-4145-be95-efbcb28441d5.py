# User function logic goes here
def main(n):
    factor = 2
    largest_prime = 1
    while n > 1:
        if n % factor == 0:
            largest_prime = factor
            n //= factor
        else:
            factor += 1
    return largest_prime

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 15, 'expectedOutput': 5},
        {'input': 13195, 'expectedOutput': 29},
        {'input': 29, 'expectedOutput': 29},
        {'input': 2, 'expectedOutput': 2}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)