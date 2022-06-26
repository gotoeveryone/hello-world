use std::env;
use std::process;

mod commands;
use commands::{
  fizzbuzz,
  hello_world,
  id_odd,
  singlar_to_prural,
};

fn main() {
  let args: Vec<String> = env::args().collect();
  if args.len() == 1 {
    hello_world();
    process::exit(0)
  }
  match args.get(1).unwrap().as_str() {
    "fizzbuzz" => fizzbuzz(),
    "is_odd" => id_odd(),
    "singlar_to_prural" => singlar_to_prural(),
    v => {
      println!("Invalid command: {}", v);
      process::exit(1);
    }
  }
}
