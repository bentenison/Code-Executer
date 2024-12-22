# User function logic goes here
import datetime

def get_day_of_week(date):
    date_obj = datetime.datetime.strptime(date, '%m/%d/%Y')
    return date_obj.strftime('%A')

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('12/25/2024',),
        ('07/04/2024',),
        ('01/01/2024',)
    ]
    expected_outputs = [
        'Wednesday',
        'Thursday',
        'Monday'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = get_day_of_week(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)