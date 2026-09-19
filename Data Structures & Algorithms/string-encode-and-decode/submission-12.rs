impl Solution {
    pub fn encode(strs: Vec<String>) -> String {
        let sep = "#"; 
        let mut concatinated_string = String::from(""); 
        
        for string in &strs{ 
            let string_len = string.len().to_string();
            concatinated_string.push_str(&(string_len + sep + string)); 
        }

        // println!("{:?}", concatinated_string);

        return concatinated_string;
    }

    pub fn decode(s: String) -> Vec<String> {
        
        let sep = "#"; 
        let mut iter = s.chars().peekable();
        let mut result: Vec<String> = Vec::new();
        let mut new_sequence: bool = false;
        let mut to_be_collected: u32 = 0;
        let mut sequence: String = String::from("");
        let mut number:String = String::from("");
        
        while let Some(c) = iter.next() {
            
            if new_sequence == false && c.is_ascii_digit() {
                 number.push_str(&c.to_string());
                 continue
            }

            to_be_collected = number.parse().unwrap();
            if to_be_collected == 0 {
                result.push(String::from("").clone());
                new_sequence=false;
                continue
            }

            if new_sequence == false && c.to_string() == sep.to_string(){
                new_sequence = true;
                to_be_collected = number.parse().unwrap();
                sequence.clear();
                continue
            }

        
            sequence.push_str(&c.to_string());
            if sequence.len() as u32 == to_be_collected{
                result.push(sequence.clone());
                new_sequence = false;
                number.clear();
            }

        }
    return result;

    }
     
}
