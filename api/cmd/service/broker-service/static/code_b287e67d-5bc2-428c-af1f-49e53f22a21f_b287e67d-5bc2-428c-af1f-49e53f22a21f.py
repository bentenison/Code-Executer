# User function logic goes here
def closest_to_zero(arr):
    arr.sort()
    left, right = 0, len(arr) - 1
    closest_sum = float('inf')
    result = []
    while left < right:
        current_sum = arr[left] + arr[right]
        if abs(current_sum) < abs(closest_sum):
            closest_sum = current_sum
            result = [arr[left], arr[right]]
        if current_sum < 0:
            left += 1
        else:
            right -= 1
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([-1, 2, 3, -4, 5]),
        ([1, 60, -10, 70, -80, 85])
    ]
    expected_outputs = [
        [-1, 2],
        [-10, 10]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = closest_to_zero(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)