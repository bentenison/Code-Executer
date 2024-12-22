# User function logic goes here
def sort_colors(nums):
    low, mid, high = 0, 0, len(nums) - 1
    while mid <= high:
        if nums[mid] == 0:
            nums[low], nums[mid] = nums[mid], nums[low]
            low += 1
            mid += 1
        elif nums[mid] == 1:
            mid += 1
        else:
            nums[mid], nums[high] = nums[high], nums[mid]
            high -= 1
    return nums

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([2,0,2,1,1,0])
    ]
    expected_outputs = [
        [0, 0, 1, 1, 2, 2]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = sort_colors(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)