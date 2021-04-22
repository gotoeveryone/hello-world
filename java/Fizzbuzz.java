import java.util.stream.IntStream;

public class Fizzbuzz {
    public static void main(String[] args) {
        IntStream.range(1, 31).forEach(v -> {
            if (v % 15 == 0) {
                System.out.println("fizzbuzz");
            } else if (v % 3 == 0) {
                System.out.println("fizz");
            } else if (v % 5 == 0) {
                System.out.println("buzz");
            } else {
                System.out.println(v);
            }
        });
    }
}
