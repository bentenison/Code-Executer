# User function logic goes here
def main(n):
    return n > 0 and 1162261467 % n == 0

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (27,),
        (9,),
        (18,),
        (81,)
    ]
    expected_outputs = [
        True,
        True,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)