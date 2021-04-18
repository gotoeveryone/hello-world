import json
from os import path

json_dir = path.dirname(path.dirname(path.abspath(__file__)))
with open(json_dir + '/test.json', mode='r', encoding='utf-8') as f:
  json_data = json.loads(f.read())
  for d in json_data:
    print('id: {id}, name: {name}'.format(
      id=d.get('id'),
      name=d.get('name'),
    ))
