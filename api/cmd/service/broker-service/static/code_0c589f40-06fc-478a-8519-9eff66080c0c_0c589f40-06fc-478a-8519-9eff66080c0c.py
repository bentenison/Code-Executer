# User function logic goes here
def sum_of_digits(n):
    return sum(int(digit) for digit in str(n))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (1234,),
        (9876,),
        (456,),
        (0,)
    ]
    expected_outputs = [
        10,
        30,
        15,
        0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = sum_of_digits(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)