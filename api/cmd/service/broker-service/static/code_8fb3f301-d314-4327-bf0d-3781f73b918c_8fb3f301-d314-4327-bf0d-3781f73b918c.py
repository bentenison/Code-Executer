# User function logic goes here
def main(s):
    n = len(s)
    if n == 0:
        return 0
    table = [[False] * n for _ in range(n)]
    max_len = 1
    for i in range(n):
        table[i][i] = True
    for start in range(n-2, -1, -1):
        for end in range(start+1, n):
            if s[start] == s[end] and (end - start == 1 or table[start+1][end-1]):
                table[start][end] = True
                max_len = max(max_len, end - start + 1)
    return max_len

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('babad',),
        ('cbbd',),
        ('a',),
        ('ac',)
    ]
    expected_outputs = [
        3,
        2,
        1,
        1
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)