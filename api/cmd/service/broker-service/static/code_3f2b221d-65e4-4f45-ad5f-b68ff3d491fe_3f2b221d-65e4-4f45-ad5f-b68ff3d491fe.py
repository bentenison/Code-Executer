# User function logic goes here
import heapq

def rearrange_string(s):
    if not s:
        return ''
    freq_map = {}
    for char in s:
        freq_map[char] = freq_map.get(char, 0) + 1
    max_heap = [(-count, char) for char, count in freq_map.items()]
    heapq.heapify(max_heap)
    result = []
    prev_count, prev_char = 0, ''
    while max_heap:
        count, char = heapq.heappop(max_heap)
        result.append(char)
        if prev_count < 0:
            heapq.heappush(max_heap, (prev_count, prev_char))
        prev_count, prev_char = count + 1, char
    return ''.join(result) if len(result) == len(s) else ''

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