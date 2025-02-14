package org.study.block1;

import java.util.Arrays;

public class ArraysTopic2 {
    /*
    Учёные провели ряд экспериментов, но из-за неисправности оборудования некоторые результаты оказались ошибочными.
    Ошибочные результаты представлены числом element.
    Ваша задача — удалить все вхождения числа element из массива, сохранив порядок остальных чисел.
     */

    public int[] clearByElement(int length, int[] array, int element) {
        // 7
        // 1 2 -1 4 5 -1 6
        // -1
        // output: 1 2 4 5 6
        int[] clearArray = new int[length];
        int pointer = 0;
        for (int i = 0; i < length; i++) {
            if (array[i] != element) {
                clearArray[pointer] = array[i];
                pointer++;
            }
        }
        return Arrays.copyOfRange(clearArray, 0, pointer);
    }
}
