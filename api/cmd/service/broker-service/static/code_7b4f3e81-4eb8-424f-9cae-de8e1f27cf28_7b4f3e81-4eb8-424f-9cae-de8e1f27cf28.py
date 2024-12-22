# User function logic goes here
def extract_digits_from_tuple_list(tuples):
    result = ''.join(''.join([ch for ch in t if ch.isdigit()]) for t in sum(tuples, []))
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        [('a1b2', 'c3d4'), ('e5f6',)],
        [('123', 'abc'), ('def',)],
        [('xyz', '789')]
    ]
    expected_outputs = [
        '123456',
        '123',
        '789'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = extract_digits_from_tuple_list(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)