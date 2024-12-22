# User function logic goes here
def main(d):
    return {v: k for k, v in d.items()}

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2, 'c': 3},),
        ({'x': 1, 'y': 2},)
    ]
    expected_outputs = [
        {1: 'a', 2: 'b', 3: 'c'},
        {1: 'x', 2: 'y'}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)