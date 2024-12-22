# User function logic goes here
def main(dict1, dict2):
    dict1.update(dict2)
    return dict1

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1}, {'b': 2}),
        ({'x': 5, 'y': 10}, {'z': 15, 'a': 1})
    ]
    expected_outputs = [
        {'a': 1, 'b': 2},
        {'x': 5, 'y': 10, 'z': 15, 'a': 1}
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print('Test Passed')