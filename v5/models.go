package openrec

import (
	_ "encoding/json"
	"time"
	// "gopkg.in/guregu/null.v4"
)

const (
	SORT_LIVE_VIEWS = "live_views"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Errors  []struct {
		Field    []string `json:"field"`
		Location string   `json:"location"`
		Messages []string `json:"messages"`
		Types    []string `json:"types"`
	} `json:"errors"`
}

type Movie struct {
	ID                string        `json:"id"`
	MovieID           int           `json:"movie_id"`
	Title             string        `json:"title"`
	Introduction      string        `json:"introduction"`
	IsLive            bool          `json:"is_live"`
	OnairStatus       int           `json:"onair_status"`
	MovieType         string        `json:"movie_type"`
	UploadStatus      string        `json:"upload_status"`
	MonetizeStatus    int           `json:"monetize_status"`
	ThumbnailURL      string        `json:"thumbnail_url"`
	LThumbnailURL     string        `json:"l_thumbnail_url"`
	SThumbnailURL     string        `json:"s_thumbnail_url"`
	LSpriteImageURL   string        `json:"l_sprite_image_url"`
	SSpriteImageURL   string        `json:"s_sprite_image_url"`
	SpriteIntervals   []int         `json:"sprite_intervals"`
	SpriteImage       SpriteImage   `json:"sprite_image"`
	CoverImageURL     string        `json:"cover_image_url"`
	IsCoverImageIcon  bool          `json:"is_cover_image_icon"`
	LiveViews         int           `json:"live_views"`
	TotalViews        int           `json:"total_views"`
	TotalYells        int           `json:"total_yells"`
	IsMobile          bool          `json:"is_mobile"`
	IsLowLatency      bool          `json:"is_low_latency"`
	IsDvr             bool          `json:"is_dvr"`
	IsCapture         bool          `json:"is_capture"`
	IsFixedPhrase     bool          `json:"is_fixed_phrase"`
	EnabledAd         bool          `json:"enabled_ad"`
	EnabledYell       bool          `json:"enabled_yell"`
	EnabledOwnerYell  bool          `json:"enabled_owner_yell"`
	EnabledCastYell   bool          `json:"enabled_cast_yell"`
	Ad                bool          `json:"ad"`
	CreatedAt         time.Time     `json:"created_at"`
	PublishedAt       time.Time     `json:"published_at"`
	StartedAt         time.Time     `json:"started_at"`
	EndedAt           time.Time     `json:"ended_at"`
	WillStartAt       interface{}   `json:"will_start_at"`
	WillEndAt         interface{}   `json:"will_end_at"`
	StartTime         int           `json:"start_time"`
	PlayTime          int           `json:"play_time"`
	IsBanned          bool          `json:"is_banned"`
	BanType           int           `json:"ban_type"`
	Orientation       int           `json:"orientation"`
	DeviceType        int           `json:"device_type"`
	BroadcastType     string        `json:"broadcast_type"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	ConnectCount      int           `json:"connect_count"`
	PlayListID        int           `json:"play_list_id"`
	Tags              []interface{} `json:"tags"`
	Media             Media         `json:"media"`
	SubsTrialMedia    SubTrialMedia `json:"subs_trial_media"`
	Channel           Channel       `json:"channel"`
	SubsChannel       interface{}   `json:"subs_channel"`
	ChatSetting       ChatSetting   `json:"chat_setting"`
	Game              Game          `json:"game"`
	Next              interface{}   `json:"next"`
	Casts             []interface{} `json:"casts"`
	Popularity        int           `json:"popularity"`
	IsLiveViewers     bool          `json:"is_live_viewers"`
	PublicType        string        `json:"public_type"`
	ChatPublicType    string        `json:"chat_public_type"`
	ArchivePublicType string        `json:"archive_public_type"`
	Poll              interface{}   `json:"poll"`
	PollUseTarget     interface{}   `json:"poll_use_target"`
	ViewHistory       ViewHistory   `json:"view_history"`
}

type SpriteImage struct {
	URL       string `json:"url"`
	Interval  int    `json:"interval"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Cols      int    `json:"cols"`
	Rows      int    `json:"rows"`
	StartPage int    `json:"start_page"`
	Ext       string `json:"ext"`
}

type Media struct {
	URL           string      `json:"url"`
	URLSource     interface{} `json:"url_source"`
	URLPublic     string      `json:"url_public"`
	URLTrailer    interface{} `json:"url_trailer"`
	URLDvr        string      `json:"url_dvr"`
	URLDvrAudio   string      `json:"url_dvr_audio"`
	URLAudio      string      `json:"url_audio"`
	URLLowLatency string      `json:"url_low_latency"`
	URLUll        string      `json:"url_ull"`
}

type SubTrialMedia struct {
	URL           string `json:"url"`
	URLLowLatency string `json:"url_low_latency"`
	URLUll        string `json:"url_ull"`
	URLDvr        string `json:"url_dvr"`
	URLAudio      string `json:"url_audio"`
	URLDvrAudio   string `json:"url_dvr_audio"`
}

type Channel struct {
	ID                   string        `json:"id"`
	RecxuserID           int           `json:"recxuser_id"`
	OpenrecUserID        int           `json:"openrec_user_id"`
	Name                 string        `json:"name"`
	Nickname             string        `json:"nickname"`
	Introduction         string        `json:"introduction"`
	IconImageURL         string        `json:"icon_image_url"`
	LIconImageURL        string        `json:"l_icon_image_url"`
	CoverImageURL        string        `json:"cover_image_url"`
	LCoverImageURL       string        `json:"l_cover_image_url"`
	Movies               int           `json:"movies"`
	Views                int           `json:"views"`
	Follows              int           `json:"follows"`
	Followers            int           `json:"followers"`
	IsPremium            bool          `json:"is_premium"`
	IsOfficial           bool          `json:"is_official"`
	IsFresh              bool          `json:"is_fresh"`
	IsWarned             bool          `json:"is_warned"`
	IsLeagueYell         bool          `json:"is_league_yell"`
	IsTeam               bool          `json:"is_team"`
	IsLive               bool          `json:"is_live"`
	LiveViews            int           `json:"live_views"`
	UserStatus           string        `json:"user_status"`
	IsCreator            bool          `json:"is_creator"`
	ChatRule             string        `json:"chat_rule"`
	LastUpdatedAt        time.Time     `json:"last_updated_at"`
	TwitterScreenName    interface{}   `json:"twitter_screen_name"`
	GaTrackingID         interface{}   `json:"ga_tracking_id"`
	Blacklist            []interface{} `json:"blacklist"`
	Games                interface{}   `json:"games"`
	OnairBroadcastMovies interface{}   `json:"onair_broadcast_movies"`
}

type ChatSetting struct {
	NameColor                   string `json:"name_color"`
	IsPremiumHidden             bool   `json:"is_premium_hidden"`
	LimitedContinuousChat       bool   `json:"limited_continuous_chat"`
	ContinuousChatThreshold     int    `json:"continuous_chat_threshold"`
	LimitedUnfollowerChat       bool   `json:"limited_unfollower_chat"`
	UnfollowerChatThreshold     int    `json:"unfollower_chat_threshold"`
	LimitedFreshUserChat        bool   `json:"limited_fresh_user_chat"`
	FreshUserChatThreshold      int    `json:"fresh_user_chat_threshold"`
	LimitedTemporaryBlacklist   bool   `json:"limited_temporary_blacklist"`
	TemporaryBlacklistThreshold int    `json:"temporary_blacklist_threshold"`
	LimitedWarnedUserChat       bool   `json:"limited_warned_user_chat"`
	ChatRule                    string `json:"chat_rule"`
}

type Game struct {
	ID                string      `json:"id"`
	GameID            int         `json:"game_id"`
	Title             string      `json:"title"`
	Introduction      string      `json:"introduction"`
	TitleImageURL     string      `json:"title_image_url"`
	CoverImageURL     string      `json:"cover_image_url"`
	AppStoreURL       string      `json:"app_store_url"`
	PlayStoreURL      string      `json:"play_store_url"`
	SchemaURL         string      `json:"schema_url"`
	PackageName       string      `json:"package_name"`
	LicenseStatus     int         `json:"license_status"`
	NoticeMessage     string      `json:"notice_message"`
	MonetizeStatus    int         `json:"monetize_status"`
	CeroRating        int         `json:"cero_rating"`
	IsPortrait        bool        `json:"is_portrait"`
	ReleasedAt        time.Time   `json:"released_at"`
	Maker             interface{} `json:"maker"`
	PlatformMobile    string      `json:"platform_mobile"`
	PlatformPc        string      `json:"platform_pc"`
	PlatformConsole   string      `json:"platform_console"`
	PlatformArcade    string      `json:"platform_arcade"`
	IsPlatformMobile  bool        `json:"is_platform_mobile"`
	IsPlatformPc      bool        `json:"is_platform_pc"`
	IsPlatformConsole bool        `json:"is_platform_console"`
	IsPlatformArcade  bool        `json:"is_platform_arcade"`
	Stats             struct {
		LiveViews    int         `json:"live_views"`
		MovieCount   interface{} `json:"movie_count"`
		TotalViews   int         `json:"total_views"`
		CreatorCount interface{} `json:"creator_count"`
		LastUpdateAt time.Time   `json:"last_update_at"` // time?
	} `json:"stats"`
}

type ViewHistory struct {
	ViewsAt      interface{} `json:"views_at"`
	PlayPosition int         `json:"play_position"`
}
