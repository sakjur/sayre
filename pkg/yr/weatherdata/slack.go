package weatherdata

import (
	"fmt"
	"strings"
)

func (weatherdata Weatherdata) ToSlack() string {
	maxLen := 4
	forecasts := make([]string, maxLen)

	for i, forecast := range weatherdata.Forecasts {
		if i >= maxLen {
			break
		}
		forecasts[i] = forecast.ToSlack()
	}

	return fmt.Sprintf(
		"*Weather forecast for %s, %s*\n\n%s\n\n_%s <%s>_",
		weatherdata.Location.Name,
		weatherdata.Location.Country,
		strings.Join(forecasts, "\n"),
		weatherdata.Credit.Text,
		weatherdata.Credit.Url,
	)
}

func (forecast Forecast) ToSlack() string {
	symbol := GetSymbolByNumber(forecast.Symbol.Number)

	return fmt.Sprintf(
		"%s %s %s [%.1fmm %.0f°C] %.0f m/s %s",
		forecast.From,
		symbol.SlackEmoji,
		symbol.Name,
		forecast.Precipitation.Value,
		forecast.Temperature.Value,
		forecast.WindSpeed.MPS,
		forecast.WindDirection.Code,
	)
}
