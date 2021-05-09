sub isLeapYear {
  my ($year) = @_;
  if ($year % 4 != 0) {
    return 0;
  }
  if ($year % 100 == 0) {
    return $year % 400 == 0 ? 1 : 0;
  }
  return 1;
}

print("Please enter the year\n");
my $year = <STDIN>;
chomp($year);

if ($year !~ /(\d+)$/) {
  print("'${year}' is not integer\n");
  exit(1);
}

if (&isLeapYear(int($year)) == 1) {
  print("'${year}' is leap year\n");
} else {
  print("'${year}' is not leap year\n");
}
