package org.study;

import org.study.block1.ArraysTopic;
import org.study.block1.ArraysTopic2;

import java.util.Arrays;
import java.util.Scanner;
import java.util.stream.Collectors;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int length = scanner.nextInt();
        int[] array = new int[length];
        for (int i = 0; i < length; i++) {
            array[i] = scanner.nextInt();
        }
        int element = scanner.nextInt();
        ArraysTopic2 arraysTopic2 = new ArraysTopic2();
        System.out.println(Arrays.stream(arraysTopic2.clearByElement(length, array, element))
                .mapToObj(Integer::toString)
                .collect(Collectors.joining(" ")));
    }
}