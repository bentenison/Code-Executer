# User function logic goes here
def main(sentence):
    return len(sentence.split())

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        ('This is a test', 4),
        ('Hello world!', 2),
        ('Count the number of words in this sentence', 7)
    ]
    for sentence, expected in test_cases:
        result = main(sentence)
        if result != expected:
            print(f'Test failed for input {sentence}. Expected {expected}, got {result}')
    print('All test cases passed')