# User function logic goes here
def main(s):
    n = len(s) // 2
    return s[:n].upper() + s[n:]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('worldwide',),
        ('python',),
        ('chatgpt',)
    ]
    expected_outputs = [
        'HEllo',
        'WORLdwidE',
        'PYThon',
        'CHAThgt'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)