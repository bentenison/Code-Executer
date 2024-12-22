# User function logic goes here
def main(strs):
    if not strs:
        return ''
    prefix = strs[0]
    for string in strs[1:]:
        while not string.startswith(prefix):
            prefix = prefix[:-1]
            if not prefix:
                return ''
    return prefix

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ['flower', 'flow', 'flight'],
        ['dog', 'racecar', 'car'],
        ['apple', 'appreciate', 'application']
    ]
    expected_outputs = [
        'fl',
        '',
        'app'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)