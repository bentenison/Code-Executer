# User function logic goes here
def main(s1, s2):
    it = iter(s2)
    return all(c in it for c in s1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abc', 'aabbcc'),
        ('axc', 'aabbcc'),
        ('ace', 'abcde'),
        ('abc', 'abdc')
    ]
    expected_outputs = [
        True,
        False,
        True,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)