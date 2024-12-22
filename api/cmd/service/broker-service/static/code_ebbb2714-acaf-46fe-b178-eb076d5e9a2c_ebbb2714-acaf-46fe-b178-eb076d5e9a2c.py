# User function logic goes here
def main(num):
    result = str(num)
    print(result)
    return result

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [123, 0, -45, 1001]
    expected_outputs = ['123', '0', '-45', '1001']
    for i, test_input in enumerate(test_cases):
        result = main(test_input)
        if result != expected_outputs[i]:
            all_passed = False
    print(all_passed)