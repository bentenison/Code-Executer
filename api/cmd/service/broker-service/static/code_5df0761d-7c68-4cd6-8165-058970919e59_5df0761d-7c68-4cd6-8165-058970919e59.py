# User function logic goes here
def main(arr):
    n = len(arr) + 1
    expected_sum = n * (n + 1) // 2
    actual_sum = sum(arr)
    return expected_sum - actual_sum

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 5],),
        ([1, 3, 4, 5, 6],),
        ([1, 2, 4, 5, 6],),
        ([1, 3],)
    ]
    expected_outputs = [
        4,
        2,
        3,
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)