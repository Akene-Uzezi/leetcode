fn first_palindrome(words: Vec<&str>) {
    for word in &words {
        println!("{}", word)
    }
}

fn main() {
    let words = vec!["james", "john"];
    first_palindrome(words);
}
