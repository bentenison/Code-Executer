# User function logic goes here
def main(lst):
    return sum(lst)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4], 10),
        ([5, 6, 7], 18),
        ([10, 20, 30, 40], 100)
    ]
    for lst, expected in test_cases:
        result = main(lst)
        if result != expected:
            all_passed = False
    print(all_passed)