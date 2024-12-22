# User function logic goes here
def main(s):
    return ' '.join([word[0].upper() + word[1:-1] + word[-1].upper() if len(word) > 1 else word.upper() for word in s.split()])

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world',),
        ('capitalize each word',),
        ('python programming',),
        ('open ai',)
    ]
    expected_outputs = [
        'HellO WorlD',
        'CapitalizE EacH WorD',
        'PythoN ProgramminG',
        'OpeN AI'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)