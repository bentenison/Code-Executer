# User function logic goes here
class QueueUsingStacks:
    def __init__(self):
        self.stack1 = []
        self.stack2 = []
    def enqueue(self, x):
        self.stack1.append(x)
    def dequeue(self):
        if not self.stack2:
            while self.stack1:
                self.stack2.append(self.stack1.pop())
        return self.stack2.pop()

# Test cases for the function
if __name__ == '__main__':
    queue = QueueUsingStacks()
    queue.enqueue(1)
    queue.enqueue(2)
    result = queue.dequeue()
    print(result == 1)  # Should print True