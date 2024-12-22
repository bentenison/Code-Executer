# User function logic goes here
def main(arr):
    max_sum = current_sum = arr[0]
    for num in arr[1:]:
        current_sum = max(num, current_sum + num)
        max_sum = max(max_sum, current_sum)
    return max_sum

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, -2, 5],),
        ([-1, -2, -3, -4],),
        ([4, -1, 2, 1],),
        ([2, 3, 4, 5, -10, 10],)
    ]
    expected_outputs = [
        9,
        -1,
        6,
        10
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)