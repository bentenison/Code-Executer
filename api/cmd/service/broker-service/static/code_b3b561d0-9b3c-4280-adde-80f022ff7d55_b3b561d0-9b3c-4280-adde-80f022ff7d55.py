# User function logic goes here
def check_start_end(s, char):
    return 'True' if s.startswith(char) or s.endswith(char) else 'False'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', 'h'),
        ('world', 'w'),
        ('python', 't')
    ]
    expected_outputs = [
        'True',
        'True',
        'False'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_start_end(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)