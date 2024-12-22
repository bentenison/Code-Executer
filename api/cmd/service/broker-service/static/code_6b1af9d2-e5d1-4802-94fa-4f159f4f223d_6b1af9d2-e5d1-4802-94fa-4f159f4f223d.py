# User function logic goes here
def main(a, b, c):
    return max(a, b, c)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (3, 7, 5, 7),
        (10, 5, 7, 10),
        (1, 1, 1, 1),
        (0, -5, 2, 2)
    ]
    for a, b, c, expected in test_cases:
        result = main(a, b, c)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: ({a}, {b}, {c}). Expected: {expected}, Got: {result}')
    if all_passed:
        print('All test cases passed!')
    else:
        print('Some test cases failed.')