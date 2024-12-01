# User function logic goes here
def are_anagrams(str1, str2):
    return sorted(str1) == sorted(str2)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('listen', 'silent'),
        ('hello', 'world'),
        ('evil', 'vile'),
        ('python', 'typhon')
    ]
    expected_outputs = [
        True,
        False,
        True,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = are_anagrams(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)