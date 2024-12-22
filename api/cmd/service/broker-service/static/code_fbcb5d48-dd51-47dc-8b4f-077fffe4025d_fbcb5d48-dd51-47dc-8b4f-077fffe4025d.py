# User function logic goes here
def findLongestChain(pairs):
    pairs.sort(key=lambda x: x[1])
    count, end = 0, float('-inf')
    for pair in pairs:
        if pair[0] > end:
            count += 1
            end = pair[1]
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([[1,2], [2,3], [3,4]])
    ]
    expected_outputs = [
        2
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = findLongestChain(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)