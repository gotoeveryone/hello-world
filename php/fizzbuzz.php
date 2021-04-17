<?php

for ($i = 1; $i <= 30; $i++) {
  $m = [];
  if ($i % 3 === 0) {
      $m[] = 'fizz';
  }
  if ($i % 5 === 0) {
      $m[] = 'buzz';
  }
  print(($m ? implode('', $m) : $i) . PHP_EOL);
}
