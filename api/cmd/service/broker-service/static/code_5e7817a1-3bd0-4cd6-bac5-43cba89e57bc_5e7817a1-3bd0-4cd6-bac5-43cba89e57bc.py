# User function logic goes here
def main(lst):
    unique_values = list(set(lst))
    unique_values.sort(reverse=True)
    return unique_values[1] if len(unique_values) > 1 else unique_values[0]

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        {'input': [10, 20, 4, 45, 99], 'expectedOutput': 45},
        {'input': [1, 2, 3, 4], 'expectedOutput': 3},
        {'input': [100, 100, 100], 'expectedOutput': 100},
        {'input': [5, 1], 'expectedOutput': 1}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            print(f'Test failed for input: {case['input']}. Expected: {case['expectedOutput']}, Got: {result}')
            exit(1)
    print('All tests passed.')