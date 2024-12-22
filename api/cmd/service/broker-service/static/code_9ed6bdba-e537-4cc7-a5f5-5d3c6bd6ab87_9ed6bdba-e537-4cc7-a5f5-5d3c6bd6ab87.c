#include <stdio.h>
#include <string.h>

// User function logic goes here
void removeCharacters(char str1[], char str2[]) {
    int len1 = strlen(str1);
    int len2 = strlen(str2);
    for (int i = 0; i < len1; i++) {
        for (int j = 0; j < len2; j++) {
            if (str1[i] == str2[j]) {
                for (int k = i; k < len1; k++) {
                    str1[k] = str1[k + 1];
                }
                len1--;
                i--;
                break;
            }
        }
    }
}

// Test cases for the function
int main() {
    char test_cases[2][2][50] = {
        {"abcdef", "ace"},
        {"hello world", "lo"}
    };
    for (int i = 0; i < 2; i++) {
        removeCharacters(test_cases[i][0], test_cases[i][1]);
        printf("Modified string: %s\n", test_cases[i][0]);
    }
    return 0;
}