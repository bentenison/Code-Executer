# User function logic goes here
def main(d):
    return len(d) == 0

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({},),
        ({'a': 1},)
    ]
    expected_outputs = [
        True,
        False
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)