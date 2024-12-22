# User function logic goes here
def findPeakElement(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] > nums[mid + 1]:
            right = mid
        else:
            left = mid + 1
    return left

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,2,3,1])
    ]
    expected_outputs = [
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = findPeakElement(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)