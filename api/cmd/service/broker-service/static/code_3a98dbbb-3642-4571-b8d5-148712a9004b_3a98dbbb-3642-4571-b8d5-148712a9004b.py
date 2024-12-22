# User function logic goes here
def merge_and_add(d1, d2):
    merged = d1.copy()
    for k, v in d2.items():
        if k in merged:
            merged[k] += v
        else:
            merged[k] = v
    return merged

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ({'a': 2, 'b': 3}, {'a': 4, 'c': 5}),
        ({'x': 1}, {'x': 2, 'y': 3}),
    ]
    expected_outputs = [
        {'a': 6, 'b': 3, 'c': 5},
        {'x': 3, 'y': 3},
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = merge_and_add(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)