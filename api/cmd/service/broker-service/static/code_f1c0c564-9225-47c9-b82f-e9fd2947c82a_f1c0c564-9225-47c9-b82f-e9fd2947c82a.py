# User function logic goes here
def reverse_words(sentence):
    words = sentence.split()
    return ' '.join(reversed(words))

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ('hello world',),
        ('this is a test',),
        ('python is great',),
        ('openai is amazing',)
    ]
    expected_outputs = [
        'world hello',
        'test a is this',
        'great is python',
        'amazing is openai'
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = reverse_words(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)