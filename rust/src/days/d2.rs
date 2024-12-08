
pub fn run(file: &str) {
    let reports = file
        .lines()
        .map(|line| {
            return line.split_whitespace()
            .map(|x| x.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        })
        .collect::<Vec<Vec<i32>>>();
    
    let p1 = part_one(&reports);
    let p2 = part_two(&reports);

    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);
}

fn part_one(reports: &[Vec<i32>]) -> usize {
    reports.iter()
        .map(|report| {
            has_no_problems(report)
        })
        .filter(|x| *x)
        .count()
}

fn part_two(reports: &[Vec<i32>]) -> usize {
    reports.iter()
        .map(|report| {
            if has_no_problems(report) {
                return true;
            }
            
            let size = report.len();
            for i in 0..size {
                let mut new_report = report.clone();
                new_report.remove(i);
                if has_no_problems(&new_report) {
                    return true;
                }
            }
            
            false
        })
        .filter(|x| *x)
        .count()
}

fn has_no_problems(report: &[i32]) -> bool {
    return report
        .windows(3)
        .map(|x| {
            let y = x[0] - x[1];
            let z = x[1] - x[2];
            if y.abs() < 1 || y.abs() > 3 || z.abs() < 1 || z.abs() > 3 {
                return false;
            }
            if y * z < 0 {
                return false;
            }
            true
        })
        .all(|x| x);
}
