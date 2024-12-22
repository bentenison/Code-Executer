# User function logic goes here
def main(s):
    freq = {}
    for char in s:
        freq[char] = freq.get(char, 0) + 1
    return freq

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ('hello',),
        ('mississippi',)
    ]
    expected_outputs = [
        {'h': 1, 'e': 1, 'l': 2, 'o': 1},
        {'m': 1, 'i': 4, 's': 4, 'p': 2}
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print('Test Passed')