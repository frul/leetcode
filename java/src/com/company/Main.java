package com.company;

public class Main {

    public static void main(String[] args) {
	    int[] nums = new int[args.length - 1];
	    int target = Integer.parseInt(args[args.length - 1]);
	    for (int i = 0; i < args.length - 1; ++i)
        {
            nums[i] = Integer.parseInt(args[i]);
        }

	    TwoSum twoSum = new TwoSum();
	    int[] result = twoSum.twoSum(nums, target);
	    System.out.printf("The result is [ %d, %d ]", result[0], result[1]);
    }
}
