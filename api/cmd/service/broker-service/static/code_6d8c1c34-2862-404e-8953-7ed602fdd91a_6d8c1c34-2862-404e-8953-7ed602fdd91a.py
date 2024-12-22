# User function logic goes here
def all_substrings_frequency(s):
    substr_freq = {}
    for i in range(len(s)):
        for j in range(i+1, len(s)+1):
            substr = s[i:j]
            substr_freq[substr] = substr_freq.get(substr, 0) + 1
    return substr_freq

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('banana',),
        ('abca',),
        ('abcabc',)
    ]
    expected_outputs = [
        {'b': 2, 'ba': 1, 'ban': 1, 'bana': 1, 'banan': 1, 'banana': 1, 'a': 3, 'an': 2, 'ana': 2, 'anan': 1, 'n': 2, 'na': 2, 'nan': 1},
        {'a': 2, 'ab': 1, 'abc': 1, 'b': 1, 'bc': 1, 'c': 1},
        {'a': 2, 'ab': 2, 'abc': 2, 'b': 2, 'bc': 2, 'c': 2}
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = all_substrings_frequency(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)