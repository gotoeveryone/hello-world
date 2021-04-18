<?php

$jsonDir = dirname(__DIR__);
$jsonData = json_decode(file_get_contents("{$jsonDir}/test.json"), true);

foreach ($jsonData as $v) {
  print("id: {$v['id']}, name: {$v['name']}" . PHP_EOL);
}
