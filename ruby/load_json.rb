require 'json'

json_path = File.expand_path '../../test.json', __FILE__
json_data = JSON.parse File.read(json_path)
json_data.each do |v|
  puts "id: #{v['id']}, name: #{v['name']}"
end
