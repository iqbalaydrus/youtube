x = []
File.stream!("measurements.txt")
|> Stream.take(10)
|> Stream.map(&String.trim/1)
|> Stream.with_index
|> Stream.map(fn ({line, index}) -> IO.puts "#{index + 1} #{line}" end)
|> Stream.map(fn ({line, _}) -> ^x = [line | x] end)
|> Stream.run
IO.puts(x)
