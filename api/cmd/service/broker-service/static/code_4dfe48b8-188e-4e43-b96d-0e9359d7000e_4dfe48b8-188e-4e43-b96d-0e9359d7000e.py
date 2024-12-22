# User function logic goes here
def main(arr1, arr2):
    i, j = 0, 0
    merged_array = []
    while i < len(arr1) and j < len(arr2):
        if arr1[i] < arr2[j]:
            merged_array.append(arr1[i])
            i += 1
        else:
            merged_array.append(arr2[j])
            j += 1
    while i < len(arr1):
        merged_array.append(arr1[i])
        i += 1
    while j < len(arr2):
        merged_array.append(arr2[j])
        j += 1
    return merged_array

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 3, 5], [2, 4, 6]),
        ([], [1, 2, 3]),
        ([1, 2, 3], []),
        ([1, 2], [3, 4])
    ]
    expected_outputs = [
        [1, 2, 3, 4, 5, 6],
        [1, 2, 3],
        [1, 2, 3],
        [1, 2, 3, 4]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)