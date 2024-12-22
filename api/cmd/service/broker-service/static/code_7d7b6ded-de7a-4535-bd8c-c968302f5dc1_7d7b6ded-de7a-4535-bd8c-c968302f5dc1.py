# User function logic goes here
def main(s):
    return s == s[::-1]

# Test cases for the function
def run_tests():
    test_cases = [
        {'input': 'madam', 'expectedOutput': True},
        {'input': 'hello', 'expectedOutput': False},
        {'input': 'racecar', 'expectedOutput': True},
        {'input': 'world', 'expectedOutput': False}
    ]
    
    all_passed = True
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
            print(f"Failed for Input: {case['input']}. Expected: {case['expectedOutput']}, Got: {result}")
    
    return all_passed

if __name__ == '__main__':
    print(run_tests())