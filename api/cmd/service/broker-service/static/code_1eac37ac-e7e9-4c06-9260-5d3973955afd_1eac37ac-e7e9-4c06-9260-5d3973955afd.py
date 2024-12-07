# User function logic goes here
def main(temp, conversion_type):
    if conversion_type == 'CtoF':
        return (temp * 9/5) + 32
    elif conversion_type == 'FtoC':
        return (temp - 32) * 5/9

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (0, 'CtoF'),
        (32, 'FtoC'),
        (100, 'CtoF'),
        (212, 'FtoC')
    ]
    expected_outputs = [
        32.0,
        0.0,
        212.0,
        100.0
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)