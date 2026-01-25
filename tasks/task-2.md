### **Task 2: One Goroutine, Two Channels**

* One goroutine generates numbers `1..10`
* If number is **even**, send to `evenChan`
* If **odd**, send to `oddChan`
* Two separate goroutines:

    * one prints evens
    * one prints odds

**Rules:**

* No shared variables
* No sleep hacks

**Concepts:**

✔ one goroutine → multiple channels