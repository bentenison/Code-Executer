# User function logic goes here
def is_power_of_two(n):
    return n > 0 and (n & (n - 1)) == 0

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (16),
        (18)
    ]
    expected_outputs = [
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_power_of_two(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)