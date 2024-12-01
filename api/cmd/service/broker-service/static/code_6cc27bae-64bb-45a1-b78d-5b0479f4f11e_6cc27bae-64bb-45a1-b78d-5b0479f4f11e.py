# User function logic goes here
def factorial(n):
    if n == 0:
        return 1
    return n * factorial(n - 1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5,),
        (0,),
        (7,),
        (10,)
    ]
    expected_outputs = [
        120,
        1,
        5040,
        3628800
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = factorial(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)