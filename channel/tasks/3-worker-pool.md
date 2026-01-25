### **Task 3: Worker Army**

* Create **1 channel**
* Spawn **5 worker goroutines**
* Each worker:

    * reads a number from the channel
    * prints `workerID processed X`
* `main` sends numbers `1..20` into the channel

**Rules:**

* All workers share the same channel
* No race conditions
* Program must exit cleanly

**Concepts:**

✔ multiple goroutines → one channel

✔ fan-out pattern