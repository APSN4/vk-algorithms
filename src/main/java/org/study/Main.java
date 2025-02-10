package org.study;

import java.util.*;
import java.util.stream.Collectors;

public class Main {
    static int count = 0;

    public static void countCombinations(int target, List<Integer> coins, int index) {
        if (target == 0) {
            count++;
            return;
        }
        if (target < 0 || index >= coins.size()) return;

        int coin = coins.get(index);

        countCombinations(target, coins, index + 1);

        if (target >= coin) {
            countCombinations(target - coin, coins, index);
        }
    }

    public static void printCombinations(int target, List<Integer> coins, int index, List<Integer> path, StringBuilder output) {
        if (target == 0) {
            output.append(path.size()).append(" ")
                    .append(path.stream().map(String::valueOf).collect(Collectors.joining(" ")))
                    .append("\n");
            return;
        }
        if (target < 0 || index >= coins.size()) return;

        int coin = coins.get(index);

        printCombinations(target, coins, index + 1, path, output);

        if (target >= coin) {
            path.add(coin);
            printCombinations(target - coin, coins, index, path, output);
            path.remove(path.size() - 1);
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int target = scanner.nextInt();
        List<Integer> coins = Arrays.asList(1, 5, 10);

        countCombinations(target, coins, 0);
        System.out.println(count);

        StringBuilder output = new StringBuilder();
        printCombinations(target, coins, 0, new ArrayList<>(), output);
        System.out.print(output);
    }
}
