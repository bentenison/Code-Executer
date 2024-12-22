# User function logic goes here
def union_arrays(arr1, arr2):
    return list(set(arr1) | set(arr2))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3], [3, 4, 5]),
        ([1, 1, 2, 2], [2, 3, 3])
    ]
    expected_outputs = [
        [1, 2, 3, 4, 5],
        [1, 2, 3]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = union_arrays(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)