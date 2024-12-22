# User function logic goes here
def main(s):
    if '@' in s and '.' in s.split('@')[-1]:
        return True
    return False

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('test@example.com',),
        ('user@domain',),
        ('test@com',),
        ('user@domain.com',)
    ]
    expected_outputs = [
        True,
        False,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)