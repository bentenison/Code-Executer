# User function logic goes here
def max_subarray_sum(arr):
    def cross_sum(arr, left, right, mid):
        left_sum = float('-inf')
        right_sum = float('-inf')
        temp_sum = 0
        for i in range(mid, left-1, -1):
            temp_sum += arr[i]
            left_sum = max(left_sum, temp_sum)
        temp_sum = 0
        for i in range(mid + 1, right + 1):
            temp_sum += arr[i]
            right_sum = max(right_sum, temp_sum)
        return left_sum + right_sum

    def divide(arr, left, right):
        if left == right:
            return arr[left]
        mid = (left + right) // 2
        left_sum = divide(arr, left, mid)
        right_sum = divide(arr, mid + 1, right)
        cross_sum_value = cross_sum(arr, left, right, mid)
        return max(left_sum, right_sum, cross_sum_value)

    return divide(arr, 0, len(arr) - 1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([-2, 1, -3, 4, -1, 2, 1, -5, 4],),
        ([-1, -2, -3, -4],),
    ]
    expected_outputs = [
        6,
        -1,
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = max_subarray_sum(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)