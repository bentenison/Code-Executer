# User function logic goes here
def numSubarrayProductLessThanK(nums, k):
    left = 0
    product = 1
    count = 0
    for right in range(len(nums)):
        product *= nums[right]
        while product >= k and left <= right:
            product //= nums[left]
            left += 1
        count += right - left + 1
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([10, 5, 2, 6], 100),
        ([1, 2, 3], 6)
    ]
    expected_outputs = [
        8,
        5
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = numSubarrayProductLessThanK(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)