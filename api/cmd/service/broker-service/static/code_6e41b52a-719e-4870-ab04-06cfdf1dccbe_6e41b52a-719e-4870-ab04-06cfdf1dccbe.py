# User function logic goes here
def longest_increasing_subsequence(arr):
    n = len(arr)
    dp = [1] * n
    for i in range(1, n):
        for j in range(i):
            if arr[i] > arr[j] and dp[i] < dp[j] + 1:
                dp[i] = dp[j] + 1
    return max(dp)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([10, 22, 9, 33, 21, 50, 41, 60, 80],),
        ([3, 10, 2, 1, 20],),
    ]
    expected_outputs = [
        6,
        3,
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = longest_increasing_subsequence(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)