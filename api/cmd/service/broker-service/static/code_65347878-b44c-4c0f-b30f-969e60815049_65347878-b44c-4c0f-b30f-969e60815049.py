# User function logic goes here
def main(sentence):
    uppercase = sum(c.isupper() for c in sentence)
    lowercase = sum(c.islower() for c in sentence)
    return uppercase, lowercase

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('Hello World!',),
        ('PYTHON programming',),
        ('This Is A Test',),
        ('lowercaseonly',)
    ]
    expected_outputs = [
        (2, 8),
        (6, 10),
        (4, 7),
        (0, 13)
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)