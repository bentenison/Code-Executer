# User function logic goes here
def longest_consecutive(nums):
    num_set = set(nums)
    longest = 0
    for num in num_set:
        if num - 1 not in num_set:
            current_num = num
            current_length = 1
            while current_num + 1 in num_set:
                current_num += 1
                current_length += 1
            longest = max(longest, current_length)
    return longest

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([100, 4, 200, 1, 3, 2]),
        ([0, 1, 2, 3, 4, 5])
    ]
    expected_outputs = [
        4,
        6
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = longest_consecutive(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)