# User function logic goes here
def main(n):
    # Your code here
    if n < 0:
        return 0
    elif n == 0:
        return 1
    result = 1
    for i in range(1, n + 1):
        result *= i
    print(result)
    return result
if __name__ == '__main__':
    main(5)

# Test cases for the function
