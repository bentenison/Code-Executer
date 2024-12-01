# User function logic goes here
def print_multiples(number, limit):
    for i in range(number, limit + 1, number):
        print(i)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (3, 20),
        (5, 50),
        (7, 30),
        (2, 15)
    ]
    expected_outputs = [
        '3\n6\n9\n12\n15\n18\n20',
        '5\n10\n15\n20\n25\n30\n35\n40\n45\n50',
        '7\n14\n21\n28',
        '2\n4\n6\n8\n10\n12\n14\n15'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = print_multiples(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)