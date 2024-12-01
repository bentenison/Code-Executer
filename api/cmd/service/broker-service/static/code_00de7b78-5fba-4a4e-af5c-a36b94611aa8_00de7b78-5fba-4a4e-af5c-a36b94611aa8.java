public class Main {

    // User function logic goes here
    public static int findLength(String str) {
    int length = 0;
    for (char c : str.toCharArray()) {
        length++;
    }
    return length;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[] testCases = {
        "Hello",  // Test case 1: Length should be 5
        "World"   // Test case 2: Length should be 5
    };
    for (String testCase : testCases) {
        int expectedOutput = testCase.length();
        int output = findLength(testCase);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}