# User function logic goes here
def main(s):
    from collections import Counter
    freq = Counter(s)
    max_freq = max(freq.values())
    candidates = [ch for ch in freq if freq[ch] == max_freq]
    return min(candidates)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('testcase',),
        ('aaaaabbc',),
        ('abcabcabc',)
    ]
    expected_outputs = [
        'l',
        't',
        'a',
        'a'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)