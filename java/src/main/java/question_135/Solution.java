package question_135;

class Solution {
    public int candy(int[] ratings) {
       int ret = 0;

       int[] candies = new int[ratings.length];
       for (int i = 0; i < candies.length; i++) {
           if (i > 0 && ratings[i] > ratings[i - 1]) {
               candies[i] = candies[i - 1] + 1;
               ret += candies[i];
               continue;
           }

           candies[i] = 1;
           ret++;
           continue;
       }
       for (int i = candies.length - 2; i >= 0; i--) {
           if (ratings[i] > ratings[i + 1] && candies[i] <= candies[i + 1]) {
               ret += candies[i + 1] + 1 - candies[i];
               candies[i] = candies[i + 1] + 1;
           }
       }

       return ret;
   }
}