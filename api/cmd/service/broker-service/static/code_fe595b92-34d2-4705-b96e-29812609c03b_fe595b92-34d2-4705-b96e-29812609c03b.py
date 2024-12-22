# User function logic goes here
def is_rotation(s1, s2):
    return len(s1) == len(s2) and s2 in (s1 + s1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abc', 'cab'),
        ('hello', 'ohell'),
        ('abc', 'abcd'),
        ('', '')
    ]
    expected_outputs = [
        True,
        True,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_rotation(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)