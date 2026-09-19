impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut groups:HashMap<String, Vec<i32>> =  HashMap::new();
        let mut result:Vec<Vec<String>> = Vec::new();

       for (i, string) in strs.iter().enumerate(){
         let sorted_string = Self::sort_string(string);
         groups.entry(sorted_string).or_insert_with(Vec::new).push(i as i32);
       }

       for (k,v) in &groups{
        let mut temp:Vec<String> = Vec::new();
        for i in v{
           temp.push(strs[*i as usize].clone());
        }
        result.push(temp);
       }
        
    return result;

    }
    fn sort_string(s: &str) -> String {
        let mut chars: Vec<char> = s.chars().collect();
        chars.sort_unstable();
        chars.into_iter().collect()
}
}
