# User function logic goes here
def main(arr):
    unique_elements = list(set(arr))
    unique_elements.sort()
    return unique_elements[-2] if len(unique_elements) > 1 else None

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3],),
        ([5, 6, 7, 8, 9],),
        ([10, 10, 10],),
        ([3, 1],)
    ]
    expected_outputs = [
        2,
        8,
        None,
        1
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)