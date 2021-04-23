import re
import sys

word = input('Please enter the singular word:').strip()

if re.match(r'[s|sh|ch|o|x]$', word):
    print(word + 'es')
    sys.exit(0)

m = re.match(r'(.*)(fe?)', word)
if m:
    print(m.group(1) + 'ves')
    sys.exit(0)

m = re.match(r'(.*[^aiueo])y$', word)
if m:
    print(m.group(1) + 'ies')
    sys.exit(0)

print(word + 's')
