# User function logic goes here
def rotate_right(arr, k):
    k = k % len(arr)
    return arr[-k:] + arr[:-k]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4, 5], 2),
        ([1, 2, 3, 4, 5], 3)
    ]
    expected_outputs = [
        [4, 5, 1, 2, 3],
        [3, 4, 5, 1, 2]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = rotate_right(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)