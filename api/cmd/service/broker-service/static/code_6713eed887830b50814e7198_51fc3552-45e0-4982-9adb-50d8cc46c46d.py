def main(n):
    # User's main logic starts here
     if n <= 0:
        return []
    elif n == 1:
        return [0]
    
    fib = [0, 1]
    for i in range(2, n):
        fib.append(fib[i - 1] + fib[i - 2])
    return fib
    # User's main logic ends here

if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5,),
        (0,),
        (10,),
        (1,)
    ]
    expected_outputs = [
        [0, 1, 1, 2, 3],
        [],
        [0, 1, 1, 2, 3, 5, 8, 13, 21, 34],
        [0]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)