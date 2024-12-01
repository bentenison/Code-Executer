#include <stdio.h>

// User function logic goes here
int sumOfDigits(int num) {
    int sum = 0;
    while (num > 0) {
        sum += num % 10;
        num /= 10;
    }
    return sum;
}

// Test cases for the function
int main() {
    int test_cases[] = {123, 456, 789};
    int expected[] = {6, 15, 24};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int result = sumOfDigits(test_cases[i]);
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