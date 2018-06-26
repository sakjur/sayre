package weatherdata

type Symbol struct {
	Number     int
	Name       string
	SlackEmoji string
}

var Symbols = []Symbol{
	{Number: 1, Name: "Sunny", SlackEmoji: ":sunny:"},
	{Number: 2, Name: "Fair", SlackEmoji: ":sunny:"},
	{Number: 3, Name: "Partly cloudy", SlackEmoji: ":partly_sunny:"},
	{Number: 4, Name: "Cloudy", SlackEmoji: ":cloud:"},
	{Number: 40, Name: "Light rain showers", SlackEmoji: ":umbrella:"},
	{Number: 5, Name: "Rain showers", SlackEmoji: ":umbrella:"},
	{Number: 41, Name: "Heavy rain showers", SlackEmoji: ":umbrella:"},
	{Number: 24, Name: "Light rain showers and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 6, Name: "Rain showers and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 25, Name: "Heavy rain showers and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 42, Name: "Light sleet showers", SlackEmoji: ":snowflake:"},
	{Number: 7, Name: "Sleet showers", SlackEmoji: ":snowflake:"},
	{Number: 43, Name: "Heavy sleet showers", SlackEmoji: ":snowflake:"},
	{Number: 26, Name: "Light sleet showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 20, Name: "Sleet showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 27, Name: "Heavy sleet showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 44, Name: "Light snow showers", SlackEmoji: ":snowflake:"},
	{Number: 8, Name: "Snow showers", SlackEmoji: ":snowflake:"},
	{Number: 45, Name: "Heavy snow showers", SlackEmoji: ":snowflake:"},
	{Number: 28, Name: "Light snow showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 21, Name: "Snow showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 29, Name: "Heavy snow showers and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 46, Name: "Light rain", SlackEmoji: ":umbrella:"},
	{Number: 9, Name: "Rain", SlackEmoji: ":umbrella:"},
	{Number: 10, Name: "Heavy rain", SlackEmoji: ":umbrella:"},
	{Number: 30, Name: "Light rain and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 22, Name: "Rain and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 11, Name: "Heavy rain and thunder", SlackEmoji: ":zap: :umbrella:"},
	{Number: 47, Name: "Light sleet", SlackEmoji: ":snowflake:"},
	{Number: 12, Name: "Sleet", SlackEmoji: ":snowflake:"},
	{Number: 48, Name: "Heavy sleet", SlackEmoji: ":snowflake:"},
	{Number: 31, Name: "Light sleet and thunder", SlackEmoji: ":snowflake:"},
	{Number: 23, Name: "Sleet and thunder", SlackEmoji: ":snowflake:"},
	{Number: 32, Name: "Heavy sleet and thunder", SlackEmoji: ":snowflake:"},
	{Number: 49, Name: "Light snow", SlackEmoji: ":snowflake:"},
	{Number: 13, Name: "Snow", SlackEmoji: ":snowflake:"},
	{Number: 50, Name: "Heavy snow", SlackEmoji: ":snowflake:"},
	{Number: 33, Name: "Light snow and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 14, Name: "Snow and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 34, Name: "Heavy snow and thunder", SlackEmoji: ":zap: :snowflake:"},
	{Number: 15, Name: "Fog", SlackEmoji: ":foggy:"},
}

func GetSymbolByNumber(number int) Symbol {
	for _, symbol := range Symbols {
		if symbol.Number == number {
			return symbol
		}
	}

	return Symbol{Name: "Unknown", SlackEmoji: ":question:"}
}
