use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut count_part_one: i32 = 0;
    let mut count_part_two: i32 = 0;
    // File hosts must exist in current path before this produces output
    if let Ok(lines) = read_lines("./input.txt") {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(ip) = line {
                let values = parse_line(&ip);
                part_one(values, &mut count_part_one);
                part_two(values, &mut count_part_two);
            }
        }
        print!("{}\n", count_part_one);
        print!("{}", count_part_two);
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn parse_line(line: &str) -> ([i32; 4]) {
    let split_comma: Vec<&str> = line.split(",").collect();
    let split_text_range_one: Vec<&str> = split_comma[0].split("-").collect();
    let split_text_range_two: Vec<&str> = split_comma[1].split("-").collect();
    let mut ints: [i32; 4] = [0; 4];
    ints[0] = split_text_range_one[0].parse().unwrap();
    ints[1] = split_text_range_one[1].parse().unwrap();
    ints[2] = split_text_range_two[0].parse().unwrap();
    ints[3] = split_text_range_two[1].parse().unwrap();
    return ints;
}

fn part_one(values: [i32; 4], count: &mut i32) -> () {
    let x1: i32 = values[0];
    let x2: i32 = values[1];
    let y1: i32 = values[2];
    let y2: i32 = values[3];
    if y1 <= x1 && y2 >= x2 || x1 <= y1 && x2 >= y2 {
        *count += 1
    }
}

fn part_two(values: [i32; 4], count: &mut i32) -> () {
    let x1: i32 = values[0];
    let x2: i32 = values[1];
    let y1: i32 = values[2];
    let y2: i32 = values[3];
    if y1 <= x1 && y2 >= x2 || x1 <= y1 && x2 >= y2 || y1 <= x1 && x1 <= y2 || y1 <= x2 && x2 <= y2
    {
        *count += 1
    }
}
