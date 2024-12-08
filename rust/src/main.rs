mod d1;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let day = &args[1];
    let test = args.get(2).map_or(false, |x| x == "test");
    let filepath = format!("src/{}/{}.txt", day, if test { "test" } else { "input" });

    match day.trim() {
        "d1" => d1::runner::run(filepath.as_str()),
        _ => panic!("Unknown entry"),
    }
}
