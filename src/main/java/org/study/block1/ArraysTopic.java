package org.study.block1;

public class ArraysTopic {
    /*
    В школе прошёл экзамен по математике. Некоторые ученики списали решения, были замечены и получили 0 баллов.
    Помогите учителю пересортировать оценки учеников.
    Все оценки, равные 0, должны быть перемещены в конец списка, при этом порядок остальных оценок должен остаться неизменным.
     */

    public int[] sort(int length, int[] array) {
        // 6
        // 0 0 6 0 9 8
        // output: 6 9 8 0 0 0
        int posPointer = 0;
        int[] sortedArray = new int[length];
        for (int i = 0; i < length; i++) {
            if (array[i] != 0) {
                sortedArray[posPointer] = array[i];
                posPointer++;
            }
        }
        return sortedArray;
    }
}
