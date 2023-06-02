package main

import (
	"fmt"
	"log"

	"github.com/lea-video/go-lea-music/db"
	"github.com/lea-video/go-lea-music/routes"

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

func addRoutes(app *fiber.App, db db.GenericDB, fileDB db.GenericFileDB) {
	wrapper := routes.RouteWrapper{DB: db, FileDB: fileDB}

	app.Get("/api/v1/artist", wrapper.GetArtists)
	app.Get("/api/v1/artist/solo", wrapper.GetArtistSolos)
	app.Post("/api/v1/artist/solo", wrapper.CreateArtistSolo)
	app.Get("/api/v1/artist/group", wrapper.GetArtistGroups)
	app.Post("/api/v1/artist/group", wrapper.CreateArtistGroup)

	app.Post("/api/V1/media", wrapper.CreateMedia)
	app.Get("/api/V1/media", wrapper.GetMedia)
	app.Get("/api/v1/media/:mid/tracks", wrapper.GetMediaTracks)
	app.Post("/api/v1/media/:mid/upload", wrapper.CreateMediaTMPFile)
	app.Post("/api/v1/upload/:uid", wrapper.AppendMediaTMPFileChunk)
}

func startServer(app *fiber.App) {
	// listen on port 8081
	listenAddr := fmt.Sprintf(":%d", PORT)
	log.Fatal(app.Listen(listenAddr))
}
