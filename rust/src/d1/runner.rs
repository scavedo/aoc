use std::fs;

pub fn run(filepath: &str) {
    let p1 = part_one(filepath);
    let p2 = part_two(filepath);

    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);
}

fn part_one(filepath: &str) -> i32 {
    let file = fs::read_to_string(filepath)
        .expect("Should have read the file");
    let (mut left, mut right): (Vec<i32>, Vec<i32>) = file
        .lines()
        .map(|line| {
            let x = line.split_whitespace().collect::<Vec<&str>>();
            let left = x[0].parse::<i32>().unwrap();
            let right = x[1].parse::<i32>().unwrap();
            (left, right)
        })
        .collect::<Vec<(i32, i32)>>()
        .into_iter()
        .unzip();

    left.sort();
    right.sort();

    return left.iter()
        .zip(right.iter())
        .map(|(l, r)| (l - r).abs())
        .sum::<i32>();
}

fn part_two(filepath: &str) -> i32 {
    let file = fs::read_to_string(filepath)
        .expect("Should have read the file");
    let (left, right): (Vec<i32>, Vec<i32>) = file
        .lines()
        .map(|line| {
            let x = line.split_whitespace().collect::<Vec<&str>>();
            let left = x[0].parse::<i32>().unwrap();
            let right = x[1].parse::<i32>().unwrap();
            (left, right)
        })
        .collect::<Vec<(i32, i32)>>()
        .into_iter()
        .unzip();

    return left.iter()
        .map(|l| {
            right.iter().filter(|r| *r == l).count() as i32 * l
        })
        .sum::<i32>();
}
