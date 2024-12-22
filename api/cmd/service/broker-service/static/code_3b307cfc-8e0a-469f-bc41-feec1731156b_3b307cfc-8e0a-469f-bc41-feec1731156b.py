# User function logic goes here
def main(s):
    from collections import Counter
    freq = Counter(s)
    min_freq = min(freq.values())
    for char in s:
        if freq[char] == min_freq:
            return char

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('aabbccdde',),
        ('abcabc',),
        ('aaaabbbbcccd',)
    ]
    expected_outputs = [
        'e',
        'a',
        'd'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)