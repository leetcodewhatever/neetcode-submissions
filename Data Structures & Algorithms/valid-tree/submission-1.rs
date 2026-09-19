impl Solution {
    pub fn valid_tree(n: i32, edges: Vec<Vec<i32>>) -> bool {
    if edges.len() < 1 {
        return true;
    }
    let mut graph: HashMap<i32, Vec<i32>> = HashMap::new();
    // 0 - 4
    for pair in &edges {
        graph.entry(pair[0]).or_insert_with(Vec::new).push(pair[1]);
        graph.entry(pair[1]).or_insert_with(Vec::new).push(pair[0]);
    }

    let mut queue: VecDeque<i32> = VecDeque::new();
    let mut visited: HashSet<i32> = HashSet::new();
    let mut parent_son: HashMap<i32, i32> = HashMap::new();
    queue.push_back(0);
    visited.insert(0);

    while queue.len() > 0 {
        let node = queue.pop_front().unwrap(); // node = 1
        for next_node in &graph[&node] {
            // next_node = 0 {1: 0}
            if visited.contains(&next_node) && parent_son.get(&node) != Some(&next_node) {
                return false;
            }
            if !visited.contains(&next_node) {
                queue.push_back(*next_node);
                visited.insert(*next_node);
                parent_son.entry(*next_node).or_insert(node);
            }
        }
    }

    return visited.len() as i32 == n;
    }
}
