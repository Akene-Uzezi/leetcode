// fn first_palindrome(words: Vec<&str>) -> Option<&str> {
//     for word in &words {
//         if is_palindrome(*word) {
//             return Some(*word);
//         }
//     }
//     None
// }
//
// fn is_palindrome(s: &str) -> bool {
//     s.chars().eq(s.chars().rev())
// }
//
// fn main() {
//     let words = vec!["james", "john"];
//     first_palindrome(words);

// }
struct Solution {}
impl Solution {
    pub fn first_palindrome(words: Vec<String>) -> Option<String> {
        for word in words {
            if Self::is_palindrome(&word) {
                return Some(word);
            }
        }
        None
    }

    pub fn is_palindrome(s: &str) -> bool {
        s.chars().eq(s.chars().rev())
    }
}

fn main() {
    let words = vec![String::from("racecar"), String::from("cool")];
    let palindrome = Solution::first_palindrome(words);
    println!("{:?}", palindrome)
}
