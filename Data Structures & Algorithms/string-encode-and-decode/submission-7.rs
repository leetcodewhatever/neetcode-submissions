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
        let mut count: u32 = 0;
        let mut new_sequence: bool = false;
        let mut to_be_collected: u32 = 0;
        let mut temp_string: String = String::from("");
        let mut number:String = String::from("");
        
        while let Some(c) = iter.next() {
    
            // let next = iter.peek();
           
            if new_sequence == false && c.is_ascii_digit() {
                 number.push_str(&c.to_string());
                 continue
            }

            // println!("digits: {:?}", number);

            if new_sequence == false && c.to_string() == sep.to_string(){
                // println!(" next inside{:?}", next);
                // start a sequence
                // println!("i am here");
                count = 0;
                new_sequence = true;
                
                // extract the full number
                // while ite//r.

                // to_be_collected = c.to_digit(10).unwrap();
                
                to_be_collected = number.parse().unwrap();
                // println!("to be collected: {:?}", to_be_collected);
                
                if to_be_collected == 0 {
                    result.push(String::from("").clone());
                    new_sequence=false;
                    continue
                }
                
                temp_string = String::from("");
            }

              

            // todo: why here it dosn't work
            //  if to_be_collected == 0 {
            //         result.push(String::from("").clone());
            //         new_sequence=false;
            //         continue
            //     }
        
            count+=1;

            if count < 2 {
                continue
            }

        

            // start collecting
            temp_string.push_str(&c.to_string());
            // println!("temp {:?}", temp_string);
            if count == to_be_collected + 1{
                // println!("temp {:?}", temp_string);
                result.push(temp_string.clone());
                new_sequence = false;
                number = String::from("");
            }

        }

        // if result.len() == 0 {
        //     return vec![String::from("")];
        // }
    return result;

    }
     
}
