impl Solution {
    pub fn has_duplicate(nums: Vec<i32>) -> bool {
         let mut count:HashMap<i32,i32> = HashMap::new();

         for v in nums{
            if count.contains_key(&v) && *count.get(&v).unwrap() == 1 {
                return true
            }
            count.entry(v).or_insert(1);
 
         }
    return false

    }
}
