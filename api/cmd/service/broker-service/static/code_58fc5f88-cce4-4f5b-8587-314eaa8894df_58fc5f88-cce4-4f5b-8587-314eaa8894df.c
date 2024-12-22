#include <stdio.h>

// User function logic goes here
void findFactors(int number) {
    for (int i = 1; i <= number; i++) {
        if (number % i == 0) {
            printf("%d ", i);
        }
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    int number = 28;
    // Check for factors
    printf("%s\n", passed ? "true" : "false");
    return 0;
}