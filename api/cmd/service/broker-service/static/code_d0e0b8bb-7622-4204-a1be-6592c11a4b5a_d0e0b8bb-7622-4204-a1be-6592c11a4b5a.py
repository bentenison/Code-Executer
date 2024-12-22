# User function logic goes here
def main(sentence):
    return ' '.join(sentence.split()[::-1])

# Test cases for the function
if __name__ == '__main__':
    test_cases = [
        {'input': 'hello world', 'expectedOutput': 'world hello'},
        {'input': 'python is fun', 'expectedOutput': 'fun is python'},
        {'input': 'reverse me', 'expectedOutput': 'me reverse'},
        {'input': 'a b c', 'expectedOutput': 'c b a'}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            print(f'Test failed for input: {case['input']}. Expected: {case['expectedOutput']}, Got: {result}')
            exit(1)
    print('All tests passed.')