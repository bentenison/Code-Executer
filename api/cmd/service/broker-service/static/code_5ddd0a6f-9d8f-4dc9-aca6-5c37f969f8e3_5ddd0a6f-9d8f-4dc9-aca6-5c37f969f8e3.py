# User function logic goes here
def check_numbers(num1, num2):
    if num1 == num2:
        return 'Equal'
    elif num1 > num2:
        return 'Greater'
    else:
        return 'Smaller'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5, 10),
        (10, 5),
        (7, 7)
    ]
    expected_outputs = [
        'Smaller',
        'Greater',
        'Equal'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_numbers(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)