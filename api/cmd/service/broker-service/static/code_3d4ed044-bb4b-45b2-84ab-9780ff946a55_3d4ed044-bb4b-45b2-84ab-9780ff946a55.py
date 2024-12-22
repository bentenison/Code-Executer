# User function logic goes here
def maxSumSubarray(nums, k):
    window_sum = sum(nums[:k])
    max_sum = window_sum
    for i in range(k, len(nums)):
        window_sum += nums[i] - nums[i - k]
        max_sum = max(max_sum, window_sum)
    return max_sum

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([2,1,5,1,3,2], 3)
    ]
    expected_outputs = [
        9
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maxSumSubarray(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)