# User function logic goes here
def gcd(a, b):
    while b:
        a, b = b, a % b
    return a

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (56, 98),
        (48, 180)
    ]
    expected_outputs = [
        14,
        12
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = gcd(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)