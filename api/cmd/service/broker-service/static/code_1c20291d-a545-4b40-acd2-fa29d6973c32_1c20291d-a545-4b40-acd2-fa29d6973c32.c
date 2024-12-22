#include <stdio.h>
#include <string.h>

// User function logic goes here
int countWords(char str[]) {
    int count = 0, i = 0;
    while (str[i] != '\0') {
        while (str[i] == ' ' && str[i] != '\0') i++;
        if (str[i] != '\0') {
            count++;
            while (str[i] != ' ' && str[i] != '\0') i++;
        }
    }
    return count;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello world", "abc def ghi", "C programming language", "No words"};
    int expected[] = {2, 3, 3, 2};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = countWords(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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