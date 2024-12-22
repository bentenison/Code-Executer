# User function logic goes here
def minCut(s):
    n = len(s)
    dp = [i for i in range(n)]
    for i in range(n):
        for j in range(i, n):
            if s[i:j + 1] == s[i:j + 1][::-1]:
                dp[j] = min(dp[j], dp[i - 1] + 1 if i > 0 else 0)
    return dp[-1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('aab'),
        ('a')
    ]
    expected_outputs = [
        1,
        0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = minCut(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)