def main(n):
    # User's main logic starts here
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True
    # User's main logic ends here

if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (2,),
        (11,),
        (15,),
        (1,)
    ]
    expected_outputs = [
        True,
        True,
        False,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)