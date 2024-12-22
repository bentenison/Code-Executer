# User function logic goes here
import re

def main(s):
    return bool(re.search(r'[^a-zA-Z0-9]', s))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello@world', True),
        ('helloworld', False),
        ('python123', False),
        ('!@#$%^&*()', True)
    ]
    for s, expected in test_cases:
        result = main(s)
        if result != expected:
            all_passed = False
    print(all_passed)