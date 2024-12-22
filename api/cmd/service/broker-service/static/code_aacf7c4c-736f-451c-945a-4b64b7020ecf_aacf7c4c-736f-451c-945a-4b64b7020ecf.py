# User function logic goes here
def subarray_sum(nums, k):
    count = 0
    prefix_sum = {0: 1}
    current_sum = 0
    for num in nums:
        current_sum += num
        if current_sum - k in prefix_sum:
            count += prefix_sum[current_sum - k]
        prefix_sum[current_sum] = prefix_sum.get(current_sum, 0) + 1
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3], 3),
        ([1, 1, 1], 2)
    ]
    expected_outputs = [
        2,
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = subarray_sum(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)