### **Task 5: Graceful Shutdown**

* One goroutine produces numbers every second
* Another goroutine consumes & prints
* After **5 seconds**, stop everything cleanly

**Rules:**

* No global variables
* No `os.Exit`

**Hint:**
There is a *very Go-ish* way to do this

**Concepts:**

✔ done channel / signal channel

✔ graceful goroutine shutdown