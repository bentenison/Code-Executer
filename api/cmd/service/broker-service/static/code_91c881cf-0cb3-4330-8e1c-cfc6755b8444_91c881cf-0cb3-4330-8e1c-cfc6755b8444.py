# User function logic goes here
def main(s1, s2):
    if len(s1) != len(s2):
        return False
    mapping = {}
    for char1, char2 in zip(s1, s2):
        if char1 in mapping and mapping[char1] != char2:
            return False
        mapping[char1] = char2
    return True

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('egg', 'add'),
        ('foo', 'bar'),
        ('paper', 'title'),
        ('ab', 'aa')
    ]
    expected_outputs = [
        True,
        False,
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)