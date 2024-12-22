#include <stdio.h>
#include <string.h>

// User function logic goes here
void sortStrings(char arr[][100], int n) {
    for (int i = 0; i < n - 1; i++) {
        for (int j = i + 1; j < n; j++) {
            if (strcmp(arr[i], arr[j]) > 0) {
                char temp[100];
                strcpy(temp, arr[i]);
                strcpy(arr[i], arr[j]);
                strcpy(arr[j], temp);
            }
        }
    }
}

// Test cases for the function
int main() {
    char test_cases[2][3][100] = {
        {"apple", "banana", "cherry"},
        {"dog", "cat", "elephant"}
    };
    int n[] = {3, 3};
    for (int i = 0; i < 2; i++) {
        sortStrings(test_cases[i], n[i]);
        for (int j = 0; j < n[i]; j++) {
            printf("%s\n", test_cases[i][j]);
        }
    }
    return 0;
}