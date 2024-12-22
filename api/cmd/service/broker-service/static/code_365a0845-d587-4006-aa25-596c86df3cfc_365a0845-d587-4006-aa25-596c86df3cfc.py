# User function logic goes here
def merge_sorted_arrays(arr1, arr2):
    merged = []
    i = j = 0
    while i < len(arr1) and j < len(arr2):
        if arr1[i] < arr2[j]:
            merged.append(arr1[i])
            i += 1
        else:
            merged.append(arr2[j])
            j += 1
    merged.extend(arr1[i:])
    merged.extend(arr2[j:])
    return merged

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 3, 5], [2, 4, 6]),
        ([1, 4, 7], [2, 5, 8])
    ]
    expected_outputs = [
        [1, 2, 3, 4, 5, 6],
        [1, 2, 4, 5, 7, 8]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = merge_sorted_arrays(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)