package api

import (
	"fmt"

	"github.com/ejagombar/SpannerBackend/internal/spotify"
	"github.com/gofiber/fiber/v2"
)

func (s *SpannerController) TopTracks(c *fiber.Ctx) error {
	tokenData, err := s.getTokenData(c)
	if err != nil {
		return err
	}

	client, err := spotify.GetClient(c.Context(), tokenData)
	if err != nil {
		return err
	}

	timerange := fmt.Sprintf("%v", c.Params("timerange"))

	if timerange != "short_term" && timerange != "medium_term" && timerange != "long_term" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid time range")
	}

	topTracks, err := spotify.GetTopTracks(client, c.Context(), timerange)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(topTracks)
}

func (s *SpannerController) TopArtists(c *fiber.Ctx) error {
	tokenData, err := s.getTokenData(c)
	if err != nil {
		return err
	}

	client, err := spotify.GetClient(c.Context(), tokenData)
	if err != nil {
		return err
	}

	timerange := fmt.Sprintf("%v", c.Params("timerange"))

	if timerange != "short_term" && timerange != "medium_term" && timerange != "long_term" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid time range")
	}

	topTracks, err := spotify.GetTopArtists(client, c.Context(), timerange)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(topTracks)
}

func (s *SpannerController) ProfileInfo(c *fiber.Ctx) error {
	tokenData, err := s.getTokenData(c)
	if err != nil {
		return err
	}

	client, err := spotify.GetClient(c.Context(), tokenData)
	if err != nil {
		return err
	}

	User, err := spotify.GetUserProfileInfo(client, c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(User)
}

func (s *SpannerController) AllUserPlaylistIds(c *fiber.Ctx) error {
	tokenData, err := s.getTokenData(c)
	if err != nil {
		return err
	}

	client, err := spotify.GetClient(c.Context(), tokenData)
	if err != nil {
		return err
	}

	userID, err := spotify.GetUserID(client, c.Context())
	if err != nil {
		return err
	}

	playlistIDs, err := spotify.AllUserPlaylistIds(client, c.Context(), userID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(playlistIDs)
}
