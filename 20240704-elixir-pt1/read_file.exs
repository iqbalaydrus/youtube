defmodule CityTemperature do
  defstruct city: "", temp_sum: 0.0, temp_count: 0, temp_max: 0.0, temp_min: 0.0
end

defmodule Main do
  def main do
    # Read and process the file stream
    city_temperature_map =
      File.stream!("measurements.txt")
      |> Stream.map(&String.trim/1)
      |> Stream.map(&String.split(&1, ";"))
      |> Stream.map(&List.to_tuple/1)
      |> Stream.map(fn {city, temperature} -> {city, String.to_float(temperature)} end)
      |> Enum.reduce(%{}, fn {city, temperature}, acc ->
        Map.update(
          acc,
          city,
          %CityTemperature{
            city: city,
            temp_sum: temperature,
            temp_count: 0,
            temp_max: temperature,
            temp_min: temperature
          },
          fn existing ->
            %CityTemperature{
              existing
              | temp_sum: existing.temp_sum + temperature,
                temp_count: existing.temp_count + 1,
                temp_max: max(existing.temp_max, temperature),
                temp_min: min(existing.temp_min, temperature)
            }
          end
        )
      end)

    samples = Enum.take(city_temperature_map, 5)
    IO.inspect(samples)
  end
end

Main.main()
