# User function logic goes here
from collections import Counter
import heapq

def rearrange_string(s):
    if not s:
        return ""
    freq_map = Counter(s)
    max_heap = [(-freq, char) for char, freq in freq_map.items()]
    heapq.heapify(max_heap)
    prev_char = None
    prev_freq = 0
    result = []

    while max_heap:
        freq, char = heapq.heappop(max_heap)
        result.append(char)
        if prev_freq < 0:
            heapq.heappush(max_heap, (prev_freq, prev_char))
        prev_char = char
        prev_freq = freq + 1

    return "".join(result) if len(result) == len(s) else ""

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('aab',),
        ('aaab',),
        ('abc',),
        ('',)
    ]
    expected_outputs = [
        'aba',
        '',
        'abc',
        ''
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = rearrange_string(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)