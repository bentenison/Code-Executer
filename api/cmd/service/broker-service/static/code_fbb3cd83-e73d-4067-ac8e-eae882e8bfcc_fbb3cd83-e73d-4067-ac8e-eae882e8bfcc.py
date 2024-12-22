# User function logic goes here
from collections import Counter

def find_most_frequent(d):
    all_values = [item for sublist in d.values() for item in sublist]
    return Counter(all_values).most_common(1)[0][0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ({'a': [1, 2, 3], 'b': [2, 3, 3]},),
        ({'x': [1, 1], 'y': [2, 3]},),
    ]
    expected_outputs = [
        3,
        1,
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_most_frequent(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)