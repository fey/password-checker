package code;

public class App {
    public static void main(String[] args) {
        System.out.println("== Генерация паролей ==");
        System.out.println("буквы и цифры:    " + PasswordGenerator.generatePassword(12, 123L));
        System.out.println("со спецсимволами: " + PasswordGenerator.generatePassword(16, 7L, true, true, true));

        System.out.println();
        System.out.println("== Проверка надёжности ==");
        System.out.println("abc        -> " + PasswordGenerator.checkPassword("abc"));
        System.out.println("abcdef1234 -> " + PasswordGenerator.checkPassword("abcdef1234"));
        System.out.println("Abcdef123! -> " + PasswordGenerator.checkPassword("Abcdef123!"));
    }
}
