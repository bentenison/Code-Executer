# User function logic goes here
def main(s):
    return s.count(' ')

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world',),
        ('count spaces in string',),
        ('python  programming',),
        ('no_space_here',)
    ]
    expected_outputs = [
        1,
        3,
        2,
        0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)