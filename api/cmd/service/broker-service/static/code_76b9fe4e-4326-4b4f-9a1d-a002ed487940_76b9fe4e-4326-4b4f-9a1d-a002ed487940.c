#include <stdio.h>
#include <string.h>

// User function logic goes here
void swapStrings(char str1[], char str2[]) {
    char temp[100];
    strcpy(temp, str1);
    strcpy(str1, str2);
    strcpy(str2, temp);
}

// Test cases for the function
int main() {
    char test_cases[2][2][50] = {
        {"hello", "world"},
        {"abc", "xyz"}
    };
    for (int i = 0; i < 2; i++) {
        swapStrings(test_cases[i][0], test_cases[i][1]);
        printf("test %d: str1: %s, str2: %s\n", i, test_cases[i][0], test_cases[i][1]);
    }
    return 0;
}