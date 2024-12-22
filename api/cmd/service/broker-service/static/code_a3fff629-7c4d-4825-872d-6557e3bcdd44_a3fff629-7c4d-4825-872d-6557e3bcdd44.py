# User function logic goes here
def longestKSubstr(s, k):
    char_count = {}
    left = 0
    max_len = -1
    for right in range(len(s)):
        char_count[s[right]] = char_count.get(s[right], 0) + 1
        while len(char_count) > k:
            char_count[s[left]] -= 1
            if char_count[s[left]] == 0:
                del char_count[s[left]]
            left += 1
        max_len = max(max_len, right - left + 1)
    return max_len

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('araaci', 2)
    ]
    expected_outputs = [
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = longestKSubstr(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)