# User function logic goes here
def longest_palindromic_substring(s):
    n = len(s)
    if n == 0:
        return ""
    start = 0
    max_length = 1

    def expand_around_center(left, right):
        while left >= 0 and right < n and s[left] == s[right]:
            left -= 1
            right += 1
        return right - left - 1

    for i in range(n):
        len1 = expand_around_center(i, i)
        len2 = expand_around_center(i, i + 1)
        max_len = max(len1, len2)
        if max_len > max_length:
            max_length = max_len
            start = i - (max_len - 1) // 2

    return s[start:start + max_length}

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('babad',),
        ('cbbd',),
        ('a',),
        ('ac',)
    ]
    expected_outputs = [
        'bab',
        'bb',
        'a',
        'a'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = longest_palindromic_substring(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)