# User function logic goes here
import datetime

def main(date):
    day = datetime.datetime.strptime(date, '%Y-%m-%d').weekday()
    return day == 5 or day == 6

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('2024-12-14',),
        ('2024-12-15',),
        ('2024-12-16',),
        ('2024-12-17',)
    ]
    expected_outputs = [
        True,
        True,
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)