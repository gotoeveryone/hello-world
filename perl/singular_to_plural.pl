print("Please enter the singular word:\n");
my $word = <STDIN>;
chomp($word);

if ($word =~ /[s|sh|ch|o|x]$/) {
  print("${word}es\n");
  exit();
}
if ($word =~ /(.*)(fe?)$/) {
  print("${1}ves\n");
  exit();
}
if ($word =~ /(.*[^aiueo])y$/) {
  print("${1}ies\n");
  exit();
}

print("${word}s\n");
