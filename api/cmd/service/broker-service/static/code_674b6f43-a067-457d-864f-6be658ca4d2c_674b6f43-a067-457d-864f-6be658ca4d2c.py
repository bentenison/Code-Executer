# User function logic goes here
def main(lst):
    even_count = sum(1 for num in lst if num % 2 == 0)
    odd_count = len(lst) - even_count
    return even_count, odd_count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4, 5],),
        ([6, 7, 8, 9],),
        ([11, 12, 13],),
        ([2, 4, 6],)
    ]
    expected_outputs = [
        (2, 3),
        (2, 2),
        (1, 2),
        (3, 0)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)