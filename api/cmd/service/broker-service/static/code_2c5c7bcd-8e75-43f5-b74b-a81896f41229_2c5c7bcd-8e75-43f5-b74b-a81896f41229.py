# User function logic goes here
def findMin(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] > nums[right]:
            left = mid + 1
        else:
            right = mid
    return nums[left]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([4,5,6,7,0,1,2])
    ]
    expected_outputs = [
        0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = findMin(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)