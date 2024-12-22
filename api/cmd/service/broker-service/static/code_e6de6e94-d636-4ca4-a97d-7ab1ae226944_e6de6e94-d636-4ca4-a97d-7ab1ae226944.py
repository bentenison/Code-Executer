# User function logic goes here
def main(s, k):
    def is_palindrome(s):
        return s == s[::-1]
    count = 0
    for i in range(len(s)):
        for j in range(i+k, len(s)+1):
            if is_palindrome(s[i:j]):
                count += 1
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('racecar', 3),
        ('abacdfgdcaba', 4),
        ('madam', 2)
    ]
    expected_outputs = [
        2,
        4,
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)