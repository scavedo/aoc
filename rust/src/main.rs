mod days;

use crate::days::{
    d1,
    d2
};
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let day = &args[1];
    let test = args.get(2).map_or(false, |x| x == "test");
    let filepath = format!("src/days/{}/{}.txt", day, if test { "test" } else { "input" });
    let file = fs::read_to_string(filepath)
        .expect("Should have read the file");

    match day.trim() {
        "d1" => d1::run(&file),
        "d2" => d2::run(&file),
        _ => panic!("Unknown entry"),
    }
}
