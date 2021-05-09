puts "Please enter the year"
set year [gets stdin]

if {![string is integer -strict $year]} {
  puts "$year is not integer"
  exit 1
}

if {$year % 4 == 0} {
  if {$year % 100 == 0} {
    if {$year % 400 == 0} {
      puts "'$year' is leap year"
    } else {
      puts "'$year' is not leap year"
    }
  } else {
    puts "'$year' is leap year"
  }
} else {
  puts "'$year' is not leap year"
}
