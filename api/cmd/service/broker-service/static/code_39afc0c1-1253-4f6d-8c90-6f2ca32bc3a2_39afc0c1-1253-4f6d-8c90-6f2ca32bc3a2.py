# User function logic goes here
def lengthOfLIS(nums):
    if not nums:
        return 0
    dp = [1] * len(nums)
    for i in range(1, len(nums)):
        for j in range(i):
            if nums[i] > nums[j]:
                dp[i] = max(dp[i], dp[j] + 1)
    return max(dp)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([10, 9, 2, 5, 3, 7, 101, 18])
    ]
    expected_outputs = [
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = lengthOfLIS(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)