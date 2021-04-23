puts 'Please enter the singular word:'
word = gets.chomp

word.match('[s|sh|ch|o|x]$') do |m|
  puts "#{word}es"
  exit
end

word.match('(.*)(fe?)') do |m|
  puts "#{m[1]}ves"
  exit
end

word.match('(.*[^aiueo])y$') do |m|
  puts "#{m[1]}ies"
  exit
end

puts "#{word}s"
