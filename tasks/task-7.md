### **Task 7: Timeout Worker**

* Worker goroutine does “work” (random delay)
* `main` waits for result
* If result takes **>2 seconds**, timeout and exit

**Rules:**

* No sleep polling
* Use Go primitives only

**Concepts:**

✔ `select`

✔ timeouts

✔ cancellation