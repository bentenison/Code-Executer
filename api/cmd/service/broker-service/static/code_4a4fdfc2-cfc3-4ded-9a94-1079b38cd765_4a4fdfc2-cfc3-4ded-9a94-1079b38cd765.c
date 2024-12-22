#include <stdio.h>

// User function logic goes here
int sumOddFactors(int n) {
    int sum = 0;
    for (int i = 1; i <= n; i += 2) {
        if (n % i == 0) {
            sum += i;
        }
    }
    return sum;
}

// Test cases for the function
int main() {
    int test_cases[] = {30, 15, 12, 21, 25};
    int expected[] = {10, 24, 4, 32, 31};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = sumOddFactors(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1;
    } else {
        printf("false\n");
        return 0;
    }
}