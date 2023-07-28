struct Solution;

impl Solution {
    fn remove_non_alphanumeric(s: &str) -> String {
        s.chars()
            .filter(|c| c.is_alphanumeric())
            .collect()
    }
    
    pub fn is_palindrome(s: String) -> bool {
        if s.is_empty() {
            return true;
        }
        let lowercase = s.to_lowercase();
        let s2: String = Self::remove_non_alphanumeric(lowercase.as_str());
        let s_chars: Vec<char> = s2.chars().collect();
        if s_chars.is_empty() {
            return true;
        }
        let mut i = 0;
        let mut j = s_chars.len() - 1;
        while i < j {
            if s_chars[i] != s_chars[j] {
                return false;
            }
            i += 1;
            j -= 1;
        }
        true
    }
}

fn main() {
    let phrase = " ";
    println!("{}", Solution::is_palindrome(phrase.to_string()));
}