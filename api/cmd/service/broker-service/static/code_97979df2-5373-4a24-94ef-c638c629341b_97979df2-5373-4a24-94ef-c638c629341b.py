# User function logic goes here
def main(lst):
    return lst[::-1]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4], [4, 3, 2, 1]),
        ([5, 6, 7], [7, 6, 5]),
        ([10, 20, 30, 40], [40, 30, 20, 10])
    ]
    for lst, expected in test_cases:
        result = main(lst)
        if result != expected:
            all_passed = False
    print(all_passed)