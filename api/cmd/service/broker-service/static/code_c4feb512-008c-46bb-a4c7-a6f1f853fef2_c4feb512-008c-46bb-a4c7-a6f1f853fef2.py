# User function logic goes here
def maximalSquare(matrix):
    if not matrix:
        return 0
    rows, cols = len(matrix), len(matrix[0])
    dp = [[0] * (cols + 1) for _ in range(rows + 1)]
    max_side = 0
    for i in range(1, rows + 1):
        for j in range(1, cols + 1):
            if matrix[i - 1][j - 1] == '1':
                dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]) + 1
                max_side = max(max_side, dp[i][j])
    return max_side * max_side

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([['1','0','1','0','0'],['1','0','1','1','1'],['1','1','1','1','1'],['1','0','0','1','0']])
    ]
    expected_outputs = [
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maximalSquare(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)