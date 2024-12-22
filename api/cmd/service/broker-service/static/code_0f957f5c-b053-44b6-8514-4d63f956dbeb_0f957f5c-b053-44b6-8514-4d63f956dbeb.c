#include <stdio.h>

// User function logic goes here
void reverseWords(char str[]) {
    int start = 0, end = 0;
    while (str[end] != '\0') {
        while (str[end] != ' ' && str[end] != '\0') end++;
        int i = start, j = end - 1;
        while (i < j) {
            char temp = str[i];
            str[i] = str[j];
            str[j] = temp;
            i++, j--;
        }
        if (str[end] != '\0') end++;
        start = end;
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello World", "Coding is fun", "abc def", "SingleWord"};
    char expected[][50] = {"olleH dlroW", "gnidoC si nuf", "cba fed", "droWelgniS"};
    int all_passed = 1;
    for (int i = 0; i < 4; i++) {
        char str[50];
        strcpy(str, test_cases[i]);
        reverseWords(str);
        if (strcmp(str, expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], str);
        }
    }
    if (all_passed) {
        printf("true\n");
    } else {
        printf("false\n");
    }
    return 0;
}