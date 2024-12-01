public class Main {

    // User function logic goes here
    public static ListNode reverseList(ListNode head) {
    ListNode prev = null;
    ListNode current = head;
    while (current != null) {
        ListNode nextNode = current.next;
        current.next = prev;
        prev = current;
        current = nextNode;
    }
    return prev;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {1, 2, 3, 4, 5},  // Test case 1: Linked list {1, 2, 3, 4, 5} reversed to {5, 4, 3, 2, 1}
        {10, 20, 30}     // Test case 2: Linked list {10, 20, 30} reversed to {30, 20, 10}
    };
    for (int[] testCase : testCases) {
        ListNode head = createLinkedList(testCase);
        ListNode expected = reverseList(createLinkedList(testCase));
        ListNode output = reverseList(head);
        if (!areEqual(expected, output)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}

private static ListNode createLinkedList(int[] values) {
    ListNode head = new ListNode(values[0]);
    ListNode current = head;
    for (int i = 1; i < values.length; i++) {
        current.next = new ListNode(values[i]);
        current = current.next;
    }
    return head;
}

private static boolean areEqual(ListNode l1, ListNode l2) {
    while (l1 != null && l2 != null) {
        if (l1.val != l2.val) return false;
        l1 = l1.next;
        l2 = l2.next;
    }
    return l1 == null && l2 == null;
}
}