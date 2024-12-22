# User function logic goes here
def main(n):
    return sum(i ** 2 for i in range(1, n + 1))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (3,),
        (5,),
        (10,)
    ]
    expected_outputs = [
        14,
        55,
        385
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)