# frozen_string_literal: true

(1..30).each do |v|
  m = []
  m.append('fizz') if (v % 3).zero?
  m.append('buzz') if (v % 5).zero?
  puts m.empty? ? v.to_s : m.join('')
end
