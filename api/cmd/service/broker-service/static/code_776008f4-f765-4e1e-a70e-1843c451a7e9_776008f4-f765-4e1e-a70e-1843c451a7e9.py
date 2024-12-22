# User function logic goes here
def first_missing_positive(nums):
    nums = set(nums)
    i = 1
    while i in nums:
        i += 1
    return i

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 0]),
        ([3, 4, -1, 1])
    ]
    expected_outputs = [
        3,
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = first_missing_positive(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)