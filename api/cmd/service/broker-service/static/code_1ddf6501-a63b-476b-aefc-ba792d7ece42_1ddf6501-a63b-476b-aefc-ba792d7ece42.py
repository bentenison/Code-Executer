# User function logic goes here
def check_prime(num):
    if num > 1:
        for i in range(2, num):
            if num % i == 0:
                return 'Not Prime'
        return 'Prime'
    else:
        return 'Not Prime'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5,),
        (10,),
        (13,)
    ]
    expected_outputs = [
        'Prime',
        'Not Prime',
        'Prime'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_prime(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)