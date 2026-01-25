### **Task 6: Pipeline**

Build a **3-stage pipeline**:

1. Generator → sends numbers `1..10`
2. Squarer → squares numbers
3. Printer → prints results

Each stage:

* runs in its own goroutine
* uses channels only

**Concepts:**

✔ pipeline pattern

✔ channel chaining