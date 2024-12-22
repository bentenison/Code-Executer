# User function logic goes here
def check_character(char):
    if char.isdigit():
        return 'Digit'
    elif char.isalpha():
        return 'Letter'
    else:
        return 'Symbol'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('a',),
        ('1',),
        ('@',)
    ]
    expected_outputs = [
        'Letter',
        'Digit',
        'Symbol'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_character(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)