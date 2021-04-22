fun main() {
    for (v in 1..30) {
        if (v % 15 == 0) {
            println("fizzbuzz")
        } else if (v % 3 == 0) {
            println("fizz")
        } else if (v % 5 == 0) {
            println("buzz")
        } else {
            println(v)
        }
    }
}
