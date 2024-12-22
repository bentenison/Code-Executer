# User function logic goes here
def main(s):
    return s[::2]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'abcdef', 'expectedOutput': 'ace'},
        {'input': 'hello', 'expectedOutput': 'hlo'},
        {'input': 'python', 'expectedOutput': 'pto'},
        {'input': 'world', 'expectedOutput': 'wrd'}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
            print(f'Failed for Input: {case["input"]}. Expected: {case["expectedOutput"]}, Got: {result}')
    print(all_passed)