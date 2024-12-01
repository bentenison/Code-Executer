# User function logic goes here
def fibonacci(n):
    a, b = 0, 1
    for _ in range(n):
        print(a)
        a, b = b, a + b

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5,),
        (8,),
        (10,),
        (3,)
    ]
    expected_outputs = [
        '0\n1\n1\n2\n3',
        '0\n1\n1\n2\n3\n5\n8\n13',
        '0\n1\n1\n2\n3\n5\n8\n13\n21\n34',
        '0\n1\n1'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = fibonacci(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)