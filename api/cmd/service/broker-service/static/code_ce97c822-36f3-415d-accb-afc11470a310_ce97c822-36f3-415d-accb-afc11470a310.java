public class Main {

    // User function logic goes here
    public static int countWords(String str) {
    if (str == null || str.isEmpty()) return 0;
    String[] words = str.trim().split("\s+");
    return words.length;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[] testCases = {
        "Hello World!",  // Test case 1: Two words
        "Java is fun"    // Test case 2: Three words
    };
    for (String testCase : testCases) {
        String input = testCase;
        int expectedOutput = (input.split("\s+")).length;
        int output = countWords(testCase);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}