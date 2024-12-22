# User function logic goes here
import heapq

def find_kth_largest(nums, k):
    return heapq.nlargest(k, nums)[-1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([3, 2, 1, 5, 6, 4], 2),
        ([3, 2, 3, 1, 2, 4, 5, 5, 6], 4)
    ]
    expected_outputs = [
        5,
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_kth_largest(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)