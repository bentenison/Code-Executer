# User function logic goes here
def main(a, b, c):
    return max(a, b, c)

# Test cases for the function
def validate():
    test_cases = [
        (3, 7, 5, 7),
        (10, 5, 7, 10),
        (1, 1, 1, 1),
        (0, -5, 2, 2)
    ]
    for a, b, c, expected in test_cases:
        result = main(a, b, c)
        if result != expected:
            return False
    return True