# User function logic goes here
def remove_duplicates(arr):
    arr.sort()
    i = 0
    for j in range(1, len(arr)):
        if arr[i] != arr[j]:
            i += 1
            arr[i] = arr[j]
    return arr[:i+1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 2, 3, 4, 4, 5]),
        ([10, 10, 20, 20, 30])
    ]
    expected_outputs = [
        [1, 2, 3, 4, 5],
        [10, 20, 30]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = remove_duplicates(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)