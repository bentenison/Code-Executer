# User function logic goes here
def main(s):
    from collections import Counter
    count = Counter(s)
    for char in s:
        if count[char] == 1:
            return char
    return '_'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('abcdab',),
        ('aabbcc',),
        ('abc',),
        ('aabb',)
    ]
    expected_outputs = [
        'c',
        '_',
        'a',
        '_'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)