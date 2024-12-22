# User function logic goes here
def max_product(nums):
    nums.sort()
    return max(nums[0] * nums[1], nums[-1] * nums[-2])

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4]),
        ([-10, -10, 5, 2])
    ]
    expected_outputs = [
        12,
        100
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = max_product(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)