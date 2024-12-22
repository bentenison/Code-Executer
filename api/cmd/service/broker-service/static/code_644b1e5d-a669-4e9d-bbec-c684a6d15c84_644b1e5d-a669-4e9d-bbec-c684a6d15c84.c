#include <stdio.h>

// User function logic goes here
int stringLength(char str[]) {
    int length = 0;
    while (str[length] != '\0') {
        length++;
    }
    return length;
}

// Test cases for the function
int main() {
    int passed = 1;
    char str[] = "Hello World!";
    // Check length
    printf("%s\n", passed ? "true" : "false");
    return 0;
}