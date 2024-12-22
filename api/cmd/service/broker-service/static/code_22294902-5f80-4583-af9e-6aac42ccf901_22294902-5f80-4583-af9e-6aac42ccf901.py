# User function logic goes here
def grade_student(marks):
    if marks >= 90:
        return 'A'
    elif marks >= 70:
        return 'B'
    elif marks >= 50:
        return 'C'
    elif marks >= 30:
        return 'D'
    return 'F'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (95,),
        (70,),
        (50,)
    ]
    expected_outputs = [
        'A',
        'B',
        'C'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = grade_student(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)