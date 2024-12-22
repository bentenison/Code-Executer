# User function logic goes here
def main(s1, s2):
    return sorted(s1) == sorted(s2)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': ('listen', 'silent'), 'expectedOutput': True},
        {'input': ('triangle', 'integral'), 'expectedOutput': True},
        {'input': ('apple', 'pale'), 'expectedOutput': False},
        {'input': ('aabbcc', 'abcabc'), 'expectedOutput': True}
    ]
    for case in test_cases:
        result = main(case['input'][0], case['input'][1])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)