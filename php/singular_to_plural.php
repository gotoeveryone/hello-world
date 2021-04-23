<?php
print('Please enter the singular word:' . PHP_EOL);
$word = trim(fgets(STDIN));

if (preg_match('/[s|sh|ch|o|x]$/', $word)) {
    print("{$word}es" . PHP_EOL);
    exit();
}
if (preg_match('/(.*)(fe?)$/', $word, $m)) {
    print("{$m[1]}ves" . PHP_EOL);
    exit();
}
if (preg_match('/(.*[^aiueo])y$/', $word, $m)) {
    print("{$m[1]}ies" . PHP_EOL);
    exit();
}

print("{$word}s" . PHP_EOL);
