# User function logic goes here
def main(str1, str2):
    result = str1 + str2
    print(result)
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [("Hello", "World"), ("Python", "Rocks"), ("123", "456"), ("", "Test")]
    expected_outputs = ["HelloWorld", "PythonRocks", "123456", "Test"]
    for i, test_input in enumerate(test_cases):
        result = main(test_input[0], test_input[1])
        if result != expected_outputs[i]:
            all_passed = False
    print(all_passed)