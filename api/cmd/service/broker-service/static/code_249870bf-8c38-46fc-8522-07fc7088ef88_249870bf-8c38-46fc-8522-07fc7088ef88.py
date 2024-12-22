# User function logic goes here
def main(start, end):
    return [num for num in range(start, end + 1) if num % 2 == 0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': (1, 10), 'expectedOutput': [2, 4, 6, 8, 10]},
        {'input': (4, 4), 'expectedOutput': [4]},
        {'input': (5, 10), 'expectedOutput': [6, 8, 10]},
        {'input': (10, 20), 'expectedOutput': [10, 12, 14, 16, 18, 20]}
    ]
    for case in test_cases:
        result = main(case['input'][0], case['input'][1])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)