db.questions.insertMany([
  {
    "title": "Palindrome Checker",
    "description": "Write a function that checks whether a given input is a palindrome.",
    "template_code": "def main(input):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        (\"{{ .Input }}\", {{ .ExpectedOutput }}),\n        {{ end }}\n    ]\n    for test_input, expected in test_cases:\n        result = main(test_input)\n        print(f'Input: {test_input}, Pass: {result == expected}')",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": "A man a plan a canal Panama", "expected_output": true},
      {"input": "racecar", "expected_output": true},
      {"input": "hello", "expected_output": false}
    ],
    "difficulty": "easy",
    "tags": ["string", "palindrome", "beginner"]
  },
  {
    "title": "FizzBuzz",
    "description": "Write a function that prints the numbers from 1 to n. But for multiples of three print 'Fizz' instead of the number and for the multiples of five print 'Buzz'. For numbers which are multiples of both three and five print 'FizzBuzz'.",
    "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for n in test_cases:\n        main(n)",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": 15, "expected_output": null},
      {"input": 3, "expected_output": null},
      {"input": 5, "expected_output": null}
    ],
    "difficulty": "easy",
    "tags": ["loops", "conditional", "beginner"]
  },
  {
    "title": "Two Sum",
    "description": "Given an array of integers, return indices of the two numbers such that they add up to a specific target.",
    "template_code": "def main(nums, target):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for nums, target in test_cases:\n        print(main(nums, target))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": [[2, 7, 11, 15], 9], "expected_output": [0, 1]},
      {"input": [[3, 2, 4], 6], "expected_output": [1, 2]},
      {"input": [[3, 3], 6], "expected_output": [0, 1]}
    ],
    "difficulty": "medium",
    "tags": ["array", "hashmap", "medium"]
  },
  {
    "title": "Reverse String",
    "description": "Write a function that reverses a string. The input string is given as an array of characters.",
    "template_code": "def main(s):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for s in test_cases:\n        print(main(s))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": ["h", "e", "l", "l", "o"], "expected_output": ["o", "l", "l", "e", "h"]},
      {"input": ["H", "a", "n", "n", "a", "h"], "expected_output": ["h", "a", "n", "n", "a", "H"]}
    ],
    "difficulty": "easy",
    "tags": ["string", "array", "easy"]
  },
  {
    "title": "Valid Anagram",
    "description": "Given two strings s and t, return true if t is an anagram of s and false otherwise.",
    "template_code": "def main(s, t):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for s, t in test_cases:\n        print(main(s, t))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": ["anagram", "nagaram"], "expected_output": true},
      {"input": ["rat", "car"], "expected_output": false}
    ],
    "difficulty": "easy",
    "tags": ["string", "hashmap", "easy"]
  },
  {
    "title": "Climbing Stairs",
    "description": "You are climbing a staircase. It takes n steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?",
    "template_code": "def main(n):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for n in test_cases:\n        print(main(n))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": 2, "expected_output": 2},
      {"input": 3, "expected_output": 3}
    ],
    "difficulty": "easy",
    "tags": ["dynamic programming", "array", "easy"]
  },
  {
    "title": "Maximum Subarray",
    "description": "Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.",
    "template_code": "def main(nums):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for nums in test_cases:\n        print(main(nums))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": [-2, 1, -3, 4, -1, 2, 1, -5, 4], "expected_output": 6},
      {"input": [1], "expected_output": 1}
    ],
    "difficulty": "easy",
    "tags": ["array", "dynamic programming", "easy"]
  },
  {
    "title": "Contains Duplicate",
    "description": "Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.",
    "template_code": "def main(nums):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for nums in test_cases:\n        print(main(nums))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": [1, 2, 3, 1], "expected_output": true},
      {"input": [1, 2, 3, 4], "expected_output": false}
    ],
    "difficulty": "easy",
    "tags": ["array", "hashmap", "easy"]
  },
  {
    "title": "Best Time to Buy and Sell Stock",
    "description": "You are given an array prices where prices[i] is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.",
    "template_code": "def main(prices):\n    # User's main logic starts here\n    {{ .Logic }}\n    # User's main logic ends here\n\nif __name__ == '__main__':\n    test_cases = [\n        {{ range .TestCases }}\n        ({{ .Input }}),\n        {{ end }}\n    ]\n    for prices in test_cases:\n        print(main(prices))",
    "language": "python",
    "language_code": "py",
    "test_cases": [
      {"input": [7, 1, 5, 3, 6, 4], "expected_output": 5},
      {"input": [7, 6, 4, 3, 1], "expected_output": 0}
    ],
    "difficulty": "easy",
    "tags": ["array", "dynamic programming", "easy"]
  }
])

















When you write code, execute it within a Docker container running Python, and then read its stdout, several things can happen depending on the code's logic and structure. Here’s a breakdown of the process and possible outcomes:

### 1. **Setup**: 
   - You would create a Docker container with Python installed.
   - You would copy your Python script or the template into the container.
   - You would run the script, which may include dynamically generated code and test cases.

### 2. **Execution Flow**:
   - The Python interpreter inside the Docker container will execute your script. If you’ve integrated the logic for handling multiple test cases using a text template, it will execute the user's logic against those test cases.

### 3. **Standard Output (stdout)**:
   - **Successful Execution**: If the code executes successfully and passes all test cases, it will print the expected outputs, or any success messages you’ve coded in.
   - **Errors**: If there are errors in the code (e.g., syntax errors, exceptions), the interpreter will print error messages to stdout, which may include:
     - Syntax errors (e.g., missing colons, unmatched parentheses)
     - Runtime errors (e.g., division by zero, index out of range)
     - Logic errors (where the program runs without crashing but produces incorrect results).

### 4. **Collecting stdout**:
   - You can capture the stdout from the Docker container. Here’s how you can do it:
     - Use the `docker run` command with `-it` to run the container interactively, or `-d` for detached mode. 
     - You can redirect stdout to a file or read it directly in your application using Docker SDK for Python or similar tools.
   - Example command to run and capture output:
     ```bash
     docker run --rm python:3.9 python /path/to/your_script.py > output.txt
     ```

### 5. **Potential Scenarios**:
   - **All Test Cases Pass**: Outputs indicating successful tests, such as "Pass" or "All tests passed."
   - **Some Test Cases Fail**: Outputs showing which test cases failed and their expected versus actual outputs.
   - **Errorc Messages**: If there are any issues in the code, it will print stack traces or error messages to stdout.

### Example of Capturing Output

Here’s a simple example of how you might structure your code execution and output capture:

1. **Dockerfile** (to set up your environment):
    ```dockerfile
    FROM python:3.9

    WORKDIR /app

    COPY your_script.py .

    CMD ["python", "your_script.py"]
    ```

2. **Python Script** (`your_script.py`):
    ```python
    def main(input):
        # User's main logic starts here
        cleaned_input = input.lower().replace(" ", "")
        return cleaned_input == cleaned_input[::-1]
        # User's main logic ends here

    if __name__ == "__main__":
        test_cases = [
            ("A man a plan a canal Panama", True),
            ("racecar", True),
            ("hello", False)
        ]
        
        for test_input, expected in test_cases:
            result = main(test_input)
            print(f'Input: {test_input}, Expected: {expected}, Got: {result}')
    ```

3. **Run and Capture Output**:
   ```bash
   docker build -t python-test .
   docker run --rm python-test > output.txt
   cat output.txt
   ```

### Conclusion
Overall, when executing code in a Docker Python container and capturing stdout, you'll either receive outputs that confirm the correctness of your code (test results) or error messages that help you debug your code. This approach is useful for testing user-submitted code in a controlled and isolated environment.