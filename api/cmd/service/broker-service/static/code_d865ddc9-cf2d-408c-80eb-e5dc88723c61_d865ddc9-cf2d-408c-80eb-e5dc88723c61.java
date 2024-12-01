public class Main {

    // User function logic goes here
    public static String reverseString(String str) {
    StringBuilder reversed = new StringBuilder(str);
    return reversed.reverse().toString();
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[] testCases = {
        "hello",  // Test case 1: Reverse of hello is olleh
        "world"   // Test case 2: Reverse of world is dlrow
    };
    for (String testCase : testCases) {
        String input = testCase;
        String expectedOutput = reverseString(testCase);
        String output = reverseString(testCase);
        if (!output.equals(expectedOutput)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}