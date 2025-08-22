package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RelatedSearchesResponse represents the RelatedSearchesResponse schema from the OpenAPI specification
type RelatedSearchesResponse struct {
	Terms []string `json:"terms"` // Related search terms
}

// CuratedListSimple represents the CuratedListSimple schema from the OpenAPI specification
type CuratedListSimple struct {
	Description string `json:"description,omitempty"` // This curated list's description.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Title string `json:"title,omitempty"` // Curated list name.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Minimum meta data of up to 5 podcasts in this curated list.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
}

// DeletePodcastResponse represents the DeletePodcastResponse schema from the OpenAPI specification
type DeletePodcastResponse struct {
	Status string `json:"status"` // The status of this podcast deletion request.
}

// SpellCheckResponse represents the SpellCheckResponse schema from the OpenAPI specification
type SpellCheckResponse struct {
	Corrected_text_html string `json:"corrected_text_html"` // The corrected text for the entire search term (multiple words/tokens), where misspelled tokens are replaced with the correct texts and html tags <b><i>
	Tokens []map[string]interface{} `json:"tokens"` // The word in the text query string that is not spelled correctly
}

// EpisodeSimple represents the EpisodeSimple schema from the OpenAPI specification
type EpisodeSimple struct {
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Podcast PodcastMinimum `json:"podcast,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Title string `json:"title,omitempty"` // Episode name.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Description string `json:"description,omitempty"` // Html of this episode's full description
}

// PodcastFull represents the PodcastFull schema from the OpenAPI specification
type PodcastFull struct {
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Genre_ids []int `json:"genre_ids,omitempty"`
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Title string `json:"title,omitempty"` // Podcast name.
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // Passed to the **next_episode_pub_date** parameter of `GET /podcasts/{id}` to paginate through episodes of that podcast.
	Extra PodcastExtraField `json:"extra,omitempty"`
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Episodes []EpisodeMinimum `json:"episodes,omitempty"`
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
}

// GetRegionsResponse represents the GetRegionsResponse schema from the OpenAPI specification
type GetRegionsResponse struct {
	Regions map[string]interface{} `json:"regions"`
}

// TypeaheadResponse represents the TypeaheadResponse schema from the OpenAPI specification
type TypeaheadResponse struct {
	Genres []Genre `json:"genres,omitempty"` // Genre suggestions. It'll show up when the **show_genres** parameter is *1*.
	Podcasts []PodcastTypeaheadResult `json:"podcasts,omitempty"` // Podcast suggestions. It'll show up when the **show_podcasts** parameter is 1.
	Terms []string `json:"terms"` // Search term suggestions.
}

// GetEpisodeRecommendationsResponse represents the GetEpisodeRecommendationsResponse schema from the OpenAPI specification
type GetEpisodeRecommendationsResponse struct {
	Recommendations []EpisodeSimple `json:"recommendations"`
}

// PlaylistResponse represents the PlaylistResponse schema from the OpenAPI specification
type PlaylistResponse struct {
	TypeField string `json:"type,omitempty"` // The type of this playlist, which should be either **episode_list** or **podcast_list**.
	Last_timestamp_ms int `json:"last_timestamp_ms,omitempty"` // Passed to the **last_timestamp_ms** parameter of `GET /playlists/{id}` to paginate through items of that playlist.
	Visibility string `json:"visibility,omitempty"` // Visibility of this playlist.
	Id string `json:"id,omitempty"` // A 11-character playlist id, which can be used to further fetch detailed playlist metadata via `GET /playlists/{id}`.
	Items []PlaylistItem `json:"items,omitempty"` // A list of playlist items.
	Name string `json:"name,omitempty"` // Playlist name.
	Total_audio_length_sec int `json:"total_audio_length_sec,omitempty"` // Total audio length of all episodes in this playlist, in seconds. It will have a valid value only when type is **episode_list**. In other words, it will be 0 if type is **podcast_list**.
	Image string `json:"image,omitempty"` // High resolution image url of the playlist.
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of the playlist.
	Total int `json:"total,omitempty"` // Total number of items in this playlist.
	Description string `json:"description,omitempty"` // Playlist description.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this playlist on ListenNotes.com.
}

// Genre represents the Genre schema from the OpenAPI specification
type Genre struct {
	Id int `json:"id,omitempty"` // Genre id
	Name string `json:"name,omitempty"` // Genre name.
	Parent_id int `json:"parent_id,omitempty"` // Parent genre id.
}

// PodcastMinimum represents the PodcastMinimum schema from the OpenAPI specification
type PodcastMinimum struct {
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title string `json:"title,omitempty"` // Podcast name.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
}

// EpisodeMinimum represents the EpisodeMinimum schema from the OpenAPI specification
type EpisodeMinimum struct {
	Title string `json:"title,omitempty"` // Episode name.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
}

// GetEpisodesInBatchForm represents the GetEpisodesInBatchForm schema from the OpenAPI specification
type GetEpisodesInBatchForm struct {
	Ids string `json:"ids"` // Comma-separated list of episode ids.
}

// PodcastDomainResponse represents the PodcastDomainResponse schema from the OpenAPI specification
type PodcastDomainResponse struct {
	Next_page_number int `json:"next_page_number,omitempty"`
	Page_number int `json:"page_number,omitempty"`
	Podcasts []PodcastSimple `json:"podcasts,omitempty"`
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
}

// PodcastSearchResult represents the PodcastSearchResult schema from the OpenAPI specification
type PodcastSearchResult struct {
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Genre_ids []int `json:"genre_ids,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Description_original string `json:"description_original,omitempty"` // Plain text of podcast description
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of podcast description
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
}

// GetLanguagesResponse represents the GetLanguagesResponse schema from the OpenAPI specification
type GetLanguagesResponse struct {
	Languages []string `json:"languages"`
}

// TrendingSearchesResponse represents the TrendingSearchesResponse schema from the OpenAPI specification
type TrendingSearchesResponse struct {
	Terms []string `json:"terms"` // Trending search terms
}

// PodcastAudienceResponse represents the PodcastAudienceResponse schema from the OpenAPI specification
type PodcastAudienceResponse struct {
	By_regions []map[string]interface{} `json:"by_regions,omitempty"`
}

// GetCuratedPodcastsResponse represents the GetCuratedPodcastsResponse schema from the OpenAPI specification
type GetCuratedPodcastsResponse struct {
	Curated_lists []CuratedListSimple `json:"curated_lists"`
	Has_next bool `json:"has_next"`
	Has_previous bool `json:"has_previous"`
	Next_page_number int `json:"next_page_number"`
	Page_number int `json:"page_number"`
	Previous_page_number int `json:"previous_page_number"`
	Total int `json:"total"`
}

// PodcastMinimumRss represents the PodcastMinimumRss schema from the OpenAPI specification
type PodcastMinimumRss struct {
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Title string `json:"title,omitempty"` // Podcast name.
}

// BestPodcastsResponse represents the BestPodcastsResponse schema from the OpenAPI specification
type BestPodcastsResponse struct {
	Listennotes_url string `json:"listennotes_url"` // Url of the list of best podcasts on [ListenNotes.com](https://www.ListenNotes.com).
	Next_page_number int `json:"next_page_number"`
	Page_number int `json:"page_number"`
	Total int `json:"total"`
	Id int `json:"id"` // The id of this genre
	Previous_page_number int `json:"previous_page_number"`
	Has_next bool `json:"has_next"`
	Has_previous bool `json:"has_previous"`
	Name string `json:"name"` // This genre's name.
	Parent_id int `json:"parent_id"` // The id of parent genre.
	Podcasts []PodcastSimple `json:"podcasts"`
}

// GetPodcastsInBatchResponse represents the GetPodcastsInBatchResponse schema from the OpenAPI specification
type GetPodcastsInBatchResponse struct {
	Podcasts []PodcastSimple `json:"podcasts"`
	Latest_episodes []EpisodeSimple `json:"latest_episodes,omitempty"` // Up to 10 latest episodes from these podcasts, sorted by **pub_date**. This field shows up only when **show_latest_episodes** is 1.
}

// PodcastTypeaheadResult represents the PodcastTypeaheadResult schema from the OpenAPI specification
type PodcastTypeaheadResult struct {
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of podcast name.
	Title_original string `json:"title_original,omitempty"` // Plain text of podcast name.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Publisher_highlighted string `json:"publisher_highlighted,omitempty"` // Highlighted segment of this podcast's publisher name.
	Publisher_original string `json:"publisher_original,omitempty"` // Plain text of this podcast's publisher name.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
}

// EpisodeSearchResult represents the EpisodeSearchResult schema from the OpenAPI specification
type EpisodeSearchResult struct {
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Link string `json:"link,omitempty"` // Web link of this episode.
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this episode's title
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this episode's description
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Podcast map[string]interface{} `json:"podcast,omitempty"` // The podcast that this episode belongs to.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Transcripts_highlighted []string `json:"transcripts_highlighted,omitempty"` // Up to 2 highlighted segments of the audio transcript of this episode.
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	Title_original string `json:"title_original,omitempty"` // Plain text of this episode' title
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Description_original string `json:"description_original,omitempty"` // Plain text of this episode's description
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
}

// GetPodcastsInBatchForm represents the GetPodcastsInBatchForm schema from the OpenAPI specification
type GetPodcastsInBatchForm struct {
	Spotify_ids string `json:"spotify_ids,omitempty"` // Comma-separated Spotify ids, e.g., 3DDfEsKDIDrTlnPOiG4ZF4
	Ids string `json:"ids,omitempty"` // Comma-separated list of podcast ids.
	Itunes_ids string `json:"itunes_ids,omitempty"` // Comma-separated Apple Podcasts (iTunes) ids, e.g., 659155419
	Next_episode_pub_date int `json:"next_episode_pub_date,omitempty"` // For latest episodes pagination. It's the value of **next_episode_pub_date** from the response of last request. If not specified, just return latest 15 episodes.
	Rsses string `json:"rsses,omitempty"` // Comma-separated rss urls.
	Show_latest_episodes int `json:"show_latest_episodes,omitempty"` // Whether or not to fetch up to 15 latest episodes from these podcasts, sorted by pub_date. 1 is yes, and 0 is no.
}

// PodcastSimple represents the PodcastSimple schema from the OpenAPI specification
type PodcastSimple struct {
	Id string `json:"id,omitempty"` // Podcast id, which can be used to further fetch detailed podcast metadata via `GET /podcasts/{id}`.
	Title string `json:"title,omitempty"` // Podcast name.
	Total_episodes int `json:"total_episodes,omitempty"` // Total number of episodes in this podcast.
	Listen_score int `json:"listen_score,omitempty"` // The estimated popularity score of a podcast compared to all other rss-based public podcasts in the world on a scale from 0 to 100. If the score is not available, it'll be null. Learn more at listennotes.com/listen-score
	Image string `json:"image,omitempty"` // Image url for this podcast's artwork. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image url for this podcast's artwork (300x300).
	Genre_ids []int `json:"genre_ids,omitempty"`
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this podcast on [ListenNotes.com](https://www.ListenNotes.com).
	Rss string `json:"rss,omitempty"` // RSS url of this podcast. This field is available only in the PRO/ENTERPRISE plan.
	TypeField string `json:"type,omitempty"` // The type of this podcast - episodic or serial.
	Update_frequency_hours int `json:"update_frequency_hours,omitempty"` // How frequently does this podcast release a new episode? In hours. For example, if the value is 166, then it's every 166 hours (or weekly).
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Average audio length of all episodes of this podcast. In seconds.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Looking_for PodcastLookingForField `json:"looking_for,omitempty"`
	Website string `json:"website,omitempty"` // Website url of this podcast.
	Language string `json:"language,omitempty"` // The language of this podcast. You can get all supported languages from `GET /languages`.
	Latest_pub_date_ms int `json:"latest_pub_date_ms,omitempty"` // The published date of the latest episode of this podcast. In milliseconds
	Publisher string `json:"publisher,omitempty"` // Podcast publisher name.
	Country string `json:"country,omitempty"` // The country where this podcast is produced.
	Itunes_id int `json:"itunes_id,omitempty"` // iTunes id for this podcast.
	Earliest_pub_date_ms int `json:"earliest_pub_date_ms,omitempty"` // The published date of the oldest episode of this podcast. In milliseconds
	Extra PodcastExtraField `json:"extra,omitempty"`
	Is_claimed bool `json:"is_claimed,omitempty"` // Whether this podcast is claimed by its producer on [ListenNotes.com](https://www.ListenNotes.com).
	Latest_episode_id string `json:"latest_episode_id,omitempty"` // The id of the most recently published episode of this podcast, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Email string `json:"email,omitempty"` // The email of this podcast's producer. This field is available only in the PRO/ENTERPRISE plan.
	Listen_score_global_rank string `json:"listen_score_global_rank,omitempty"` // The estimated popularity ranking of a podcast compared to all other rss-based public podcasts in the world. For example, if the value is 0.5%, then this podcast is one of the top 0.5% most popular shows out of all podcasts globally, ranked by Listen Score. If the ranking is not available, it'll be null. Learn more at listennotes.com/listen-score
}

// GetEpisodesInBatchResponse represents the GetEpisodesInBatchResponse schema from the OpenAPI specification
type GetEpisodesInBatchResponse struct {
	Episodes []EpisodeSimple `json:"episodes"`
}

// DeletedItem represents the DeletedItem schema from the OpenAPI specification
type DeletedItem struct {
	ErrorField string `json:"error,omitempty"` // Why this episode or podcast is deleted?
	Id string `json:"id,omitempty"` // Episode id or podcast id.
	Status string `json:"status,omitempty"` // The status of this episode or podcast. For now, the only possible value is **deleted**.
	Title string `json:"title,omitempty"` // Episode title or podcast title.
}

// PodcastExtraField represents the PodcastExtraField schema from the OpenAPI specification
type PodcastExtraField struct {
	Amazon_music_url string `json:"amazon_music_url,omitempty"` // Amazon Music url for this podcast
	Instagram_handle string `json:"instagram_handle,omitempty"` // Instagram username affiliated with this podcast
	Linkedin_url string `json:"linkedin_url,omitempty"` // LinkedIn url affiliated with this podcast
	Twitter_handle string `json:"twitter_handle,omitempty"` // Twitter username affiliated with this podcast
	Url2 string `json:"url2,omitempty"` // Url affiliated with this podcast
	Spotify_url string `json:"spotify_url,omitempty"` // Spotify url for this podcast
	Url1 string `json:"url1,omitempty"` // Url affiliated with this podcast
	Facebook_handle string `json:"facebook_handle,omitempty"` // Facebook username affiliated with this podcast
	Google_url string `json:"google_url,omitempty"` // Google Podcasts url for this podcast
	Patreon_handle string `json:"patreon_handle,omitempty"` // Patreon username affiliated with this podcast
	Url3 string `json:"url3,omitempty"` // Url affiliated with this podcast
	Wechat_handle string `json:"wechat_handle,omitempty"` // WeChat username affiliated with this podcast
	Youtube_url string `json:"youtube_url,omitempty"` // YouTube url affiliated with this podcast
}

// PodcastLookingForField represents the PodcastLookingForField schema from the OpenAPI specification
type PodcastLookingForField struct {
	Cohosts bool `json:"cohosts,omitempty"` // Whether this podcast is looking for cohosts.
	Cross_promotion bool `json:"cross_promotion,omitempty"` // Whether this podcast is looking for cross promotion opportunities with other podcasts.
	Guests bool `json:"guests,omitempty"` // Whether this podcast is looking for guests.
	Sponsors bool `json:"sponsors,omitempty"` // Whether this podcast is looking for sponsors.
}

// PlaylistItem represents the PlaylistItem schema from the OpenAPI specification
type PlaylistItem struct {
	Added_at_ms int `json:"added_at_ms,omitempty"` // Timestamp (in milliseconds) when this item is added.
	Data interface{} `json:"data,omitempty"`
	Id int `json:"id,omitempty"` // Playlist item id.
	Notes string `json:"notes,omitempty"` // Notes for this item.
	TypeField string `json:"type,omitempty"` // The type of this playlist item. If a playlist is **episode_list**, then an item could be either **episode** or **custom_audio**. If it's **podcast_list**, then an item can only be **podcast**.
}

// GetGenresResponse represents the GetGenresResponse schema from the OpenAPI specification
type GetGenresResponse struct {
	Genres []Genre `json:"genres"`
}

// EpisodeFull represents the EpisodeFull schema from the OpenAPI specification
type EpisodeFull struct {
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length of this episode. In seconds.
	Thumbnail string `json:"thumbnail,omitempty"` // Thumbnail image (300x300) url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork thumbnail image.
	Audio string `json:"audio,omitempty"` // Audio url of this episode, which can be played directly.
	Transcript string `json:"transcript,omitempty"` // The transcript of this episode, in plain text (with the newline character \n). If there's not transcript, it is null. This field is available only in the PRO/ENTERPRISE plan.
	Listennotes_edit_url string `json:"listennotes_edit_url,omitempty"` // Edit url of this episode where you can update the audio url if you find the audio is broken.
	Podcast PodcastSimple `json:"podcast,omitempty"`
	Title string `json:"title,omitempty"` // Episode name.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this episode on [ListenNotes.com](https://www.ListenNotes.com).
	Id string `json:"id,omitempty"` // Episode id, which can be used to further fetch detailed episode metadata via `GET /episodes/{id}`.
	Image string `json:"image,omitempty"` // Image url for this episode. If an episode doesn't have its own image, then this field would be the url of the podcast artwork image. If you are using PRO/ENTERPRISE plan, then it's a high resolution image (1400x1400). If you are using FREE plan, then it's the same as **thumbnail**, low resolution image (300x300).
	Maybe_audio_invalid bool `json:"maybe_audio_invalid,omitempty"` // Whether or not this episode's audio is invalid. Podcasters may delete the original audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date for this episode. In millisecond.
	Explicit_content bool `json:"explicit_content,omitempty"` // Whether this podcast contains explicit language.
	Description string `json:"description,omitempty"` // Html of this episode's full description
	Link string `json:"link,omitempty"` // Web link of this episode.
}

// CuratedListFull represents the CuratedListFull schema from the OpenAPI specification
type CuratedListFull struct {
	Podcasts []PodcastSimple `json:"podcasts,omitempty"` // Complete meta data of all podcasts in this curated list.
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Description string `json:"description,omitempty"` // This curated list's description.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Title string `json:"title,omitempty"` // Curated list name.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
}

// SubmitPodcastResponse represents the SubmitPodcastResponse schema from the OpenAPI specification
type SubmitPodcastResponse struct {
	Podcast PodcastMinimum `json:"podcast"`
	Status string `json:"status"` // The status of this submission.
}

// GetPodcastRecommendationsResponse represents the GetPodcastRecommendationsResponse schema from the OpenAPI specification
type GetPodcastRecommendationsResponse struct {
	Recommendations []PodcastSimple `json:"recommendations"`
}

// CustomAudio represents the CustomAudio schema from the OpenAPI specification
type CustomAudio struct {
	Audio string `json:"audio,omitempty"` // Audio url, which can be played directly.
	Audio_length_sec int `json:"audio_length_sec,omitempty"` // Audio length in seconds.
	Image string `json:"image,omitempty"` // High resolution image url of this custom audio.
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date (in milliseconds) of this custom audio. For now, it's the same as **added_at_ms** of this playlist item.
	Thumbnail string `json:"thumbnail,omitempty"` // Low resolution image url of this custom audio.
	Title string `json:"title,omitempty"` // Custom audio title.
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	Count int `json:"count,omitempty"` // The number of search results in this page.
	Next_offset int `json:"next_offset,omitempty"` // Pass this value to the **offset** parameter to do pagination of search results.
	Results []interface{} `json:"results,omitempty"` // A list of search results.
	Took float64 `json:"took,omitempty"` // The time it took to fetch these search results. In seconds.
	Total int `json:"total,omitempty"` // The total number of search results.
}

// CuratedListSearchResult represents the CuratedListSearchResult schema from the OpenAPI specification
type CuratedListSearchResult struct {
	Title_original string `json:"title_original,omitempty"` // Plain text of this curated list's title
	Listennotes_url string `json:"listennotes_url,omitempty"` // The url of this curated list on [ListenNotes.com](https://www.ListenNotes.com).
	Podcasts []PodcastMinimum `json:"podcasts,omitempty"` // Up to 5 podcasts in this curated list.
	Description_original string `json:"description_original,omitempty"` // Plain text of this curated list's description
	Pub_date_ms int `json:"pub_date_ms,omitempty"` // Published date of this curated list. In milliseconds.
	Source_url string `json:"source_url,omitempty"` // Url of the source of this curated list.
	Title_highlighted string `json:"title_highlighted,omitempty"` // Highlighted segment of this curated list's title
	Id string `json:"id,omitempty"` // Curated list id, which can be used to further fetch detailed curated list metadata via `GET /curated_podcasts/{id}`.
	Total int `json:"total,omitempty"` // The total number of podcasts in this curated list.
	Description_highlighted string `json:"description_highlighted,omitempty"` // Highlighted segment of this curated list's description
	Source_domain string `json:"source_domain,omitempty"` // The domain name of the source of this curated list.
}

// PlaylistsResponse represents the PlaylistsResponse schema from the OpenAPI specification
type PlaylistsResponse struct {
	Previous_page_number int `json:"previous_page_number,omitempty"`
	Total int `json:"total,omitempty"`
	Has_next bool `json:"has_next,omitempty"`
	Has_previous bool `json:"has_previous,omitempty"`
	Next_page_number int `json:"next_page_number,omitempty"`
	Page_number int `json:"page_number,omitempty"`
	Playlists []map[string]interface{} `json:"playlists,omitempty"`
}

// SubmitPodcastForm represents the SubmitPodcastForm schema from the OpenAPI specification
type SubmitPodcastForm struct {
	Email string `json:"email,omitempty"` // A valid email address. If **email** is specified, then we'll notify this email address once the podcast is accepted.
	Rss string `json:"rss"` // A valid podcast rss url.
}
