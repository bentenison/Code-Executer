# User function logic goes here
def check_palindrome(num):
    original = num
    reversed_num = 0
    while num > 0:
        reversed_num = reversed_num * 10 + num % 10
        num //= 10
    if original == reversed_num:
        return 'Palindrome'
    else:
        return 'Not a palindrome'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (121,),
        (12321,),
        (10,)
    ]
    expected_outputs = [
        'Palindrome',
        'Palindrome',
        'Not a palindrome'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = check_palindrome(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)