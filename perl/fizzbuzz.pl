for ($i = 1; $i <= 30; $i++) {
  my @m = ();
  if ($i % 3 == 0) {
    push(@m, 'fizz');
  }
  if ($i % 5 == 0) {
    push(@m, 'buzz');
  }
  print((@m ? join('', @m) : $i) . "\n");
}
