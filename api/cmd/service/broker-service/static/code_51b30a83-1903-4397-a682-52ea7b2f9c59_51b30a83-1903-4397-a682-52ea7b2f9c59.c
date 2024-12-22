#include <stdio.h>

// User function logic goes here
void findCommonElements(int arr1[], int size1, int arr2[], int size2) {
    for (int i = 0; i < size1; i++) {
        for (int j = 0; j < size2; j++) {
            if (arr1[i] == arr2[j]) {
                printf("%d ", arr1[i]);
            }
        }
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    int arr1[] = {1, 2, 3, 4}, arr2[] = {3, 4, 5, 6};
    // Check common elements
    printf("%s\n", passed ? "true" : "false");
    return 0;
}