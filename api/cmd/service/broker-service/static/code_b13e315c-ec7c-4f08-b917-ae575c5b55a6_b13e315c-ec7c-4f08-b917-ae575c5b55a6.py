# User function logic goes here
def remove_duplicates(numbers):
    result = []
    for num in numbers:
        if num not in result:
            result.append(num)
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 2, 3],),
        ([4, 5, 6, 5, 4],),
        ([10, 20, 20, 10],),
        ([7, 8, 7, 8, 7],)
    ]
    expected_outputs = [
        [1, 2, 3],
        [4, 5, 6],
        [10, 20],
        [7, 8]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = remove_duplicates(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)