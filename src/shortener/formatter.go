package shortener

import (
	"reku-technical-test/src/entity"
)

type ShortFormatter struct {
	ID         string `json:"id" gorm:"unique"`
	TargetURL  string `json:"target_url"`
	ExpiryDate int    `json:"expiry_date"`
	Clicks     int    `json:"clicks"`
}

type FetchFormatter struct {
	ID        string `json:"id" gorm:"unique"`
	TargetURL string `json:"target_url"`
}

func FormatShortened(short entity.Shortener) ShortFormatter {
	formatShortened := ShortFormatter{
		ID:         short.ID,
		TargetURL:  short.TargetURL,
		ExpiryDate: short.ExpiryDate,
		Clicks:     short.Clicks,
	}

	return formatShortened
}

func FormatShorteneds(shorts []entity.Shortener) []ShortFormatter {
	shortsFormatter := []ShortFormatter{}

	for _, shorted := range shorts {
		shortFormatter := FormatShortened(shorted)
		shortsFormatter = append(shortsFormatter, shortFormatter)
	}

	return shortsFormatter
}

func FormatShort(shortened, targetUrl string) FetchFormatter {
	formatShortened := FetchFormatter{
		ID:        shortened,
		TargetURL: targetUrl,
	}

	return formatShortened
}
