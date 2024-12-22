# User function logic goes here
def main(s):
    return [word for word in s.split() if len(word) % 2 == 0]

# Test cases for the function
