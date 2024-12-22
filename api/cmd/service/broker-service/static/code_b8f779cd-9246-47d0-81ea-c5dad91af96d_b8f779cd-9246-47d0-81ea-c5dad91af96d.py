# User function logic goes here
def longest_consecutive_subsequence(arr):
    nums = set(arr)
    longest = 0
    for num in arr:
        if num - 1 not in nums:
            current_num = num
            current_streak = 1
            while current_num + 1 in nums:
                current_num += 1
                current_streak += 1
            longest = max(longest, current_streak)
    return longest

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([100, 4, 200, 1, 3, 2]),
        ([1, 9, 3, 10, 4, 20, 2])
    ]
    expected_outputs = [
        4,
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = longest_consecutive_subsequence(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)