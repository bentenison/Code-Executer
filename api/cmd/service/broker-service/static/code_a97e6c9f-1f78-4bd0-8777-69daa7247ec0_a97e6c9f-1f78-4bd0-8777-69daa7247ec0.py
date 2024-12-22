# User function logic goes here
def reverse_string(s):
    reversed_str = ''
    for char in s:
        reversed_str = char + reversed_str
    return reversed_str

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello'),
        ('world')
    ]
    expected_outputs = [
        'olleh',
        'dlrow'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = reverse_string(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)