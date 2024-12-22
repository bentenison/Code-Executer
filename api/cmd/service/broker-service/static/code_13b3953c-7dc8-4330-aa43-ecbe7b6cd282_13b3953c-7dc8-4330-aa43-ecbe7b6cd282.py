# User function logic goes here
def three_sum(nums, target):
    nums.sort()
    result = []
    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        left, right = i + 1, len(nums) - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total == target:
                result.append([nums[i], nums[left], nums[right]])
                while left < right and nums[left] == nums[left + 1]:
                    left += 1
                while left < right and nums[right] == nums[right - 1]:
                    right -= 1
                left += 1
                right -= 1
            elif total < target:
                left += 1
            else:
                right -= 1
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([-1, 0, 1, 2, -1, -4], 0)
    ]
    expected_outputs = [
        [[-1, -1, 2], [-1, 0, 1]]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = three_sum(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)