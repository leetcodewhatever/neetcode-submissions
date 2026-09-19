impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        // a charachter can be repeated as long as the size of the array.
        let n = nums.len();
        let mut result:Vec<i32> = Vec::new();
        let mut freq:HashMap<i32,i32> = HashMap::new();
        let mut buckets: Vec<Vec<i32>> = vec![Vec::new();n + 1];

        for num in nums{
            *freq.entry(num).or_insert(0) += 1;
            // nums[10,3,3,0] 
            // {10:1, 3:2}
        }
        
        for (k,v) in freq{
           buckets[v as usize].push(k);
        }

      let mut count:i32 = k;
      for bucket in buckets.iter().rev(){
          for top_k in bucket{
            result.push(*top_k);
            count -=1;
            if count == 0 {
                return result;
            }
          } 
      }
    return vec![];
    }
}
