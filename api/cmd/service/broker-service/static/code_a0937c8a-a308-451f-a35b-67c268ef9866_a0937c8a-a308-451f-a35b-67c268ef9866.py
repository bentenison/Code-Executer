# User function logic goes here
import re

def check_time_format(time_str):
    pattern = '^([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$'
    if re.match(pattern, time_str):
        return 'Valid Time'
    return 'Invalid Time'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('12:30:45',),
        ('25:61:00',),
        ('11:10:09',)
    ]
    expected_outputs = [
        'Valid Time',
        'Invalid Time',
        'Valid Time'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_time_format(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)