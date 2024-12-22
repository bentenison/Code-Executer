# User function logic goes here
def main(d1, d2):
    for key, value in d2.items():
        d1[key] = d1.get(key, 0) + value
    return d1

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 2}, {'b': 3, 'c': 4}),
        ({'x': 5, 'y': 10}, {'x': 1, 'y': 2})
    ]
    expected_outputs = [
        {'a': 1, 'b': 5, 'c': 4},
        {'x': 6, 'y': 12}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)