# User function logic goes here
def main(d1, d2):
    return d1 == d2

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2}, {'a': 1, 'b': 2}),
        ({'a': 1, 'b': 2}, {'a': 1, 'c': 3})
    ]
    expected_outputs = [
        True,
        False
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)