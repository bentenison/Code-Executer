# User function logic goes here
def main(s):
    vowels = 'aeiou'
    count_vowels = sum(1 for char in s if char in vowels)
    count_consonants = len(s) - count_vowels
    return count_vowels, count_consonants

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello',),
        ('world',),
        ('python',),
        ('aeiou',)
    ]
    expected_outputs = [
        (2, 3),
        (1, 4),
        (1, 5),
        (5, 0)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)