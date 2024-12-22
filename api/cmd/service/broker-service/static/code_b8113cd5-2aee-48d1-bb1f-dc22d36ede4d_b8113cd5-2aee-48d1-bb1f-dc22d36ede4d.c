void removeChar(char *s, int i) {
    int len = strlen(s);
    if (i < 0 || i >= len) return;
    for (int j = i; j < len; j++) {
        s[j] = s[j + 1];
    }
}

int main() {
    int all_passed = 1;
    char s1[] = "hello";
    removeChar(s1, 1);
    if (strcmp(s1, "hllo") != 0) all_passed = 0;

    char s2[] = "python";
    removeChar(s2, 3);
    if (strcmp(s2, "pyton") != 0) all_passed = 0;

    char s3[] = "world";
    removeChar(s3, 0);
    if (strcmp(s3, "orld") != 0) all_passed = 0;

    char s4[] = "abcdef";
    removeChar(s4, 5);
    if (strcmp(s4, "abcde") != 0) all_passed = 0;

    if (all_passed) {
        printf("true");
    } else {
        printf("false");
    }
    return 0;
}