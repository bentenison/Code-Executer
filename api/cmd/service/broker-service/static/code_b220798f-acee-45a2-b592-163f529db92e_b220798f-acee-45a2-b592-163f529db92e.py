# User function logic goes here
def main(sentence):
    return [word for word in sentence.split() if len(word) % 2 == 0]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'This is a test', 'expectedOutput': ['This', 'is', 'test']},
        {'input': 'Hello world', 'expectedOutput': ['world']},
        {'input': 'Python is fun', 'expectedOutput': ['Python']},
        {'input': '', 'expectedOutput': []}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)