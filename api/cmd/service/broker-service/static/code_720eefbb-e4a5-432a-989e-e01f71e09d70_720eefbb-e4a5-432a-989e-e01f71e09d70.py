# User function logic goes here
def main(s):
    from collections import Counter
    freq = Counter(s)
    odd_freq_chars = [ch for ch in freq if freq[ch] % 2 != 0]
    return ''.join(sorted(odd_freq_chars))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('testcase',),
        ('abcabcabc',),
        ('aabbcc',)
    ]
    expected_outputs = [
        'eh',
        'aes',
        'a',
        ''
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)