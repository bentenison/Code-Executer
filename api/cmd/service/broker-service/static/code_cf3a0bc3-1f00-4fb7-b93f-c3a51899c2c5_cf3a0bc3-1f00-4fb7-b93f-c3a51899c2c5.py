# User function logic goes here
def main(s):
    stack = []
    for char in s:
        if char == '(': 
            stack.append(char)
        elif char == ')': 
            if not stack: 
                return False
            stack.pop()
    return len(stack) == 0

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('()',),
        ('(()',),
        ('(())',),
        (')(',)
    ]
    expected_outputs = [
        True,
        False,
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)