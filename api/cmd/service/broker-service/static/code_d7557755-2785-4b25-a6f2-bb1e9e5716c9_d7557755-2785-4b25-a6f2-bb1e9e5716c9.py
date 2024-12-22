# User function logic goes here
def find_quadrant(x, y):
    if x > 0 and y > 0:
        return 'Quadrant I'
    elif x < 0 and y > 0:
        return 'Quadrant II'
    elif x < 0 and y < 0:
        return 'Quadrant III'
    elif x > 0 and y < 0:
        return 'Quadrant IV'
    else:
        return 'Origin'

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (1, 1),
        (-1, 1),
        (-1, -1),
        (1, -1)
    ]
    expected_outputs = [
        'Quadrant I',
        'Quadrant II',
        'Quadrant III',
        'Quadrant IV'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_quadrant(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)