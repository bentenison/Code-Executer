# User function logic goes here
def find_keys_above_threshold(d, threshold):
    return [k for k, v in d.items() if v > threshold]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ({'a': 1, 'b': 5, 'c': 10}, 4),
        ({'x': 2, 'y': 7, 'z': 3}, 4),
    ]
    expected_outputs = [
        ['b', 'c'],
        ['y'],
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_keys_above_threshold(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)