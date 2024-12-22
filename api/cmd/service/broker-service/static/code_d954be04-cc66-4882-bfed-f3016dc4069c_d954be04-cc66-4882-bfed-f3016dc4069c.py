# User function logic goes here
def longest_palindrome(s):
    if not s:
        return ''
    def expand_around_center(left, right):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return s[left + 1:right]
    longest = ''
    for i in range(len(s)):  
        odd = expand_around_center(i, i)
        even = expand_around_center(i, i + 1)
        longest = max(longest, odd, even, key=len)
    return longest

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
        result = longest_palindrome(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)