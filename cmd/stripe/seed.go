package main

import (
	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
)

var (
	movies = []model.Movie{
		{
			Title:       "Avatar",
			Description: "Jake Sully lives with his newfound family formed on the planet of Pandora. Once a familiar threat returns to finish what was previously started, Jake must work with Neytiri and the army of the Na'vi race to protect their planet.",
			Price:       7.0,
			CoverURL:    "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/280186877_327463976185980_4928049883994991289_n_480x.progressive.jpg?v=1653311094",
		},
		{
			Title:       "Thor: Love and Thunder",
			Description: "Thor enlists the help of Valkyrie, Korg and ex-girlfriend Jane Foster to fight Gorr the God Butcher, who intends to make the gods extinct.",
			Price:       12.0,
			CoverURL:    "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/thor_480x.progressive.jpg?v=1653571052",
		},
		{
			Title:       "Batman",
			Description: "When a sadistic serial killer begins murdering key political figures in Gotham, Batman is forced to investigate the city's hidden corruption and question his family's involvement.",
			Price:       4.0,
			CoverURL:    "https://cdn.shopify.com/s/files/1/0057/3728/3618/products/the-batman_tgstxmov_480x.progressive.jpg?v=1641930817",
		},
	}
	users = []model.User{
		{
			Email:      "test@gmail.com",
			Subscribed: false,
		},
	}
)
