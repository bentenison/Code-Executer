# User function logic goes here
def main(arr):
    return max(arr)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4, 5],),
        ([-5, -3, -10, 0, 7],),
        ([12, 19, 4, 2, 8],),
        ([9],)
    ]
    expected_outputs = [
        5,
        7,
        19,
        9
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)