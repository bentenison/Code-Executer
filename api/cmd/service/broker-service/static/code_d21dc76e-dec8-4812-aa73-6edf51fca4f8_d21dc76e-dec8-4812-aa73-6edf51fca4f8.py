# User function logic goes here
from collections import Counter

def count_word_frequency(text):
    words = text.split()
    return dict(Counter(words))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world hello',),
        ('apple banana apple orange',),
    ]
    expected_outputs = [
        {'hello': 2, 'world': 1},
        {'apple': 2, 'banana': 1, 'orange': 1},
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = count_word_frequency(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)