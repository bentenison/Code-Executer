# User function logic goes here
def isValidSudoku(board):
    rows = [{} for _ in range(9)]
    cols = [{} for _ in range(9)]
    boxes = [{} for _ in range(9)]
    for i in range(9):
        for j in range(9):
            if board[i][j] != '.':
                num = board[i][j]
                k = (i // 3) * 3 + j // 3
                if num in rows[i] or num in cols[j] or num in boxes[k]:
                    return False
                rows[i][num] = 1
                cols[j][num] = 1
                boxes[k][num] = 1
    return True

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([['5','3','.','.','7','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']]),
        ([['8','3','.','.','.','.','.','.','.'],['6','.','.','1','9','5','.','.','.'],['.','9','8','.','.','.','.','6','.'],['8','.','.','.','6','.','.','.','3'],['4','.','.','8','.','3','.','.','1'],['7','.','.','.','2','.','.','.','6'],['.','6','.','.','.','.','2','8','.'],['.','.','.','4','1','9','.','.','5'],['.','.','.','.','8','.','.','7','9']])
    ]
    expected_outputs = [
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = isValidSudoku(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)