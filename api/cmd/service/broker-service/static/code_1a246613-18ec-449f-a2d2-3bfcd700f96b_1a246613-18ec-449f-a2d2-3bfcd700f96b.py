# User function logic goes here
def is_multiple_of_3_or_5(num):
    if num % 3 == 0 or num % 5 == 0:
        return 'Multiple of 3 or 5'
    return 'Not a Multiple'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (9,),
        (7,),
        (15,)
    ]
    expected_outputs = [
        'Multiple of 3 or 5',
        'Not a Multiple',
        'Multiple of 3 or 5'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_multiple_of_3_or_5(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)