# User function logic goes here
def maxSumOfThreeSubarrays(arr, k):
    n = len(arr)
    left, right = [0] * n, [0] * n
    max_left_sum, max_right_sum = [0] * n, [0] * n
    window_sum, max_sum = 0, 0
    for i in range(k):
        window_sum += arr[i]
    max_left_sum[k-1] = window_sum
    for i in range(k, n):
        window_sum = window_sum + arr[i] - arr[i - k]
        max_left_sum[i] = max(max_left_sum[i-1], window_sum)
    for i in range(n - k, n):
        window_sum += arr[i] - arr[i - k]
    return max_sum

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,2,1,2,6,7,5,1], 2),
        ([1,2,1,2,6,7,5,1,1], 3)
    ]
    expected_outputs = [
        23,
        15
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maxSumOfThreeSubarrays(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)