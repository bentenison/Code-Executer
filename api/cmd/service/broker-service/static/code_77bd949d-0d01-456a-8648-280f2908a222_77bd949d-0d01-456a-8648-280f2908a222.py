# User function logic goes here
def productExceptSelf(nums):
    n = len(nums)
    left, right, result = [1] * n, [1] * n, []
    for i in range(1, n):
        left[i] = left[i - 1] * nums[i - 1]
    for i in range(n - 2, -1, -1):
        right[i] = right[i + 1] * nums[i + 1]
    for i in range(n):
        result.append(left[i] * right[i])
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,2,3,4])
    ]
    expected_outputs = [
        [24, 12, 8, 6]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = productExceptSelf(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)