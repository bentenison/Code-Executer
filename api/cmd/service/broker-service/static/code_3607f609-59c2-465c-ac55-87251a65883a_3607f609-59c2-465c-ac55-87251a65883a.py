# User function logic goes here
def concatenate_strings(s1, s2):
    return s1 + s2

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', 'world'),
        ('abc', 'def'),
        ('a', 'b'),
        ('', 'empty')
    ]
    expected_outputs = [
        'helloworld',
        'abcdef',
        'ab',
        'empty'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = concatenate_strings(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)