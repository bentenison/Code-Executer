# User function logic goes here
def main(tpl):
    return len(tpl)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': (1, 2, 3, 4), 'expectedOutput': 4},
        {'input': (), 'expectedOutput': 0},
        {'input': (100,), 'expectedOutput': 1},
        {'input': (10, 20, 30, 40, 50), 'expectedOutput': 5}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)