# User function logic goes here
def main(str1, str2):
    return sorted(str1) == sorted(str2)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': ('listen', 'silent'), 'expectedOutput': True},
        {'input': ('triangle', 'integral'), 'expectedOutput': True},
        {'input': ('hello', 'world'), 'expectedOutput': False},
        {'input': ('abc', 'cba'), 'expectedOutput': True}
    ]
    for case in test_cases:
        result = main(case['input'][0], case['input'][1])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)