# User function logic goes here
def main(lst, element):
    return element in lst

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4], 3, True),
        ([1, 2, 3, 4], 5, False),
        ([10, 20, 30], 20, True),
        ([1, 1, 1], 2, False)
    ]
    for lst, element, expected in test_cases:
        result = main(lst, element)
        if result != expected:
            all_passed = False
    print(all_passed)