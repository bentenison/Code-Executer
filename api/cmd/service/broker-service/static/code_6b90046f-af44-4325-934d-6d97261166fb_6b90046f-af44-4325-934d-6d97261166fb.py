# User function logic goes here
def main(s):
    has_letter = any(char.isalpha() for char in s)
    has_number = any(char.isdigit() for char in s)
    return has_letter and has_number

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abc123',),
        ('1234',),
        ('abcd',),
        ('1a2b3c',)
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