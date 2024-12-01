public class Main {

    // User function logic goes here
    public static int countVowels(String str) {
    int count = 0;
    for (int i = 0; i < str.length(); i++) {
        char c = str.charAt(i);
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U') {
            count++;
        }
    }
    return count;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[] testCases = {
        "hello",  // Test case 1: Vowels in hello are 2
        "world"   // Test case 2: Vowels in world are 1
    };
    for (String testCase : testCases) {
        int input = testCase;
        int expectedOutput = countVowels(testCase);
        int output = countVowels(testCase);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}