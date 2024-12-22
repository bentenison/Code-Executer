# User function logic goes here
from collections import deque

def maxSlidingWindow(nums, k):
    result = []
    dq = deque()
    for i in range(len(nums)):
        while dq and nums[dq[-1]] <= nums[i]:
            dq.pop()
        dq.append(i)
        if dq[0] == i - k:
            dq.popleft()
        if i >= k - 1:
            result.append(nums[dq[0]])
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,3,-1,-3,5,3,6,7], 3)
    ]
    expected_outputs = [
        [3,3,5,5,6,7]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maxSlidingWindow(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)