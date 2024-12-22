# User function logic goes here
def main(d1, d2):
    return {k: v for k, v in d1.items() if k in d2 and d1[k] == d2[k]}

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2}, {'b': 2, 'c': 3}),
        ({'x': 10, 'y': 20}, {'y': 20, 'z': 30})
    ]
    expected_outputs = [
        {'b': 2},
        {'y': 20}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)