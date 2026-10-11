use rand;
use std::{collections::HashMap, io};

fn welcome_message() {
    println!("Welcome to the Number Guessing Game!");
    println!("I'm thinking of a number between 1 and 100.");
    println!("You have multiple chances to guess the correct number.");
    println!();
}

fn return_difficulty(level: i32) -> i32 {
    let difficulty_number = HashMap::from([(1, 10), (2, 5), (3, 3)]);
    return difficulty_number.get(&level).copied().unwrap_or_default();
}

fn return_difficulty_string(level: i32) -> String {
    let difficulty_string = HashMap::from([(1, "Easy"), (2, "Medium"), (3, "Hard")]);
    difficulty_string
        .get(&level)
        .copied()
        .map(String::from)
        .unwrap_or_default()
}

fn select_difficult() -> (i32, i32) {
    println!("Please select the difficulty level:");
    println!("1. Easy (10 chances)");
    println!("2. Medium (5 chances)");
    println!("3. Hard (3 chances)");
    println!();

    let mut user_input = String::new();

    println!("Enter your choice: ");
    io::stdin()
        .read_line(&mut user_input)
        .expect("Failed to read input");

    let selection: i32 = user_input.trim().parse().expect("Input is not an Integer");

    println!("You selected {}", selection);

    return (selection, return_difficulty(selection));
}

fn game_loop(level: i32, chances: i32) {
    // random number between 1 and 100
    let current_rng = rand::random_range(1..=100);
    let level_string = return_difficulty_string(level);

    println!(
        "Great! You have selected the {} difficulty level.",
        level_string
    );
    println!("Let's start the game!");

    let mut guess_correct_flag: bool = false;

    for i in 0..chances {
        let mut user_input = String::new();

        println!("Enter your guess: ");
        io::stdin()
            .read_line(&mut user_input)
            .expect("Failed to read input");

        let guess: i32 = user_input.trim().parse().expect("Input is not an Integer");

        if current_rng == guess {
            println!(
                "Congratulations! You guessed the correct number in {} attempts",
                i + 1
            );
            guess_correct_flag = true;
            break;
        } else if current_rng < guess {
            println!("Incorrect! The number is less than {}", guess);
        } else {
            println!("Incorrect! The number is greater than {}", guess);
        }
    }

    if !guess_correct_flag {
        println!("You ran out of chances");
    }
}

fn main() {
    welcome_message();
    let (level, chances): (i32, i32) = select_difficult();
    game_loop(level, chances);
}
