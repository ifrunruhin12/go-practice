### **Task 1: Ping → Pong**

* Create **two channels**: `pingChan` and `pongChan`
* One goroutine sends `"ping"` into `ping`
* Another goroutine:

    * receives from `ping`
    * prints it
    * sends `"pong"` into `pong`
* `main` receives `"pong"` and prints it

**Concepts:**

✔ channel → channel communication

✔ blocking behavior