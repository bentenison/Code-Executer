# User function logic goes here
def main(score):
    if score >= 90:
        return 'A'
    elif score >= 80:
        return 'B'
    elif score >= 70:
        return 'C'
    elif score >= 60:
        return 'D'
    else:
        return 'F'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (95,),
        (85,),
        (75,),
        (65,),
        (55,)
    ]
    expected_outputs = [
        'A',
        'B',
        'C',
        'D',
        'F'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)