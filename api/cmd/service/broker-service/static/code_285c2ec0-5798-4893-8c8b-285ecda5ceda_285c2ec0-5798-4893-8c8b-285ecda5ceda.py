# User function logic goes here
def main(arr):
    return list(dict.fromkeys(arr))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 2, 3, 4, 4],),
        ([1, 1, 1, 1, 1],),
        ([5, 6, 7, 5, 7, 8],),
        ([9, 10],)
    ]
    expected_outputs = [
        [1, 2, 3, 4],
        [1],
        [5, 6, 7, 8],
        [9, 10]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)