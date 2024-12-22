# User function logic goes here
def find_disappeared_numbers(nums):
    missing_numbers = []
    for i in range(len(nums)):
        index = abs(nums[i]) - 1
        if nums[index] > 0:
            nums[index] = -nums[index]
    for i in range(len(nums)):
        if nums[i] > 0:
            missing_numbers.append(i + 1)
    return missing_numbers

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([4,3,2,7,8,2,3,1])
    ]
    expected_outputs = [
        [5,6]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_disappeared_numbers(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)