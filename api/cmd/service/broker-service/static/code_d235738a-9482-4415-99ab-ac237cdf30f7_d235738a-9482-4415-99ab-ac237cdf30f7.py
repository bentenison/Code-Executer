# User function logic goes here
def majority_element(arr):
    count = 0
    candidate = None
    for num in arr:
        if count == 0:
            candidate = num
        count += (1 if num == candidate else -1)
    count = arr.count(candidate)
    return candidate if count > len(arr) // 2 else None

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([3, 3, 4, 2, 4, 4, 2, 4, 4]),
        ([1, 2, 3, 4])
    ]
    expected_outputs = [
        4,
        None
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = majority_element(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)