# User function logic goes here
def partition_array(arr, pivot):
    smaller = [x for x in arr if x < pivot]
    greater = [x for x in arr if x > pivot]
    return smaller + [pivot] + greater

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([3, 5, 2, 8, 1, 4], 4),
        ([7, 1, 3, 9, 4, 6], 5)
    ]
    expected_outputs = [
        [3, 2, 1, 4, 8, 5],
        [1, 3, 4, 5, 9, 6]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = partition_array(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)