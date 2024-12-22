# User function logic goes here
def main(d, key):
    if key in d:
        del d[key]
    return d

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2}, 'b'),
        ({'x': 5, 'y': 10, 'z': 15}, 'y')
    ]
    expected_outputs = [
        {'a': 1},
        {'x': 5, 'z': 15}
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print('Test Passed')