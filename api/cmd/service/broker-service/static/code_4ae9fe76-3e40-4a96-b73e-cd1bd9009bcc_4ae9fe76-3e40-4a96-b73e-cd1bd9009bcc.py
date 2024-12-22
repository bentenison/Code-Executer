# User function logic goes here
def main(s):
    import string
    return ''.join(char for char in s if char not in string.punctuation)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'Hello, World!', 'expectedOutput': 'Hello World'},
        {'input': 'Welcome!!!', 'expectedOutput': 'Welcome'},
        {'input': 'Python@3.10', 'expectedOutput': 'Python310'},
        {'input': 'C# is great!', 'expectedOutput': 'C is great'}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)