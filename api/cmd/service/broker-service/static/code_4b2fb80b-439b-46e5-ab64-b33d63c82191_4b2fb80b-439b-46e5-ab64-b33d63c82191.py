# User function logic goes here
def is_valid_triangle(a, b, c):
    if a + b > c and a + c > b and b + c > a:
        return 'Valid Triangle'
    return 'Invalid Triangle'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (3, 4, 5),
        (1, 2, 3),
        (5, 5, 5)
    ]
    expected_outputs = [
        'Valid Triangle',
        'Invalid Triangle',
        'Valid Triangle'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = is_valid_triangle(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)