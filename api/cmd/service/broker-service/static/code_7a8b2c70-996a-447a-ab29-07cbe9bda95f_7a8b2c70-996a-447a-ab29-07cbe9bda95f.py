# User function logic goes here
def main(num):
    total = sum(int(digit) ** len(str(num)) for digit in str(num))
    return total == num

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 153, 'expectedOutput': True},
        {'input': 9474, 'expectedOutput': True},
        {'input': 123, 'expectedOutput': False},
        {'input': 9475, 'expectedOutput': False}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)