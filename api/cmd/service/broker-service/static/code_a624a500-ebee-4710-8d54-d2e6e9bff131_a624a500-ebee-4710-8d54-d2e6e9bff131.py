# User function logic goes here
def count_vowels(s):
    vowels = 'aeiou'
    count = 0
    for char in s:
        if char in vowels:
            count += 1
    return count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('world',),
        ('aeiou',),
        ('python',)
    ]
    expected_outputs = [
        2,
        1,
        5,
        1
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = count_vowels(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)