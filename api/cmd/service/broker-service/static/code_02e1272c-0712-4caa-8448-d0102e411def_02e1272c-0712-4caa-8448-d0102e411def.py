# User function logic goes here
from collections import Counter

def main(s):
    counts = Counter(s)
    return [char for char, count in counts.items() if count > 1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', ['l']),
        ('test', ['t']),
        ('world', []),
        ('aabbcc', ['a', 'b', 'c'])
    ]
    for s, expected in test_cases:
        result = main(s)
        if result != expected:
            all_passed = False
    print(all_passed)