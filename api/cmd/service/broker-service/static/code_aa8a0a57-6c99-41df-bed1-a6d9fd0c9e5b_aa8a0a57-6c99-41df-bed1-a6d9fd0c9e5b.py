# User function logic goes here
import string

def main(s):
    return ''.join(char for char in s if char not in string.punctuation)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'Hello, world!', 'expectedOutput': 'Hello world'},
        {'input': 'Python@3.9', 'expectedOutput': 'Python39'},
        {'input': 'Good morning!', 'expectedOutput': 'Good morning'},
        {'input': 'Remove this: ; punctuation.', 'expectedOutput': 'Remove this punctuation'}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
            print(f'Failed for Input: {case["input"]}. Expected: {case["expectedOutput"]}, Got: {result}')
    print(all_passed)