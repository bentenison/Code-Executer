#include <stdio.h>
#include <math.h>

// User function logic goes here
int isArmstrong(int num) {
    int sum = 0, temp = num, digits = 0;
    while (temp != 0) {
        digits++;
        temp /= 10;
    }
    temp = num;
    while (temp != 0) {
        int rem = temp % 10;
        sum += pow(rem, digits);
        temp /= 10;
    }
    return sum == num;
}

// Test cases for the function
int main() {
    int test_cases[] = {153, 370, 371, 123};
    int expected[] = {1, 1, 1, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = isArmstrong(test_cases[i]);
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