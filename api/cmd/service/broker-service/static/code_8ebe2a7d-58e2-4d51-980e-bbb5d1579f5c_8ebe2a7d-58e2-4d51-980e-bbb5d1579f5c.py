# User function logic goes here
def main(arr):
    n = len(arr)
    expected_sum = n * (n + 1) // 2
    actual_sum = sum(arr)
    return expected_sum - actual_sum

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([3, 0, 1],),
        ([0, 1],),
        ([9,6,4,2,3,5,7,0,1],),
        ([0],)
    ]
    expected_outputs = [2, 2, 8, 1]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)