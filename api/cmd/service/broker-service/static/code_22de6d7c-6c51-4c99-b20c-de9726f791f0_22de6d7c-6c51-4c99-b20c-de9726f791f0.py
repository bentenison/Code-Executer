# User function logic goes here
def main(s):
    return [word for word in s.split() if len(word) % 2 == 0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'This is a test string', 'expectedOutput': ['This', 'is', 'test']},
        {'input': 'Python programming is fun', 'expectedOutput': ['Python', 'is']},
        {'input': 'I love coding', 'expectedOutput': ['love']},
        {'input': 'Even length words only', 'expectedOutput': ['Even', 'length', 'words']}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
            print(f'Failed for Input: {case["input"]}. Expected: {case["expectedOutput"]}, Got: {result}')
    print(all_passed)