# User function logic goes here
def maxLengthSubarray(nums, k):
    left, total, max_length = 0, 0, 0
    for right in range(len(nums)):
        total += nums[right]
        while total > k:
            total -= nums[left]
            left += 1
        max_length = max(max_length, right - left + 1)
    return max_length

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,2,3,4,5], 11)
    ]
    expected_outputs = [
        3
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maxLengthSubarray(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)