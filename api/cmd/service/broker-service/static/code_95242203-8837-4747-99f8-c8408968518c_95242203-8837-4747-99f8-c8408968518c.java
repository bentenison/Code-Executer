public class Main {

    // User function logic goes here
    public void push(int x) {
    queue1.add(x);
}

public int pop() {
    while (queue1.size() > 1) {
        queue2.add(queue1.poll());
    }
    int popped = queue1.poll();
    Queue<Integer> temp = queue1;
    queue1 = queue2;
    queue2 = temp;
    return popped;
}

public boolean isEmpty() {
    return queue1.isEmpty();
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {1, 2},  // Test case 1: Push 1, 2, pop should return 2
        {3, 4, 5} // Test case 2: Push 3, 4, 5, pop should return 5
    };
    for (int[] testCase : testCases) {
        {{ .ClassName }} stack = new {{ .ClassName }}();
        stack.push(testCase[0]);
        stack.push(testCase[1]);
        int expectedOutput = testCase[2];
        int output = stack.pop();
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}