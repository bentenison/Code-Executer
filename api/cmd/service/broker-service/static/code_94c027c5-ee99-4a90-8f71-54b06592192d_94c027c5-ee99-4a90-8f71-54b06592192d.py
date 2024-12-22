# User function logic goes here
def largestRectangleArea(heights):
    stack = [-1]
    max_area = 0
    for i, h in enumerate(heights):
        while heights[stack[-1]] >= h:
            h_idx = stack.pop()
            max_area = max(max_area, heights[h_idx] * (i - stack[-1] - 1))
        stack.append(i)
    while heights[stack[-1]] != -1:
        h_idx = stack.pop()
        max_area = max(max_area, heights[h_idx] * (len(heights) - stack[-1] - 1))
    return max_area

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([2,1,5,6,2,3])
    ]
    expected_outputs = [
        10
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = largestRectangleArea(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)