# User function logic goes here
def is_divisible(num1, num2):
    return 'Divisible' if num1 % num2 == 0 else 'Not Divisible'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (10, 2),
        (7, 3),
        (15, 5)
    ]
    expected_outputs = [
        'Divisible',
        'Not Divisible',
        'Divisible'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_divisible(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)