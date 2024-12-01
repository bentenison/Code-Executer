# User function logic goes here
def main(num1, num2):
    # Your code here
    print("data\n",num1)
    if num1 > num2:
        return 'Greater'
    elif num1 < num2:
        return 'Less'
    return 'Equal'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5, 3),
        (2, 2),
        (1, 4),
        (0, 0)
    ]
    expected_outputs = [
        'Greater',
        'Equal',
        'Less',
        'Equal'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)