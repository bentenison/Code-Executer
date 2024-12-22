# User function logic goes here
def check_multiple_of_7_and_11(num):
    if num % 7 == 0 and num % 11 == 0:
        return 'Multiple of 7 and 11'
    else:
        return 'Not a multiple of 7 and 11'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (77,),
        (14,),
        (22,)
    ]
    expected_outputs = [
        'Multiple of 7 and 11',
        'Not a multiple of 7 and 11',
        'Not a multiple of 7 and 11'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_multiple_of_7_and_11(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)