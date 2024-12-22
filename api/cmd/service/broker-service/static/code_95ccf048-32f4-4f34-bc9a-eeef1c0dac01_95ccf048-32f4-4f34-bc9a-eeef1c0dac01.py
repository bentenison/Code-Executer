# User function logic goes here
def main(lst):
    return [x ** 2 for x in lst]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3],),
        ([4, 5, 6],)
    ]
    expected_outputs = [
        [1, 4, 9],
        [16, 25, 36]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)