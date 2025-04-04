# User function logic goes here
def sum_of_even_numbers(nums):
    return sum(num for num in nums if num % 2 == 0)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4],),
        ([10, 15, 20, 25],),
        ([5, 7, 9],),
        ([2, 4, 6, 8],)
    ]
    expected_outputs = [
        6,
        30,
        0,
        20
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = sum_of_even_numbers(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)