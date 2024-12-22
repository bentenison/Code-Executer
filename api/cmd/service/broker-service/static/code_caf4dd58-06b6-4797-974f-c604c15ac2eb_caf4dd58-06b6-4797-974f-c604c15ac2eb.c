#include <stdio.h>

// User function logic goes here
float calculateSimpleInterest(float principal, float rate, float time) {
    return (principal * rate * time) / 100;
}

// Test cases for the function
int main() {
    int passed = 1;
    float principal = 1000, rate = 5, time = 2;
    float expected = 100.0;
    if (calculateSimpleInterest(principal, rate, time) != expected) {
        passed = 0;
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}