# User function logic goes here
def main(sentence):
    letters = sum(c.isalpha() for c in sentence)
    digits = sum(c.isdigit() for c in sentence)
    return letters, digits

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('Hello World 123!',),
        ('2024 is here!',),
        ('Python 3.9',),
        ('NoDigitsHere',)
    ]
    expected_outputs = [
        (10, 3),
        (7, 4),
        (6, 1),
        (11, 0)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)