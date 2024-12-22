# User function logic goes here
def find_words_greater_than_k(s, k):
    words = s.split()
    return [word for word in words if len(word) > k]

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('The quick brown fox', 3),
        ('Hello world', 4),
        ('I am here', 1),
        ('abc def ghi', 2)
    ]
    expected_outputs = [
        ['quick', 'brown', 'fox'],
        ['Hello', 'world'],
        ['I', 'am', 'here'],
        ['def', 'ghi']
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = find_words_greater_than_k(test_input[0], test_input[1])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)