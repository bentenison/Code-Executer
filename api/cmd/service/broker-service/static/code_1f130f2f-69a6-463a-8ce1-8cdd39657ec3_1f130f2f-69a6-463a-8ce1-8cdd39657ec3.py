# User function logic goes here
def main(year):
    return year % 100 == 0 and year % 400 == 0

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (2000,),
        (1900,),
        (2024,),
        (1800,)
    ]
    expected_outputs = [
        True,
        True,
        False,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)