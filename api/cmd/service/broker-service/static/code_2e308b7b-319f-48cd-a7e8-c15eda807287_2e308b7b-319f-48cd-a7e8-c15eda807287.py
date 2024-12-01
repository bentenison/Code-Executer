# User function logic goes here
def sum_of_digits(n):
    if n == 0:
        return 0
    return n % 10 + sum_of_digits(n // 10)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (123,),
        (987,),
        (34567,),
        (5,)
    ]
    expected_outputs = [
        6,
        24,
        25,
        5
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = sum_of_digits(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)