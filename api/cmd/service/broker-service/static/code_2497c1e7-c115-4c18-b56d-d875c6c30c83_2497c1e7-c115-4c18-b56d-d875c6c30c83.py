# User function logic goes here
import string

def main(text):
    return ''.join(char for char in text if char not in string.punctuation)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'Hello, World!', 'expectedOutput': 'Hello World'},
        {'input': 'Python@3.9!', 'expectedOutput': 'Python39'},
        {'input': 'No#punctuation$here', 'expectedOutput': 'Nopunctuationhere'},
        {'input': '', 'expectedOutput': ''}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)