# User function logic goes here
def minSubArrayLen(target, nums):
    left, total, result = 0, 0, float('inf')
    for right in range(len(nums)):
        total += nums[right]
        while total >= target:
            result = min(result, right - left + 1)
            total -= nums[left]
            left += 1
    return result if result != float('inf') else 0

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (7, [2,3,1,2,4,3])
    ]
    expected_outputs = [
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = minSubArrayLen(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)