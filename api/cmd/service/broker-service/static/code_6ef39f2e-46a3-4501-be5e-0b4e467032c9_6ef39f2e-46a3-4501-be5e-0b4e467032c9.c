#include <stdio.h>

// User function logic goes here
int fibonacci(int n) {
    if (n == 0) return 0;
    if (n == 1) return 1;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

// Test cases for the function
int main() {
    int test_cases[] = {0, 1, 2, 3, 4, 5};
    int expected[] = {0, 1, 1, 2, 3, 5};
    int all_passed = 1;

    for (int i = 0; i < 6; i++) {
        int result = fibonacci(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1; // Return 1 to indicate true
    } else {
        printf("false\n");
        return 0; // Return 0 to indicate false
    }
}