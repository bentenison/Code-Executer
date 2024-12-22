#include <stdio.h>
#include <string.h>

// User function logic goes here
char* firstLetterOfEachWord(char str[]) {
    static char result[100];
    int j = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        if (i == 0 || str[i-1] == ' ') {
            result[j++] = str[i];
        }
    }
    result[j] = '\0';
    return result;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"hello world golang", "good morning", "openai rocks"};
    char expected[][50] = {"hwg", "gm", "or"};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        char* result = firstLetterOfEachWord(test_cases[i]);
        if (strcmp(result, expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], result);
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