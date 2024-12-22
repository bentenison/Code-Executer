# User function logic goes here
def main(d):
    return min(d.values()), max(d.values())

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 5, 'b': 10, 'c': 15},),
        ({'x': 20, 'y': 30, 'z': 10},)
    ]
    expected_outputs = [
        (5, 15),
        (10, 30)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print('Test Passed')