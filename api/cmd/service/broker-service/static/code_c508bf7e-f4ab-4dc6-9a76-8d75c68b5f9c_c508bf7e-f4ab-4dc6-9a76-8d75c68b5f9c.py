# User function logic goes here
def max_product(nums):
    if not nums:
        return 0
    max_prod, min_prod, result = nums[0], nums[0], nums[0]
    for i in range(1, len(nums)):
        if nums[i] < 0:
            max_prod, min_prod = min_prod, max_prod
        max_prod = max(nums[i], max_prod * nums[i])
        min_prod = min(nums[i], min_prod * nums[i])
        result = max(result, max_prod)
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([2, 3, -2, 4])
    ]
    expected_outputs = [
        6
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = max_product(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)