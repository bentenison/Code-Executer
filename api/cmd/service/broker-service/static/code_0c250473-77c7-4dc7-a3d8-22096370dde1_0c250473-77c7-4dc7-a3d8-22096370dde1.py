# User function logic goes here
def main(n):
    i = 2
    while i * i <= n:
        if n % i == 0:
            n //= i
        else:
            i += 1
    return n

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 13195, 'expectedOutput': 29},
        {'input': 49, 'expectedOutput': 7},
        {'input': 97, 'expectedOutput': 97},
        {'input': 28, 'expectedOutput': 7}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)