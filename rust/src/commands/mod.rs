use regex::Regex;

pub fn hello_world() {
    println!("hello world")
}

pub fn fizzbuzz() {
  for i in 1..31 {
    if i % 15 == 0 {
      println!("fizzbuzz")
    } else if i % 3 == 0 {
      println!("fizz")
    } else if i % 5 == 0 {
      println!("buzz")
    } else {
      println!("{}", i)
    }
  }
}

pub fn id_odd() {
  println!("Please enter the number:");
  let mut s: String = String::new();
  std::io::stdin().read_line(&mut s).ok();
  match s.trim().parse::<i32>() {
    Ok(i) => {
      if i == 0 {
        println!("Can't specify zero");
        std::process::exit(1);
      }
      println!("Is odd? {}", i % 2 == 1);
    }
    Err(e) => {
      println!("{}", e);
      std::process::exit(1);
    }
  };
}

pub fn singlar_to_prural() {
	println!("Please enter the singular word:");
  let mut s: String = String::new();
  std::io::stdin().read_line(&mut s).ok();
  let v = s.trim();

	let re = Regex::new("[s|sh|ch|o|x]$").unwrap();
  if re.is_match(v) {
    let cap = re.captures(v).unwrap();
    if cap.len() > 0 {
      println!("{}es", cap[0].to_string());
      std::process::exit(0)
    }
  }

	let re = Regex::new("(.*)(fe?)$").unwrap();
  if re.is_match(v) {
    let cap = re.captures(v).unwrap();
    if cap.len() > 1 {
      println!("{}es", cap[1].to_string());
      std::process::exit(0)
    }
  }

	let re = Regex::new("(.*[^aiueo])y$").unwrap();
  if re.is_match(v) {
    let cap = re.captures(v).unwrap();
    if cap.len() > 1 {
      println!("{}ies", cap[1].to_string());
      std::process::exit(0)
    }
  }

	println!("{}s", v)
}
