# User function logic goes here
def main(lst):
    return lst[0] == lst[-1] if len(lst) > 0 else False

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 1],),
        ([2, 3, 4, 5],)
    ]
    expected_outputs = [
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)