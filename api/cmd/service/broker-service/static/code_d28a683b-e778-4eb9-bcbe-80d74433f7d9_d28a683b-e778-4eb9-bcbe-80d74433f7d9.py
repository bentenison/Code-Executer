# User function logic goes here
def main(s1, s2):
    i = 0
    for char in s2:
        if i < len(s1) and s1[i] == char:
            i += 1
    return i == len(s1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abc', 'ahbgdc'),
        ('axc', 'ahbgdc'),
        ('', 'ahbgdc')
    ]
    expected_outputs = [
        True,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)