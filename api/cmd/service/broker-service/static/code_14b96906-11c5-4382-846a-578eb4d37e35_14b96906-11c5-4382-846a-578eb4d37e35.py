# User function logic goes here
def max_area(height):
    left, right = 0, len(height) - 1
    max_area = 0
    while left < right:
        max_area = max(max_area, min(height[left], height[right]) * (right - left))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return max_area

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1,8,6,2,5,4,8,3,7])
    ]
    expected_outputs = [
        49
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = max_area(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)