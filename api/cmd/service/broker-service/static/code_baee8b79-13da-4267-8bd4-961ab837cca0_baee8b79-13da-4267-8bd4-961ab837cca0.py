# User function logic goes here
def substring_in_strings_list(strings, substring):
    return any(substring in s for s in strings)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (['hello', 'world'], 'world'),
        (['python', 'rocks'], 'java'),
        (['substring', 'found'], 'sub'),
        (['test', 'case'], 'exam')
    ]
    expected_outputs = [
        True,
        False,
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = substring_in_strings_list(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)