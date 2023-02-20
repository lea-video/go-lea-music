package main

import (
	"fmt"
	"github.com/lea-video/go-lea-music/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const PORT = 8081

func prepServer() *fiber.App {
	app := fiber.New()
	// allow for gzip, brotli or deflate transport compression
	app.Use(compress.New())
	// handle CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))
	return app
}

func addRoutes(app *fiber.App, db routes.GenericDB) {
	wrapper := routes.RouteWrapper{DB: db}

	app.Get("/api/v1/artist", wrapper.GetArtist)
	app.Post("/api/v1/artist", wrapper.PostArtist)
	app.Patch("/api/v1/artist", wrapper.UpdateArtist)
	app.Get("/api/v1/artist/:id", wrapper.GetArtistID)
	app.Delete("/api/v1/artist/:id", wrapper.DelArtistID)

	app.Get("/api/v1/song", wrapper.GetSong)
	app.Post("/api/v1/song", wrapper.PostSong)
	app.Patch("/api/v1/song", wrapper.UpdateSong)
	app.Get("/api/v1/song/:id", wrapper.GetSongID)
	app.Delete("/api/v1/song/:id", wrapper.DelSongID)

	app.Get("/api/v1/variation", wrapper.GetVariation)
	app.Post("/api/v1/variation", wrapper.PostVariation)
	app.Patch("/api/v1/variation", wrapper.UpdateVariation)
	app.Get("/api/v1/variation/:id", wrapper.GetVariationID)
	app.Delete("/api/v1/variation/:id", wrapper.DelVariationID)

	// app.Get("/api/v1/tag", wrapper.GetTag)
	// app.Post("/api/v1/tag", wrapper.PostTag)
	// app.Patch("/api/v1/tag", wrapper.UpdateTag)
	// app.Get("/api/v1/tag/:id", wrapper.GetTagID)
	// app.Delete("/api/v1/tag/:id", wrapper.DelTagID)
}

func startServer(app *fiber.App) {
	// listen on port 8081
	listenAddr := fmt.Sprintf(":%d", PORT)
	log.Fatal(app.Listen(listenAddr))
}
