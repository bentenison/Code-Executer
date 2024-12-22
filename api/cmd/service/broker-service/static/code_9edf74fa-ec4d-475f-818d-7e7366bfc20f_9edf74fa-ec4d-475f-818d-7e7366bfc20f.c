#include <stdio.h>

// User function logic goes here
int asciiValue(char ch) {
    return (int)ch;
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_case = 'A';
    int expected = 65;
    if (asciiValue(test_case) != expected) {
        passed = 0;
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}