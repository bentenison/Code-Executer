def main(input_num):
    # User's main logic starts here
    if input_num < 2:
        return False
    for i in range(2, int(input_num ** 0.5) + 1):
        if input_num % i == 0:
            return False
    return True
    # User's main logic ends here

if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (2,),
        (17,),
        (15,)
    ]
    expected_outputs = [
        True,
        True,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)
