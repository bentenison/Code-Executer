# User function logic goes here
def main():
    return {i: i**2 for i in range(1, 6)}

# Test cases for the function
if __name__ == '__main__':
    result = main()
    expected = {1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
    print(result == expected)