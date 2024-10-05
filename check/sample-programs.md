Here are some Python problems based on different programming concepts, formatted using the template structure you provided. Each problem introduces a specific concept and includes test input, expected output, and placeholders for the user's logic.

---

### 1. **Problem 1: Palindrome Check (String Manipulation)**
#### Description:
Write a function that checks if a given string is a palindrome. A palindrome is a word, phrase, or sequence that reads the same backward as forward.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    cleaned_input = input.lower().replace(" ", "")
    return cleaned_input == cleaned_input[::-1]
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "A man a plan a canal Panama"  # User input for the test case
    expected_output = "True"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 2. **Problem 2: Factorial Calculation (Recursion)**
#### Description:
Write a function to calculate the factorial of a given number using recursion. The factorial of a number `n` is defined as `n! = n * (n-1) * (n-2) * ... * 1`.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    def factorial(n):
        return 1 if n == 0 else n * factorial(n - 1)
    
    return factorial(int(input))
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "5"  # User input for the test case
    expected_output = "120"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 3. **Problem 3: Fibonacci Sequence (Dynamic Programming)**
#### Description:
Write a function to find the nth Fibonacci number. The Fibonacci sequence is defined as:
- `Fib(0) = 0`
- `Fib(1) = 1`
- `Fib(n) = Fib(n-1) + Fib(n-2)`

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    def fibonacci(n):
        a, b = 0, 1
        for _ in range(n):
            a, b = b, a + b
        return a
    
    return fibonacci(int(input))
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "10"  # User input for the test case
    expected_output = "55"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 4. **Problem 4: Sum of Elements in a List (List Operations)**
#### Description:
Write a function that takes a list of numbers and returns the sum of all the numbers in the list.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    numbers = [int(x) for x in input.split(",")]
    return sum(numbers)
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "1,2,3,4,5"  # User input for the test case
    expected_output = "15"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 5. **Problem 5: Count Vowels in a String (String Processing)**
#### Description:
Write a function that counts the number of vowels (a, e, i, o, u) in a given string.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    vowels = "aeiou"
    return sum(1 for char in input.lower() if char in vowels)
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "Hello World"  # User input for the test case
    expected_output = "3"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 6. **Problem 6: Prime Number Check (Mathematical Logic)**
#### Description:
Write a function that checks if a given number is prime. A prime number is a number greater than 1 that has no divisors other than 1 and itself.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    n = int(input)
    if n < 2:
        return False
    for i in range(2, int(n ** 0.5) + 1):
        if n % i == 0:
            return False
    return True
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "29"  # User input for the test case
    expected_output = "True"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 7. **Problem 7: Sorting a List (Sorting Algorithms)**
#### Description:
Write a function that takes a list of numbers and returns a sorted list in ascending order.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    numbers = [int(x) for x in input.split(",")]
    return sorted(numbers)
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "5,3,8,1,2"  # User input for the test case
    expected_output = "[1, 2, 3, 5, 8]"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 8. **Problem 8: Reverse a String (String Manipulation)**
#### Description:
Write a function that takes a string and returns the string reversed.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    return input[::-1]
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "hello"  # User input for the test case
    expected_output = "olleh"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 9. **Problem 9: GCD of Two Numbers (Euclidean Algorithm)**
#### Description:
Write a function to calculate the greatest common divisor (GCD) of two numbers using the Euclidean algorithm.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    a, b = map(int, input.split(","))
    while b:
        a, b = b, a % b
    return a
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "48,18"  # User input for the test case
    expected_output = "6"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

### 10. **Problem 10: Find Duplicates in a List (Data Structures - Set)**
#### Description:
Write a function that finds and returns a list of duplicate elements in a given list.

#### Python Problem:
```python
def main(input):
    # User's main logic starts here
    numbers = [int(x) for x in input.split(",")]
    seen = set()
    duplicates = set(x for x in numbers if x in seen or seen.add(x))
    return sorted(duplicates)
    # User's main logic ends here

if __name__ == "__main__":
    test_input = "1,2,3,4,2,3,5"  # User input for the test case
    expected_output = "[2, 3]"  # Expected output
    
    result = str(main(test_input))
    print("Pass" if result == expected_output else "Fail")
```

---

These problems cover a range of concepts including string manipulation, recursion, dynamic programming, list operations, mathematical logic, and data structures. Each problem includes a clear input and output format, with the user's main logic being inserted into the provided template.