#include <stdio.h>

// User function logic goes here
int checkEvenOdd(int num) {
    if (num % 2 == 0) {
        return 1; // Even
    }
    return 0; // Odd
}

// Test cases for the function
int main() {
    int test_cases[] = {4, 7, 0, 15, 22};
    int expected[] = {1, 0, 1, 0, 1};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = checkEvenOdd(test_cases[i]);
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