#include <stdio.h>

// User function logic goes here
int checkSign(int num) {
    if (num > 0) {
        return 1; // Positive
    } else if (num < 0) {
        return -1; // Negative
    }
    return 0; // Zero
}

// Test cases for the function
int main() {
    int test_cases[] = {5, -10, 0, 100, -25};
    int expected[] = {1, -1, 0, 1, -1};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = checkSign(test_cases[i]);
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