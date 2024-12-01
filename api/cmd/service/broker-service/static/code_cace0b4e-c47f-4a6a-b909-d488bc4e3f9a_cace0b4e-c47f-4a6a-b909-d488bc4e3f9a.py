# User function logic goes here
def find_maximum(numbers):
    max_num = numbers[0]
    for num in numbers:
        if num > max_num:
            max_num = num
    return max_num

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3],),
        ([10, 20, 5],),
        ([7, 2, 9],),
        ([15, 40, 30],)
    ]
    expected_outputs = [
        3,
        20,
        9,
        40
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_maximum(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)