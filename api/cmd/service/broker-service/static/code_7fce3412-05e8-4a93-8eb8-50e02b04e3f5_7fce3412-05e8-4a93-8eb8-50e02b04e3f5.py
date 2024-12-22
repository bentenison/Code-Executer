# User function logic goes here
def find_pairs(arr, target):
    pairs = []
    seen = set()
    for num in arr:
        complement = target - num
        if complement in seen:
            pairs.append((complement, num))
        seen.add(num)
    return pairs

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 2, 3, 4, 5], 6),
        ([1, 3, 2, 4, 6], 6)
    ]
    expected_outputs = [
        [(1, 5), (2, 4)],
        [(2, 4)]
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_pairs(*test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)