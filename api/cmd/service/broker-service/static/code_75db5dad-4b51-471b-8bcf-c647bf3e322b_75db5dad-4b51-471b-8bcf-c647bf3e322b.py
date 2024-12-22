# User function logic goes here
def main(sentence):
    return len(sentence.split())

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('Hello world', 2),
        ('Python is amazing', 3),
        ('This is a longer sentence.', 5)
    ]
    for sentence, expected_output in test_cases:
        result = main(sentence)
        if result != expected_output:
            all_passed = False
            print(f'Failed for Input: {sentence}. Expected: {expected_output}, Got: {result}')
    print('All Passed:', all_passed)