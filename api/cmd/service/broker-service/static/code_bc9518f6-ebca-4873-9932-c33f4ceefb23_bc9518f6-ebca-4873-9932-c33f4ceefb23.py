# User function logic goes here
def rotate_array(arr, n):
    n = n % len(arr)
    arr[:] = arr[-n:] + arr[:-n]
    return arr

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4, 5], 2),
        ([1, 2, 3, 4, 5], 4)
    ]
    expected_outputs = [
        [4, 5, 1, 2, 3],
        [2, 3, 4, 5, 1]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = rotate_array(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)