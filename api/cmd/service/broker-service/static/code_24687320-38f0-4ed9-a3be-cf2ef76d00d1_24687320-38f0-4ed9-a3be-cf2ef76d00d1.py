# User function logic goes here
def check_substring(main_str, sub_str):
    if sub_str in main_str:
        return 'Found'
    else:
        return 'Not Found'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world', 'world'),
        ('python is fun', 'java'),
        ('check substring', 'check')
    ]
    expected_outputs = [
        'Found',
        'Not Found',
        'Found'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_substring(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)