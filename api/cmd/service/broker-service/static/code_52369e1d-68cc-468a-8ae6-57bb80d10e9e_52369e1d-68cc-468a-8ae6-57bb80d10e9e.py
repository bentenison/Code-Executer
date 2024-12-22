# User function logic goes here
def main(year):
    if (year % 4 == 0 and year % 100 != 0) or (year % 400 == 0):
        return True
    return False

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (2000,),
        (1900,),
        (2020,),
        (2024,)
    ]
    expected_outputs = [
        True,
        False,
        True,
        True
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)