# User function logic goes here
def main(start, end):
    return [num for num in range(start, end + 1) if num % 2 != 0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': (1, 10), 'expectedOutput': [1, 3, 5, 7, 9]},
        {'input': (3, 3), 'expectedOutput': [3]},
        {'input': (2, 9), 'expectedOutput': [3, 5, 7, 9]},
        {'input': (15, 20), 'expectedOutput': [15, 17, 19]}
    ]
    for case in test_cases:
        result = main(case['input'][0], case['input'][1])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)