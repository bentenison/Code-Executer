# User function logic goes here
def swap_first_last(d):
    keys = list(d.keys())
    d[keys[0]], d[keys[-1]] = d[keys[-1]], d[keys[0]]
    return d

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ({'a': 1, 'b': 2, 'c': 3},),
        ({'x': 10, 'y': 20, 'z': 30},),
    ]
    expected_outputs = [
        {'c': 3, 'b': 2, 'a': 1},
        {'z': 30, 'y': 20, 'x': 10},
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = swap_first_last(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)