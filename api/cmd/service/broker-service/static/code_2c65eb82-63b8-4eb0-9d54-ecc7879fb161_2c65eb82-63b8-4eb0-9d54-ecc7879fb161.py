# User function logic goes here
def main(s):
    return len(s.split())

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'Hello world!', 'expectedOutput': 2},
        {'input': 'Count the words in this sentence.', 'expectedOutput': 6},
        {'input': 'SingleWord', 'expectedOutput': 1},
        {'input': '  Spaces   before  and after  ', 'expectedOutput': 5}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)