#include <stdio.h>
#include <ctype.h>
#include <string.h>

// User function logic goes here
void removePunctuation(char str[], char result[]) {
    int j = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        if (isalnum(str[i]) || isspace(str[i])) {
            result[j++] = str[i];
        }
    }
    result[j] = '\0';
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello, World!", "This is great.", "#C programming?", "(Python)", "No punctuations here"};
    char expected[][50] = {"Hello World", "This is great", "C programming", "Python", "No punctuations here"};
    int all_passed = 1;
    char result[100];

    for (int i = 0; i < 5; i++) {
        removePunctuation(test_cases[i], result);
        if (strcmp(result, expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], result);
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