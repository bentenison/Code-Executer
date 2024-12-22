# User function logic goes here
def lcm(a, b):
    return abs(a * b) // gcd(a, b)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (12, 15),
        (7, 5)
    ]
    expected_outputs = [
        60,
        35
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = lcm(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)