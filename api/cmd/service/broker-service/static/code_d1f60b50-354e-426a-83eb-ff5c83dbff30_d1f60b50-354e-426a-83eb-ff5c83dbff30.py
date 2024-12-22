# User function logic goes here
def main(d, value):
    return {k: v for k, v in d.items() if v != value}

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2, 'c': 1}, 1),
        ({'x': 5, 'y': 5, 'z': 10}, 5)
    ]
    expected_outputs = [
        {'b': 2},
        {'z': 10}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)