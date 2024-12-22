# User function logic goes here
def main(s):
    char_index = {}
    left = 0
    max_length = 0
    for right in range(len(s)):
        if s[right] in char_index:
            left = max(left, char_index[s[right]] + 1)
        char_index[s[right]] = right
        max_length = max(max_length, right - left + 1)
    return max_length

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abcabcbb',),
        ('bbbbb',),
        ('pwwkew',),
        ('',)
    ]
    expected_outputs = [
        3,
        1,
        3,
        0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)