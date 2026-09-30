fn first_palindrome(words: Vec<&str>) -> Option<&str> {
    for word in &words {
        if is_palindrome(*word) {
            return Some(*word);
        }
    }
    None
}

fn is_palindrome(s: &str) -> bool {
    s.chars().eq(s.chars().rev())
}

fn main() {
    let words = vec!["james", "john"];
    first_palindrome(words);
}
