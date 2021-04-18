Enum.each(1..30, fn v ->
  cond do
    rem(v, 15) == 0 -> IO.puts("fizzbuzz")
    rem(v, 3) == 0 -> IO.puts("fizz")
    rem(v, 5) == 0 -> IO.puts("buzz")
    true -> IO.puts(v)
  end
end)
