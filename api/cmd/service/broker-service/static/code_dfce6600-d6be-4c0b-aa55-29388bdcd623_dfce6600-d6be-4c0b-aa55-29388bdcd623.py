# User function logic goes here
def sort_by_keys(d):
    return (dict(sorted(d.items())), dict(sorted(d.items(), reverse=True)))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ({'b': 1, 'a': 2, 'c': 3},),
        ({'x': 4, 'z': 2, 'y': 1},),
    ]
    expected_outputs = [
        ({'a': 2, 'b': 1, 'c': 3}, {'c': 3, 'b': 1, 'a': 2}),
        ({'y': 1, 'z': 2, 'x': 4}, {'x': 4, 'z': 2, 'y': 1}),
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = sort_by_keys(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)