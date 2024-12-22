# User function logic goes here
def findDuplicate(nums):
    slow, fast = nums[0], nums[0]
    while True:
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast:
            break
    slow = nums[0]
    while slow != fast:
        slow = nums[slow]
        fast = nums[fast]
    return slow

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,3,4,2,2])
    ]
    expected_outputs = [
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = findDuplicate(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)