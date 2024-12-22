# User function logic goes here
def main(lst):
    result = {}
    for word in lst:
        key = word[0]
        result.setdefault(key, []).append(word)
    return result

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        (['apple', 'banana', 'apricot', 'blueberry'],),
        (['cat', 'dog', 'duck', 'dove'],)
    ]
    expected_outputs = [
        {'a': ['apple', 'apricot'], 'b': ['banana', 'blueberry']},
        {'c': ['cat'], 'd': ['dog', 'duck', 'dove']}
    ]
    all_passed = True
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)