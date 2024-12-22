# User function logic goes here
def check_range(num):
    if 1 <= num <= 100:
        return 'In range'
    else:
        return 'Out of range'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (50,),
        (101,),
        (75,)
    ]
    expected_outputs = [
        'In range',
        'Out of range',
        'In range'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_range(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)