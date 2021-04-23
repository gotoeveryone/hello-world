word = String.trim(IO.gets "Please enter the singular word:")

if String.match?(word, ~r/[s|sh|ch|o|x]$/) do
  IO.puts word <> "es"
  System.halt(0)
end

matched = Regex.run(~r/(.*)(fe?)$/, word)
if matched do
  IO.puts Enum.at(matched, 1) <> "ves"
  System.halt(0)
end

matched = Regex.run(~r/(.*[^aiueo])y$/, word)
if matched do
  IO.puts Enum.at(matched, 1) <> "ies"
  System.halt(0)
end

IO.puts word <> "s"
