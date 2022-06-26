use regex::Regex;

fn main() {
	println!("Please enter the singular word:")

	var w string
	fmt.Scan(&w)

	m, e := regexp.MatchString("[s|sh|ch|o|x]$", w)
	if e != nil {
		std::process::exit(1)
	}
	if m {
		println!(w + "es")
		os.Exit(0)
	}

	r, e := regexp.Compile("(.*)(fe?)$")
	if e != nil {
		std::process::exit(1)
	}
	f := r.FindStringSubmatch(w)
	if f != nil {
		println!(f[1] + "ves")
		os.Exit(0)
	}

	r2, e := regexp.Compile("(.*[^aiueo])y$")
	if e != nil {
		std::process::exit(1)
	}
	f2 := r2.FindStringSubmatch(w)
	if f2 != nil {
		println!(f2[1] + "ies")
		os.Exit(0)
	}

	println!(w + "s")
}
