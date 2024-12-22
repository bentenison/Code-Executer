# User function logic goes here
def intersection(arr1, arr2):
    return list(set(arr1) & set(arr2))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 2, 1], [2, 2]),
        ([4, 9, 5], [9, 4, 9, 8, 4])
    ]
    expected_outputs = [
        [2],
        [9, 4]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = intersection(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)