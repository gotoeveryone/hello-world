<?php

function isLeapYear(int $year) {
    if ($year % 4 != 0) {
        return false;
    }
    if ($year % 100 == 0) {
        return $year % 400 == 0;
    }
    return true;
}

print('Please enter the year' . PHP_EOL);
$year = trim(fgets(STDIN));

if (!is_numeric($year)) {
    print("'{$year}' is not integer" . PHP_EOL);
    exit(1);
}

if (isLeapYear((int) $year)) {
    print("'{$year}' is leap year" . PHP_EOL);
} else {
    print("'{$year}' is not leap year" . PHP_EOL);
}
