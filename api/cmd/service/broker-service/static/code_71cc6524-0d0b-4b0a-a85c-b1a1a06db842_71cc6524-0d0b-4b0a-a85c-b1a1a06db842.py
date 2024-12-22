# User function logic goes here
def main(s):
    return len(s)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': {1, 2, 3}, 'expectedOutput': 3},
        {'input': {'a', 'b', 'c', 'd'}, 'expectedOutput': 4},
        {'input': set(), 'expectedOutput': 0},
        {'input': {10, 20, 30, 40, 50}, 'expectedOutput': 5}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
            print(f'Failed for Input: {case["input"]}. Expected: {case["expectedOutput"]}, Got: {result}')
    print(all_passed)