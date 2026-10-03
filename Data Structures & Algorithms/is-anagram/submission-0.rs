impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut s_chars: HashMap<char, i32> = HashMap::new();
        let mut t_chars: HashMap<char, i32> = HashMap::new();

        for (s, t) in s.chars().zip(t.chars()){
            *s_chars.entry(s).or_insert(1) += 1;
            *t_chars.entry(t).or_insert(1) += 1;
        }

        s_chars == t_chars
    }
}
