seq 30 | xargs -I{} sh -c 'if [ `expr {} % 15` -eq 0 ]; then echo fizzbuzz; elif [ `expr {} % 3` -eq 0 ]; then echo fizz; elif [ `expr {} % 5` -eq 0 ]; then echo buzz; else echo {}; fi'
