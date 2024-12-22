# User function logic goes here
from collections import Counter

def odd_frequency_chars(s):
    freq_map = Counter(s)
    return [char for char, count in freq_map.items() if count % 2 != 0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('aabcc',),
        ('aaabbbb',),
        ('abcde',),
        ('',)
    ]
    expected_outputs = [
        ['a', 'c'],
        ['a', 'b'],
        ['a', 'b', 'c', 'd', 'e'],
        []
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = odd_frequency_chars(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)