impl Solution {
    pub fn count_components(n: i32, edges: Vec<Vec<i32>>) -> i32 {
     
    let mut graph:HashMap<i32, Vec<i32>> = HashMap::new();
    let mut visited: HashSet<i32> = HashSet::new();
    let mut count = 0;
    // build graph
    for pair in &edges {
        graph.entry(pair[1]).or_insert_with(Vec::new).push(pair[0]);
        graph.entry(pair[0]).or_insert_with(Vec::new).push(pair[1]);
    
    }
    
    for node in 0..n{
        if !visited.contains(&node){
            count+=1;
            Self::dfs(&mut visited, &graph, node)
        }
    }

     
    
    return count

    }

    pub fn dfs(visited: &mut HashSet<i32>, graph: &HashMap<i32, Vec<i32>>, node: i32){
        let mut stack: Vec<i32> = Vec::new();

        // seed the stack with the first value

        stack.push(node);
        visited.insert(node);
        

        while stack.len() > 0 {
            let node = stack.pop().unwrap();
            

            if let Some(neighbors) = graph.get(&node){
                for next_node in &graph[&node] {
                    if !visited.contains(&next_node){
                        visited.insert(*next_node);
                        stack.push(*next_node);
                    }
            }
            }
            

        }
         
    }
   
}
