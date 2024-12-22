# User function logic goes here
import re

def main(s):
    s = re.sub('[^a-zA-Z0-9]', '', s).lower()
    return s == s[::-1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('A man, a plan, a canal, Panama',),
        ('race a car',),
        ('Hello',),
        ('No lemon, no melon',)
    ]
    expected_outputs = [
        True,
        False,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)