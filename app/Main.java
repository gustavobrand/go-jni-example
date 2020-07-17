package app;

public class Main {
    public native static String nativeFetch();

    public static void main(String[] args) {
        System.loadLibrary("fetch");
        System.err.println(">>: " + nativeFetch());
    }
}
