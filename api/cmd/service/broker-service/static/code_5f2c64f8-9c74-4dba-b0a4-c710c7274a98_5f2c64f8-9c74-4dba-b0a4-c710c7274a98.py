# User function logic goes here
from collections import defaultdict

def specific_char_frequency(str_list, chars):
    freq_map = defaultdict(int)
    for s in str_list:
        for char in s:
            if char in chars:
                freq_map[char] += 1
    return dict(freq_map)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (['hello', 'world'], ['h', 'e', 'l']),
        (['abc', 'def', 'ghi'], ['a', 'b', 'c']),
        (['aaa', 'bbb', 'ccc'], ['a', 'b', 'c']),
        ([], ['a'])
    ]
    expected_outputs = [
        {'h': 1, 'e': 1, 'l': 3},
        {'a': 1, 'b': 1, 'c': 1},
        {'a': 3, 'b': 3, 'c': 3},
        {} 
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = specific_char_frequency(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)