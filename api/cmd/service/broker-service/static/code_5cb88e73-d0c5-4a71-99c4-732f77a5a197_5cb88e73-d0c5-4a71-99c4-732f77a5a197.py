# User function logic goes here
from collections import Counter

def main(s):
    freq = Counter(s)
    return min(freq, key=freq.get)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', 'h'),
        ('aabbcc', 'a'),
        ('python', 'p')
    ]
    for s, expected in test_cases:
        result = main(s)
        if result != expected:
            all_passed = False
    print(all_passed)