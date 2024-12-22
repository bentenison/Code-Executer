# User function logic goes here
def trap(height):
    if not height:
        return 0
    left, right = 0, len(height) - 1
    left_max, right_max = height[left], height[right]
    water_trapped = 0
    while left < right:
        if height[left] < height[right]:
            left += 1
            left_max = max(left_max, height[left])
            water_trapped += max(0, left_max - height[left])
        else:
            right -= 1
            right_max = max(right_max, height[right])
            water_trapped += max(0, right_max - height[right])
    return water_trapped

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([0,1,0,2,1,0,1,3,2,1,2,1])
    ]
    expected_outputs = [
        6
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = trap(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)