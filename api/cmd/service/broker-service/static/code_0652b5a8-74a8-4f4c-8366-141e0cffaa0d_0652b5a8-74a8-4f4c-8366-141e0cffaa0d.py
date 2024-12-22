# User function logic goes here
def main(s, i):
    return s[:i] + s[i+1:]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', 2, 'helo'),
        ('world', 3, 'wold'),
        ('python', 0, 'ython')
    ]
    for s, i, expected in test_cases:
        result = main(s, i)
        if result != expected:
            all_passed = False
    print(all_passed)