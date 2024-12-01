# User function logic goes here
def count_down(n):
    for i in range(n, 0, -1):
        print(i)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (5,),
        (3,),
        (10,),
        (7,)
    ]
    expected_outputs = [
        '5\n4\n3\n2\n1',
        '3\n2\n1',
        '10\n9\n8\n7\n6\n5\n4\n3\n2\n1',
        '7\n6\n5\n4\n3\n2\n1'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = count_down(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)