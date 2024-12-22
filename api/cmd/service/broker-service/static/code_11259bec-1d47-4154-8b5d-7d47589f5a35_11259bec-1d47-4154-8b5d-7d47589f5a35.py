# User function logic goes here
def subarrays_divisible_by_k(nums, k):
    count = 0
    prefix_sum = {0: 1}
    current_sum = 0
    for num in nums:
        current_sum += num
        mod = current_sum % k
        if mod < 0:
            mod += k
        if mod in prefix_sum:
            count += prefix_sum[mod]
        prefix_sum[mod] = prefix_sum.get(mod, 0) + 1
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([23, 2, 4, 6, 7], 6)
    ]
    expected_outputs = [
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = subarrays_divisible_by_k(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)