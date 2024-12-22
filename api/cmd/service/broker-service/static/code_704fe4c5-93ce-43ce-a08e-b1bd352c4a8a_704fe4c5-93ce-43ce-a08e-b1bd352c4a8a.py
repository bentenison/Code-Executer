# User function logic goes here
import re

def check_date_format(date_str):
    pattern = '^\d{2}/\d{2}/\d{4}$'
    if re.match(pattern, date_str):
        return 'Valid Date'
    return 'Invalid Date'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('12/25/2024',),
        ('02/30/2024',),
        ('11/15/2023',)
    ]
    expected_outputs = [
        'Valid Date',
        'Invalid Date',
        'Valid Date'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_date_format(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)