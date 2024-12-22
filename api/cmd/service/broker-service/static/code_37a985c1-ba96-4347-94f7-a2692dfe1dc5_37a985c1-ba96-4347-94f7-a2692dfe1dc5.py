# User function logic goes here
def check_case(c):
    if c.isupper():
        return 'Uppercase'
    return 'Lowercase'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('A',),
        ('z',),
        ('a',)
    ]
    expected_outputs = [
        'Uppercase',
        'Lowercase',
        'Lowercase'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_case(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)