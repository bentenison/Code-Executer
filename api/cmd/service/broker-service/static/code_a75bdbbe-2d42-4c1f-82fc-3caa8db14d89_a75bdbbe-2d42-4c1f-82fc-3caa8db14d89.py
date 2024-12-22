# User function logic goes here
def main(s):
    length = len(s)
    if length % 2 == 1:
        return s[length // 2]
    else:
        return s[length // 2 - 1: length // 2 + 1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('abcd',),
        ('python',),
        ('abcde',)
    ]
    expected_outputs = [
        'l',
        'bc',
        't',
        'c'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)