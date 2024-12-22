# User function logic goes here
from collections import defaultdict

def frequency_of_numbers_in_string(s):
    freq_map = defaultdict(int)
    for char in s:
        if char.isdigit():
            freq_map[char] += 1
    return dict(freq_map)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('a1b2c3d1',),
        ('112233',),
        ('abcd',),
        ('12345',)
    ]
    expected_outputs = [
        {'1': 2, '2': 1, '3': 1},
        {'1': 2, '2': 2, '3': 2},
        {},
        {'1': 1, '2': 1, '3': 1, '4': 1, '5': 1}
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = frequency_of_numbers_in_string(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)