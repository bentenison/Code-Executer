# User function logic goes here
def main(s, sub):
    return s.count(sub)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world', 'o'),
        ('banana', 'na')
    ]
    expected_outputs = [
        2,
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)