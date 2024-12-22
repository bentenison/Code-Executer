# User function logic goes here
def main(d):
    return dict(sorted(d.items(), key=lambda item: item[1]))

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 5, 'b': 10, 'c': 3},),
        ({'x': 20, 'y': 10, 'z': 30},)
    ]
    expected_outputs = [
        {'c': 3, 'a': 5, 'b': 10},
        {'y': 10, 'x': 20, 'z': 30}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)