# User function logic goes here
def main(snake_case):
    return ''.join([word.capitalize() for word in snake_case.split('_')])

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        {'input': 'hello_world', 'expectedOutput': 'HelloWorld'},
        {'input': 'convert_to_pascal', 'expectedOutput': 'ConvertToPascal'},
        {'input': 'python_code', 'expectedOutput': 'PythonCode'},
        {'input': 'this_is_test', 'expectedOutput': 'ThisIsTest'}
    ]
    for case in test_cases:
        result = main(case['input'])
        if result != case['expectedOutput']:
            all_passed = False
    print(all_passed)