# User function logic goes here
def main(d):
    value_to_keys = {}
    for k, v in d.items():
        value_to_keys.setdefault(v, []).append(k)
    return [keys for keys in value_to_keys.values() if len(keys) > 1]

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ({'a': 1, 'b': 1, 'c': 2},),
        ({'x': 10, 'y': 10, 'z': 20},)
    ]
    expected_outputs = [
        ['a', 'b'],
        ['x', 'y']
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)