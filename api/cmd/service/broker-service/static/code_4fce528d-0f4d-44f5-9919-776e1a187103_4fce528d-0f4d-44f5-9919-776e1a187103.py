# User function logic goes here
def find_common_elements(arr1, arr2, arr3):
    return list(set(arr1) & set(arr2) & set(arr3))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 5, 10], [2, 5, 10], [5, 10]),
        ([1, 2, 3], [4, 5, 6], [7, 8, 9]),
    ]
    expected_outputs = [
        [5, 10],
        [],
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_common_elements(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)