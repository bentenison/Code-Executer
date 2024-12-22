# User function logic goes here
def is_digits_only(string):
    return string.isdigit()

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('1234',),
        ('abc123',),
        ('987654321',)
    ]
    expected_outputs = [
        'True',
        'False',
        'True'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_digits_only(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)