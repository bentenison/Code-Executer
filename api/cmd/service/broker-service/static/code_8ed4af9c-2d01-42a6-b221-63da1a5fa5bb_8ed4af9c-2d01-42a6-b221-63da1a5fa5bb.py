# User function logic goes here
def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([10, 7, 8, 9, 1, 5],),
        ([3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5],),
    ]
    expected_outputs = [
        [1, 5, 7, 8, 9, 10],
        [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9],
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = quicksort(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)