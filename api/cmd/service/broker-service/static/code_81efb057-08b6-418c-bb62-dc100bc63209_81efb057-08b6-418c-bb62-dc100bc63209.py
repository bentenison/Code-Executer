# User function logic goes here
def max_consecutive_substring(s):
    max_substring = ''
    max_count = 0
    current_substring = ''
    current_count = 1
    for i in range(1, len(s)):
        if s[i] == s[i-1]:
            current_count += 1
            current_substring += s[i]
        else:
            if current_count > max_count:
                max_substring = current_substring
                max_count = current_count
            current_substring = s[i]
            current_count = 1
    return max_substring, max_count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('aabbbaac',),
        ('abcabcabc',),
        ('zzzzzz',)
    ]
    expected_outputs = [
        ('aa', 2),
        ('abc', 3),
        ('z', 6)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = max_consecutive_substring(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)