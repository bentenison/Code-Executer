# User function logic goes here
def check_char_type(c):
    if c.isalpha():
        return 'Alphabet'
    elif c.isdigit():
        return 'Digit'
    return 'Special Character'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('a',),
        ('1',),
        ('@',)
    ]
    expected_outputs = [
        'Alphabet',
        'Digit',
        'Special Character'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_char_type(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)