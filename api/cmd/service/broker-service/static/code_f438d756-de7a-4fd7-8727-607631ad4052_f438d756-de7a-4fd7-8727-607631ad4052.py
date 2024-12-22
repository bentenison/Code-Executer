# User function logic goes here
def is_narcissistic(num):
    digits = [int(d) for d in str(num)]
    return 'Narcissistic Number' if sum(d ** len(digits) for d in digits) == num else 'Not a Narcissistic Number'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (153,),
        (370,),
        (1234,)
    ]
    expected_outputs = [
        'Narcissistic Number',
        'Narcissistic Number',
        'Not a Narcissistic Number'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_narcissistic(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)