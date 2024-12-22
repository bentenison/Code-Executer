# User function logic goes here
def main(keys, values):
    return dict(zip(keys, values))

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        (['a', 'b', 'c'], [1, 2, 3]),
        (['x', 'y', 'z'], [4, 5, 6])
    ]
    expected_outputs = [
        {'a': 1, 'b': 2, 'c': 3},
        {'x': 4, 'y': 5, 'z': 6}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)