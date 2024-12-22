# User function logic goes here
def main(s, n):
    n = n % len(s)  # Handle large n values
    return s[n:] + s[:n]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello', 2, 'llohe'),
        ('world', 3, 'ldwor'),
        ('python', 6, 'python'),
        ('abcd', 1, 'bcda')
    ]
    for s, n, expected in test_cases:
        result = main(s, n)
        if result != expected:
            all_passed = False
    print(all_passed)