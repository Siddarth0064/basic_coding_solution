import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class SubsetPartition {

    public static List<Integer> subsetA(List<Integer> arr) {
        // Sort the array in descending order to make it easier to select elements for A
        Collections.sort(arr, Collections.reverseOrder());

        List<Integer> subsetA = new ArrayList<>();
        int sumA = 0;
        int sumB = 0;

        // Iterate through the sorted array to construct A and calculate sums
        for (Integer num : arr) {
            if (sumA <= sumB) {
                subsetA.add(num);
                sumA += num;
            } else {
                sumB += num;
            }
        }

        return subsetA;
    }

    public static void main(String[] args) {
        List<Integer> input = Arrays.asList(5, 3, 2, 4, 1, 2);
        List<Integer> result = subsetA(input);

        System.out.println("Output: " + result);
    }
}
