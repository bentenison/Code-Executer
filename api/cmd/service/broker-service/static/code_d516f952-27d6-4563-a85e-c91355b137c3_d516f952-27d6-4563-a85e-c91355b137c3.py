# User function logic goes here
def main(c):
    return c in [' ', '\t', '\n']

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (' ',),
        ('a',),
        ('\n',),
        ('\t',)
    ]
    expected_outputs = [
        True,
        False,
        True,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)