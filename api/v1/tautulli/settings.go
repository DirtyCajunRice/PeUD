package tautulli

type Response struct {
	Result  string   `json:"result"`
	Message string   `json:"message"`
	Data    Settings `json:"data"`
}
type Settings struct {
	IFTTT      IFTTT      `json:"IFTTT"`
	Growl      Growl      `json:"Growl"`
	PMS        PMS        `json:"PMS"`
	Monitoring Monitoring `json:"Monitoring"`
	General    General    `json:"General"`
	Slack      Slack      `json:"Slack"`
	Boxcar     Boxcar     `json:"Boxcar"`
	PushBullet PushBullet `json:"PushBullet"`
	Scripts    Scripts    `json:"Scripts"`
	OSXNotify  OSXNotify  `json:"OSX_Notify"`
	NMA        NMA        `json:"NMA"`
	Hipchat    Hipchat    `json:"Hipchat"`
	Pushalot   Pushalot   `json:"Pushalot"`
	Prowl      Prowl      `json:"Prowl"`
	Advanced   Advanced   `json:"Advanced"`
	Facebook   Facebook   `json:"Facebook"`
	Email      Email      `json:"Email"`
	PlexWatch  PlexWatch  `json:"PlexWatch"`
	Pushover   Pushover   `json:"Pushover"`
	Telegram   Telegram   `json:"Telegram"`
	Twitter    Twitter    `json:"Twitter"`
	Plex       Plex       `json:"Plex"`
	Join       Join       `json:"Join"`
	XBMC       XBMC       `json:"XBMC"`
	Newsletter Newsletter `json:"Newsletter"`
	Browser    Browser    `json:"Browser"`
	Cloudinary Cloudinary `json:"Cloudinary"`
}
type IFTTT struct {
	IFTTTOnStop       bool   `json:"ifttt_on_stop"`
	IFTTTOnPause      bool   `json:"ifttt_on_pause"`
	IFTTTKey          string `json:"ifttt_key"`
	IFTTTOnConcurrent bool   `json:"ifttt_on_concurrent"`
	IFTTTOnCreated    bool   `json:"ifttt_on_created"`
	IFTTTOnBuffer     bool   `json:"ifttt_on_buffer"`
	IFTTTEnabled      bool   `json:"ifttt_enabled"`
	IFTTTOnPMSUpdate  bool   `json:"ifttt_on_pmsupdate"`
	IFTTTOnNewDevice  bool   `json:"ifttt_on_newdevice"`
	IFTTTOnIntDown    bool   `json:"ifttt_on_intdown"`
	IFTTTOnExtUp      bool   `json:"ifttt_on_extup"`
	IFTTTOnIntUp      bool   `json:"ifttt_on_intup"`
	IFTTTOnResume     bool   `json:"ifttt_on_resume"`
	IFTTTOnWatched    bool   `json:"ifttt_on_watched"`
	IFTTTEvent        string `json:"ifttt_event"`
	IFTTTOnExtDown    bool   `json:"ifttt_on_extdown"`
	IFTTTOnPlay       bool   `json:"ifttt_on_play"`
}
type Growl struct {
	GrowlEnabled      bool   `json:"growl_enabled"`
	GrowlOnPlay       bool   `json:"growl_on_play"`
	GrowlOnStop       bool   `json:"growl_on_stop"`
	GrowlOnBuffer     bool   `json:"growl_on_buffer"`
	GrowlOnPause      bool   `json:"growl_on_pause"`
	GrowlOnIntDown    bool   `json:"growl_on_intdown"`
	GrowlHost         string `json:"growl_host"`
	GrowlOnExtUp      bool   `json:"growl_on_extup"`
	GrowlOnNewDevice  bool   `json:"growl_on_newdevice"`
	GrowlOnPMSUpdate  bool   `json:"growl_on_pmsupdate"`
	GrowlPassword     string `json:"growl_password"`
	GrowlOnResume     bool   `json:"growl_on_resume"`
	GrowlOnCreated    bool   `json:"growl_on_created"`
	GrowlOnConcurrent bool   `json:"growl_on_concurrent"`
	GrowlOnWatched    bool   `json:"growl_on_watched"`
	GrowlOnIntUp      bool   `json:"growl_on_intup"`
	GrowlOnExtDown    bool   `json:"growl_on_extdown"`
}
type PMS struct {
	PMSUseBif            bool   `json:"pms_use_bif"`
	PMSUUID              string `json:"pms_uuid"`
	PMSURLManual         int    `json:"pms_url_manual"`
	PMSWebURL            string `json:"pms_web_url"`
	PMSURL               string `json:"pms_url"`
	PMSUpdateChannel     string `json:"pms_update_channel"`
	PMSUpdateDistro      string `json:"pms_update_distro"`
	PMSName              string `json:"pms_name"`
	PMSLogsFolder        string `json:"pms_logs_folder"`
	PMSPlexPass          int    `json:"pms_plexpass"`
	PMSToken             string `json:"pms_token"`
	PMSURLOverride       string `json:"pms_url_override"`
	PMSPort              int    `json:"pms_port"`
	PMSIsRemote          int    `json:"pms_is_remote"`
	PMSIdentifier        string `json:"pms_identifier"`
	PMSSSL               int    `json:"pms_ssl"`
	PMSVersion           string `json:"pms_version"`
	PMSUpdateDistroBuild string `json:"pms_update_distro_build"`
	PMSIP                string `json:"pms_ip"`
	PMSLogsLineCap       string `json:"pms_logs_line_cap"`
	PMSPlatform          string `json:"pms_platform"`
	PMSIsCloud           int    `json:"pms_is_cloud"`
}
type Monitoring struct {
	NotifyUploadPosters                 bool   `json:"notify_upload_posters"`
	MovieNotifyEnable                   bool   `json:"movie_notify_enable"`
	NotifyOnResumeBodyText              string `json:"notify_on_resume_body_text"`
	NotifyOnPMSUpdateBodyText           string `json:"notify_on_pmsupdate_body_text"`
	RefreshLibrariesOnStartup           bool   `json:"refresh_libraries_on_startup"`
	NotifyOnCreatedBodyText             string `json:"notify_on_created_body_text"`
	BufferWait                          string `json:"buffer_wait"`
	NotifyOnIntUpSubjectText            string `json:"notify_on_intup_subject_text"`
	NotifyConcurrentByIP                bool   `json:"notify_concurrent_by_ip"`
	NotifyOnNewDeviceSubjectText        string `json:"notify_on_newdevice_subject_text"`
	NotifyOnExtUpSubjectText            string `json:"notify_on_extup_subject_text"`
	NotifyOnWatchedSubjectText          string `json:"notify_on_watched_subject_text"`
	NotifyOnPauseSubjectText            string `json:"notify_on_pause_subject_text"`
	NotifyOnExtUpBodyText               string `json:"notify_on_extup_body_text"`
	TvWatchedPercent                    string `json:"tv_watched_percent"`
	NotifyOnWatchedBodyText             string `json:"notify_on_watched_body_text"`
	MusicNotifyOnPause                  bool   `json:"music_notify_on_pause"`
	NotifyGroupRecentlyAddedGrandparent bool   `json:"notify_group_recently_added_grandparent"`
	NotifyOnStartBodyText               string `json:"notify_on_start_body_text"`
	NotifyOnPauseBodyText               string `json:"notify_on_pause_body_text"`
	NotifyOnBufferBodyText              string `json:"notify_on_buffer_body_text"`
	MusicNotifyOnStop                   bool   `json:"music_notify_on_stop"`
	NotifyRecentlyAddedUpgrade          bool   `json:"notify_recently_added_upgrade"`
	NotifyOnStopBodyText                string `json:"notify_on_stop_body_text"`
	NotifyOnCreatedSubjectText          string `json:"notify_on_created_subject_text"`
	NotifyOnExtdownBodyText             string `json:"notify_on_extdown_body_text"`
	NotifyScriptsArgsText               string `json:"notify_scripts_args_text"`
	MusicLoggingEnable                  bool   `json:"music_logging_enable"`
	MovieNotifyOnStop                   bool   `json:"movie_notify_on_stop"`
	NotifyWatchedPercent                string `json:"notify_watched_percent"`
	NotifyGroupRecentlyAdded            bool   `json:"notify_group_recently_added"`
	NotifyOnConcurrentSubjectText       string `json:"notify_on_concurrent_subject_text"`
	NotifyOnConcurrentBodyText          string `json:"notify_on_concurrent_body_text"`
	TvLoggingEnable                     bool   `json:"tv_logging_enable"`
	NotifyGroupRecentlyAddedParent      bool   `json:"notify_group_recently_added_parent"`
	TvNotifyEnable                      bool   `json:"tv_notify_enable"`
	NotifyRecentlyAddedDelay            string `json:"notify_recently_added_delay"`
	NotifyOnIntDownBodyText             string `json:"notify_on_intdown_body_text"`
	VideoLoggingEnable                  bool   `json:"video_logging_enable"`
	LoggingIgnoreInterval               string `json:"logging_ignore_interval"`
	NotifyOnNewDeviceBodyText           string `json:"notify_on_newdevice_body_text"`
	RefreshUsersInterval                string `json:"refresh_users_interval"`
	MonitoringUseWebsocket              bool   `json:"monitoring_use_websocket"`
	MonitorPMSUpdates                   bool   `json:"monitor_pms_updates"`
	NotifyOnStartSubjectText            string `json:"notify_on_start_subject_text"`
	NotifyRecentlyAdded                 bool   `json:"notify_recently_added"`
	NotifyOnPMSUpdateSubjectText        string `json:"notify_on_pmsupdate_subject_text"`
	TvNotifyOnStop                      bool   `json:"tv_notify_on_stop"`
	MusicNotifyEnable                   bool   `json:"music_notify_enable"`
	MovieNotifyOnStart                  bool   `json:"movie_notify_on_start"`
	RefreshUsersOnStartup               bool   `json:"refresh_users_on_startup"`
	BufferThreshold                     string `json:"buffer_threshold"`
	NotifyOnExtdownSubjectText          string `json:"notify_on_extdown_subject_text"`
	RefreshLibrariesInterval            string `json:"refresh_libraries_interval"`
	TvNotifyOnStart                     bool   `json:"tv_notify_on_start"`
	NotifyRecentlyAddedGrandparent      bool   `json:"notify_recently_added_grandparent"`
	NotifyOnResumeSubjectText           string `json:"notify_on_resume_subject_text"`
	NotifyConcurrentThreshold           string `json:"notify_concurrent_threshold"`
	NotifyOnStopSubjectText             string `json:"notify_on_stop_subject_text"`
	MovieNotifyOnPause                  bool   `json:"movie_notify_on_pause"`
	NotifyConsecutive                   bool   `json:"notify_consecutive"`
	MovieLoggingEnable                  bool   `json:"movie_logging_enable"`
	NotifyOnIntUpBodyText               string `json:"notify_on_intup_body_text"`
	NotifyOnIntDownSubjectText          string `json:"notify_on_intdown_subject_text"`
	MusicNotifyOnStart                  bool   `json:"music_notify_on_start"`
	TvNotifyOnPause                     bool   `json:"tv_notify_on_pause"`
	MonitorRemoteAccess                 bool   `json:"monitor_remote_access"`
	ImgurClientID                       string `json:"imgur_client_id"`
	MovieWatchedPercent                 string `json:"movie_watched_percent"`
	MonitoringInterval                  string `json:"monitoring_interval"`
	MusicWatchedPercent                 string `json:"music_watched_percent"`
	NotifyOnBufferSubjectText           string `json:"notify_on_buffer_subject_text"`
	LoggingLiveTv                       bool   `json:"logging_live_tv"`
	NotifyContinuedSessionThreshold     string `json:"notify_continued_session_threshold"`
}
type GetFileSizesHold struct {
	SectionIds []string `json:"section_ids"`
	RatingKeys []string `json:"rating_keys"`
}
type General struct {
	WeekStartMonday             bool             `json:"week_start_monday"`
	GitToken                    string           `json:"git_token"`
	HomeStatsCards              []string         `json:"home_stats_cards"`
	BackupInterval              string           `json:"backup_interval"`
	GraphDays                   string           `json:"graph_days"`
	GraphType                   string           `json:"graph_type"`
	HTTPPlexAdmin               bool             `json:"http_plex_admin"`
	GetFileSizes                bool             `json:"get_file_sizes"`
	HTTPBasicAuth               bool             `json:"http_basic_auth"`
	APISQL                      bool             `json:"api_sql"`
	UpdateShowChangelog         bool             `json:"update_show_changelog"`
	HTTPRoot                    string           `json:"http_root"`
	HTTPProxy                   bool             `json:"http_proxy"`
	FreezeDb                    bool             `json:"freeze_db"`
	ShowAdvancedSettings        bool             `json:"show_advanced_settings"`
	IPLoggingEnable             bool             `json:"ip_logging_enable"`
	LaunchBrowser               bool             `json:"launch_browser"`
	UpdateLabels                bool             `json:"update_labels"`
	HTTPUsername                string           `json:"http_username"`
	GitRemote                   string           `json:"git_remote"`
	HTTPSCreateCert             bool             `json:"https_create_cert"`
	DoNotOverrideGitBranch      bool             `json:"do_not_override_git_branch"`
	HomeStatsLength             string           `json:"home_stats_length"`
	HTTPHost                    string           `json:"http_host"`
	HTTPSCertChain              string           `json:"https_cert_chain"`
	FirstRunComplete            bool             `json:"first_run_complete"`
	GitRepo                     string           `json:"git_repo"`
	HomeStatsRecentlyAddedCount string           `json:"home_stats_recently_added_count"`
	EnableHTTPS                 bool             `json:"enable_https"`
	CheckGithub                 bool             `json:"check_github"`
	APIEnabled                  bool             `json:"api_enabled"`
	GitPath                     string           `json:"git_path"`
	TheMovieDBAPIKey            string           `json:"themoviedb_apikey"`
	TVMazeLookup                bool             `json:"tvmaze_lookup"`
	CheckGithubInterval         string           `json:"check_github_interval"`
	HTTPSIP                     string           `json:"https_ip"`
	UpdateNotifiersDb           bool             `json:"update_notifiers_db"`
	GroupHistoryTables          bool             `json:"group_history_tables"`
	HTTPSDomain                 string           `json:"https_domain"`
	GraphTab                    string           `json:"graph_tab"`
	GeoipDb                     string           `json:"geoip_db"`
	HTTPPassword                string           `json:"http_password"`
	HomeStatsCount              string           `json:"home_stats_count"`
	BackupDir                   string           `json:"backup_dir"`
	GitUser                     string           `json:"git_user"`
	Interface                   string           `json:"interface"`
	CacheImages                 bool             `json:"cache_images"`
	LogBlacklist                bool             `json:"log_blacklist"`
	GraphMonths                 string           `json:"graph_months"`
	TheMovieDBLookup            bool             `json:"themoviedb_lookup"`
	CleanupFiles                bool             `json:"cleanup_files"`
	PlexpyAutoUpdate            bool             `json:"plexpy_auto_update"`
	CheckGithubOnStartup        bool             `json:"check_github_on_startup"`
	HTTPPort                    string           `json:"http_port"`
	LogDir                      string           `json:"log_dir"`
	UpdateDbInterval            string           `json:"update_db_interval"`
	AllowGuestAccess            bool             `json:"allow_guest_access"`
	TimeFormat                  string           `json:"time_format"`
	CacheDir                    string           `json:"cache_dir"`
	HomeSections                []string         `json:"home_sections"`
	BackupDays                  string           `json:"backup_days"`
	HTTPSCert                   string           `json:"https_cert"`
	HomeRefreshInterval         string           `json:"home_refresh_interval"`
	APIKey                      string           `json:"api_key"`
	UpdateSectionIds            bool             `json:"update_section_ids"`
	DateFormat                  string           `json:"date_format"`
	HomeStatsType               bool             `json:"home_stats_type"`
	UpdateLibrariesDbNotify     bool             `json:"update_libraries_db_notify"`
	HTTPEnvironment             string           `json:"http_environment"`
	HistoryTableActivity        bool             `json:"history_table_activity"`
	HTTPBaseURL                 string           `json:"http_base_url"`
	HTTPSKey                    string           `json:"https_key"`
	HTTPHashPassword            bool             `json:"http_hash_password"`
	HomeLibraryCards            []string         `json:"home_library_cards"`
	HTTPHashedPassword          bool             `json:"http_hashed_password"`
	GitBranch                   string           `json:"git_branch"`
	AnonRedirect                string           `json:"anon_redirect"`
	WinSysTray                  bool             `json:"win_sys_tray"`
	MusicBrainzLookup           bool             `json:"musicbrainz_lookup"`
	MaxmindLicenseKey           string           `json:"maxmind_license_key"`
	GeoipDbInstalled            string           `json:"geoip_db_installed"`
	GeoipDbUpdateDays           string           `json:"geoip_db_update_days"`
	LaunchStartup               bool             `json:"launch_startup"`
	SysTrayIcon                 bool             `json:"sys_tray_icon"`
	GetFileSizesHold            GetFileSizesHold `json:"get_file_sizes_hold"`
	InterfaceList               []string         `json:"interface_list"`
}
type Slack struct {
	SlackOnStop       bool   `json:"slack_on_stop"`
	SlackOnWatched    bool   `json:"slack_on_watched"`
	SlackOnExtUp      bool   `json:"slack_on_extup"`
	SlackOnResume     bool   `json:"slack_on_resume"`
	SlackHook         string `json:"slack_hook"`
	SlackInclSubject  bool   `json:"slack_incl_subject"`
	SlackOnPMSUpdate  bool   `json:"slack_on_pmsupdate"`
	SlackChannel      string `json:"slack_channel"`
	SlackOnBuffer     bool   `json:"slack_on_buffer"`
	SlackOnExtdown    bool   `json:"slack_on_extdown"`
	SlackUsername     string `json:"slack_username"`
	SlackInclPMSlink  bool   `json:"slack_incl_pmslink"`
	SlackEnabled      bool   `json:"slack_enabled"`
	SlackInclPoster   bool   `json:"slack_incl_poster"`
	SlackOnPause      bool   `json:"slack_on_pause"`
	SlackOnConcurrent bool   `json:"slack_on_concurrent"`
	SlackOnIntUp      bool   `json:"slack_on_intup"`
	SlackOnIntDown    bool   `json:"slack_on_intdown"`
	SlackOnCreated    bool   `json:"slack_on_created"`
	SlackOnPlay       bool   `json:"slack_on_play"`
	SlackOnNewDevice  bool   `json:"slack_on_newdevice"`
	SlackIconEmoji    string `json:"slack_icon_emoji"`
}
type Boxcar struct {
	BoxcarSound        string `json:"boxcar_sound"`
	BoxcarOnNewDevice  bool   `json:"boxcar_on_newdevice"`
	BoxcarOnWatched    bool   `json:"boxcar_on_watched"`
	BoxcarOnBuffer     bool   `json:"boxcar_on_buffer"`
	BoxcarOnStop       bool   `json:"boxcar_on_stop"`
	BoxcarOnExtdown    bool   `json:"boxcar_on_extdown"`
	BoxcarOnPMSUpdate  bool   `json:"boxcar_on_pmsupdate"`
	BoxcarOnIntDown    bool   `json:"boxcar_on_intdown"`
	BoxcarOnPause      bool   `json:"boxcar_on_pause"`
	BoxcarOnPlay       bool   `json:"boxcar_on_play"`
	BoxcarOnCreated    bool   `json:"boxcar_on_created"`
	BoxcarOnIntUp      bool   `json:"boxcar_on_intup"`
	BoxcarOnResume     bool   `json:"boxcar_on_resume"`
	BoxcarOnConcurrent bool   `json:"boxcar_on_concurrent"`
	BoxcarEnabled      bool   `json:"boxcar_enabled"`
	BoxcarOnExtUp      bool   `json:"boxcar_on_extup"`
	BoxcarToken        string `json:"boxcar_token"`
}
type PushBullet struct {
	PushbulletOnPlay       bool   `json:"pushbullet_on_play"`
	PushbulletChannelTag   string `json:"pushbullet_channel_tag"`
	PushbulletOnBuffer     bool   `json:"pushbullet_on_buffer"`
	PushbulletAPIKey       string `json:"pushbullet_apikey"`
	PushbulletOnPMSUpdate  bool   `json:"pushbullet_on_pmsupdate"`
	PushbulletOnResume     bool   `json:"pushbullet_on_resume"`
	PushbulletOnConcurrent bool   `json:"pushbullet_on_concurrent"`
	PushbulletOnCreated    bool   `json:"pushbullet_on_created"`
	PushbulletOnPause      bool   `json:"pushbullet_on_pause"`
	PushbulletOnExtdown    bool   `json:"pushbullet_on_extdown"`
	PushbulletOnStop       bool   `json:"pushbullet_on_stop"`
	PushbulletOnExtUp      bool   `json:"pushbullet_on_extup"`
	PushbulletEnabled      bool   `json:"pushbullet_enabled"`
	PushbulletOnWatched    bool   `json:"pushbullet_on_watched"`
	PushbulletOnIntDown    bool   `json:"pushbullet_on_intdown"`
	PushbulletDeviceid     string `json:"pushbullet_deviceid"`
	PushbulletOnIntUp      bool   `json:"pushbullet_on_intup"`
	PushbulletOnNewDevice  bool   `json:"pushbullet_on_newdevice"`
}
type Scripts struct {
	ScriptsOnNewDevice        bool   `json:"scripts_on_newdevice"`
	ScriptsOnExtdownScript    string `json:"scripts_on_extdown_script"`
	ScriptsOnIntUpScript      string `json:"scripts_on_intup_script"`
	ScriptsOnStop             bool   `json:"scripts_on_stop"`
	ScriptsOnIntDownScript    string `json:"scripts_on_intdown_script"`
	ScriptsOnNewDeviceScript  string `json:"scripts_on_newdevice_script"`
	ScriptsFolder             string `json:"scripts_folder"`
	ScriptsOnPMSUpdate        bool   `json:"scripts_on_pmsupdate"`
	ScriptsOnBuffer           bool   `json:"scripts_on_buffer"`
	ScriptsOnConcurrent       bool   `json:"scripts_on_concurrent"`
	ScriptsOnIntDown          bool   `json:"scripts_on_intdown"`
	ScriptsOnResume           bool   `json:"scripts_on_resume"`
	ScriptsOnExtUp            bool   `json:"scripts_on_extup"`
	ScriptsOnPlayScript       string `json:"scripts_on_play_script"`
	ScriptsOnCreatedScript    string `json:"scripts_on_created_script"`
	ScriptsOnPMSUpdateScript  string `json:"scripts_on_pmsupdate_script"`
	ScriptsOnResumeScript     string `json:"scripts_on_resume_script"`
	ScriptsOnStopScript       string `json:"scripts_on_stop_script"`
	ScriptsOnWatchedScript    string `json:"scripts_on_watched_script"`
	ScriptsEnabled            bool   `json:"scripts_enabled"`
	ScriptsOnExtUpScript      string `json:"scripts_on_extup_script"`
	ScriptsOnExtdown          bool   `json:"scripts_on_extdown"`
	ScriptsOnCreated          bool   `json:"scripts_on_created"`
	ScriptsOnConcurrentScript string `json:"scripts_on_concurrent_script"`
	ScriptsOnPauseScript      string `json:"scripts_on_pause_script"`
	ScriptsOnPause            bool   `json:"scripts_on_pause"`
	ScriptsOnIntUp            bool   `json:"scripts_on_intup"`
	ScriptsTimeout            string `json:"scripts_timeout"`
	ScriptsOnWatched          bool   `json:"scripts_on_watched"`
	ScriptsOnBufferScript     string `json:"scripts_on_buffer_script"`
	ScriptsOnPlay             bool   `json:"scripts_on_play"`
}
type OSXNotify struct {
	OsxNotifyOnPause      bool   `json:"osx_notify_on_pause"`
	OsxNotifyOnIntDown    bool   `json:"osx_notify_on_intdown"`
	OsxNotifyOnBuffer     bool   `json:"osx_notify_on_buffer"`
	OsxNotifyApp          string `json:"osx_notify_app"`
	OsxNotifyOnCreated    bool   `json:"osx_notify_on_created"`
	OsxNotifyOnPMSUpdate  bool   `json:"osx_notify_on_pmsupdate"`
	OsxNotifyOnStop       bool   `json:"osx_notify_on_stop"`
	OsxNotifyEnabled      bool   `json:"osx_notify_enabled"`
	OsxNotifyOnIntUp      bool   `json:"osx_notify_on_intup"`
	OsxNotifyOnWatched    bool   `json:"osx_notify_on_watched"`
	OsxNotifyOnConcurrent bool   `json:"osx_notify_on_concurrent"`
	OsxNotifyOnNewDevice  bool   `json:"osx_notify_on_newdevice"`
	OsxNotifyOnPlay       bool   `json:"osx_notify_on_play"`
	OsxNotifyOnExtdown    bool   `json:"osx_notify_on_extdown"`
	OsxNotifyOnExtUp      bool   `json:"osx_notify_on_extup"`
	OsxNotifyOnResume     bool   `json:"osx_notify_on_resume"`
}
type NMA struct {
	NmaOnPMSUpdate  bool   `json:"nma_on_pmsupdate"`
	NmaPriority     bool   `json:"nma_priority"`
	NmaOnIntDown    bool   `json:"nma_on_intdown"`
	NmaOnBuffer     bool   `json:"nma_on_buffer"`
	NmaOnPlay       bool   `json:"nma_on_play"`
	NmaOnNewDevice  bool   `json:"nma_on_newdevice"`
	NmaOnConcurrent bool   `json:"nma_on_concurrent"`
	NmaOnResume     bool   `json:"nma_on_resume"`
	NmaOnWatched    bool   `json:"nma_on_watched"`
	NmaAPIKey       string `json:"nma_apikey"`
	NmaOnIntUp      bool   `json:"nma_on_intup"`
	NmaOnStop       bool   `json:"nma_on_stop"`
	NmaOnCreated    bool   `json:"nma_on_created"`
	NmaOnPause      bool   `json:"nma_on_pause"`
	NmaEnabled      bool   `json:"nma_enabled"`
	NmaOnExtdown    bool   `json:"nma_on_extdown"`
	NmaOnExtUp      bool   `json:"nma_on_extup"`
}
type Hipchat struct {
	HipchatEmoticon     string `json:"hipchat_emoticon"`
	HipchatInclSubject  bool   `json:"hipchat_incl_subject"`
	HipchatOnExtdown    bool   `json:"hipchat_on_extdown"`
	HipchatOnIntUp      bool   `json:"hipchat_on_intup"`
	HipchatEnabled      bool   `json:"hipchat_enabled"`
	HipchatOnCreated    bool   `json:"hipchat_on_created"`
	HipchatOnNewDevice  bool   `json:"hipchat_on_newdevice"`
	HipchatOnPlay       bool   `json:"hipchat_on_play"`
	HipchatOnResume     bool   `json:"hipchat_on_resume"`
	HipchatOnBuffer     bool   `json:"hipchat_on_buffer"`
	HipchatInclPoster   bool   `json:"hipchat_incl_poster"`
	HipchatOnWatched    bool   `json:"hipchat_on_watched"`
	HipchatColor        string `json:"hipchat_color"`
	HipchatOnConcurrent bool   `json:"hipchat_on_concurrent"`
	HipchatURL          string `json:"hipchat_url"`
	HipchatOnPMSUpdate  bool   `json:"hipchat_on_pmsupdate"`
	HipchatOnExtUp      bool   `json:"hipchat_on_extup"`
	HipchatInclPMSlink  bool   `json:"hipchat_incl_pmslink"`
	HipchatOnPause      bool   `json:"hipchat_on_pause"`
	HipchatOnIntDown    bool   `json:"hipchat_on_intdown"`
	HipchatOnStop       bool   `json:"hipchat_on_stop"`
}
type Pushalot struct {
	PushalotAPIKey       string `json:"pushalot_apikey"`
	PushalotOnPMSUpdate  bool   `json:"pushalot_on_pmsupdate"`
	PushalotEnabled      bool   `json:"pushalot_enabled"`
	PushalotOnExtdown    bool   `json:"pushalot_on_extdown"`
	PushalotOnCreated    bool   `json:"pushalot_on_created"`
	PushalotOnPause      bool   `json:"pushalot_on_pause"`
	PushalotOnStop       bool   `json:"pushalot_on_stop"`
	PushalotOnBuffer     bool   `json:"pushalot_on_buffer"`
	PushalotOnIntUp      bool   `json:"pushalot_on_intup"`
	PushalotOnConcurrent bool   `json:"pushalot_on_concurrent"`
	PushalotOnNewDevice  bool   `json:"pushalot_on_newdevice"`
	PushalotOnWatched    bool   `json:"pushalot_on_watched"`
	PushalotOnPlay       bool   `json:"pushalot_on_play"`
	PushalotOnExtUp      bool   `json:"pushalot_on_extup"`
	PushalotOnIntDown    bool   `json:"pushalot_on_intdown"`
	PushalotOnResume     bool   `json:"pushalot_on_resume"`
}
type Prowl struct {
	ProwlOnCreated    bool   `json:"prowl_on_created"`
	ProwlOnPlay       bool   `json:"prowl_on_play"`
	ProwlOnExtdown    bool   `json:"prowl_on_extdown"`
	ProwlOnPause      bool   `json:"prowl_on_pause"`
	ProwlOnIntDown    bool   `json:"prowl_on_intdown"`
	ProwlOnIntUp      bool   `json:"prowl_on_intup"`
	ProwlEnabled      bool   `json:"prowl_enabled"`
	ProwlPriority     bool   `json:"prowl_priority"`
	ProwlOnPMSUpdate  bool   `json:"prowl_on_pmsupdate"`
	ProwlOnNewDevice  bool   `json:"prowl_on_newdevice"`
	ProwlOnExtUp      bool   `json:"prowl_on_extup"`
	ProwlOnWatched    bool   `json:"prowl_on_watched"`
	ProwlOnBuffer     bool   `json:"prowl_on_buffer"`
	ProwlOnResume     bool   `json:"prowl_on_resume"`
	ProwlOnStop       bool   `json:"prowl_on_stop"`
	ProwlOnConcurrent bool   `json:"prowl_on_concurrent"`
	ProwlKeys         string `json:"prowl_keys"`
}
type Advanced struct {
	VerifySSLCert               bool   `json:"verify_ssl_cert"`
	SessionDbWriteAttempts      string `json:"session_db_write_attempts"`
	SystemAnalytics             bool   `json:"system_analytics"`
	JournalMode                 string `json:"journal_mode"`
	NotificationThreads         string `json:"notification_threads"`
	RemoteAccessPingThreshold   string `json:"remote_access_ping_threshold"`
	WebsocketConnectionTimeout  string `json:"websocket_connection_timeout"`
	ConfigVersion               string `json:"config_version"`
	MetadataCacheSeconds        string `json:"metadata_cache_seconds"`
	PMSTimeout                  string `json:"pms_timeout"`
	WebsocketMonitorPingPong    bool   `json:"websocket_monitor_ping_pong"`
	WebsocketConnectionAttempts string `json:"websocket_connection_attempts"`
	JwtSecret                   string `json:"jwt_secret"`
	CacheSizeMB                 string `json:"cache_sizemb"`
	SynchronousMode             string `json:"synchronous_mode"`
	PMSUpdateCheckInterval      string `json:"pms_update_check_interval"`
	JwtUpdateSecret             bool   `json:"jwt_update_secret"`
	VerboseLogs                 bool   `json:"verbose_logs"`
	AddLiveTvLibrary            bool   `json:"add_live_tv_library"`
	RemoteAccessPingInterval    string `json:"remote_access_ping_interval"`
}
type Facebook struct {
	FacebookAppSecret    string `json:"facebook_app_secret"`
	FacebookOnStop       bool   `json:"facebook_on_stop"`
	FacebookToken        string `json:"facebook_token"`
	FacebookOnBuffer     bool   `json:"facebook_on_buffer"`
	FacebookOnPlay       bool   `json:"facebook_on_play"`
	FacebookEnabled      bool   `json:"facebook_enabled"`
	FacebookOnIntUp      bool   `json:"facebook_on_intup"`
	FacebookInclPoster   bool   `json:"facebook_incl_poster"`
	FacebookOnPMSUpdate  bool   `json:"facebook_on_pmsupdate"`
	FacebookOnIntDown    bool   `json:"facebook_on_intdown"`
	FacebookOnExtUp      bool   `json:"facebook_on_extup"`
	FacebookOnResume     bool   `json:"facebook_on_resume"`
	FacebookOnConcurrent bool   `json:"facebook_on_concurrent"`
	FacebookOnExtdown    bool   `json:"facebook_on_extdown"`
	FacebookGroup        string `json:"facebook_group"`
	FacebookOnPause      bool   `json:"facebook_on_pause"`
	FacebookAppID        string `json:"facebook_app_id"`
	FacebookInclPMSlink  bool   `json:"facebook_incl_pmslink"`
	FacebookOnNewDevice  bool   `json:"facebook_on_newdevice"`
	FacebookOnWatched    bool   `json:"facebook_on_watched"`
	FacebookRedirectURI  string `json:"facebook_redirect_uri"`
	FacebookInclSubject  bool   `json:"facebook_incl_subject"`
	FacebookOnCreated    bool   `json:"facebook_on_created"`
}
type Email struct {
	EmailFromName     string `json:"email_from_name"`
	EmailBcc          string `json:"email_bcc"`
	EmailOnPause      bool   `json:"email_on_pause"`
	EmailSMTPPassword string `json:"email_smtp_password"`
	EmailTLS          bool   `json:"email_tls"`
	EmailSMTPPort     string `json:"email_smtp_port"`
	EmailOnWatched    bool   `json:"email_on_watched"`
	EmailSMTPServer   string `json:"email_smtp_server"`
	EmailOnIntDown    bool   `json:"email_on_intdown"`
	EmailOnConcurrent bool   `json:"email_on_concurrent"`
	EmailOnIntUp      bool   `json:"email_on_intup"`
	EmailEnabled      bool   `json:"email_enabled"`
	EmailOnBuffer     bool   `json:"email_on_buffer"`
	EmailOnResume     bool   `json:"email_on_resume"`
	EmailSMTPUser     string `json:"email_smtp_user"`
	EmailCc           string `json:"email_cc"`
	EmailOnNewDevice  bool   `json:"email_on_newdevice"`
	EmailOnStop       bool   `json:"email_on_stop"`
	EmailOnExtdown    bool   `json:"email_on_extdown"`
	EmailHTMLSupport  bool   `json:"email_html_support"`
	EmailOnExtUp      bool   `json:"email_on_extup"`
	EmailOnCreated    bool   `json:"email_on_created"`
	EmailOnPMSUpdate  bool   `json:"email_on_pmsupdate"`
	EmailTo           string `json:"email_to"`
	EmailFrom         string `json:"email_from"`
	EmailOnPlay       bool   `json:"email_on_play"`
}
type PlexWatch struct {
	GroupingUserHistory   bool   `json:"grouping_user_history"`
	PlexwatchDatabase     string `json:"plexwatch_database"`
	GroupingGlobalHistory bool   `json:"grouping_global_history"`
	GroupingCharts        bool   `json:"grouping_charts"`
}
type Pushover struct {
	PushoverEnabled      bool   `json:"pushover_enabled"`
	PushoverOnExtUp      bool   `json:"pushover_on_extup"`
	PushoverOnIntUp      bool   `json:"pushover_on_intup"`
	PushoverOnBuffer     bool   `json:"pushover_on_buffer"`
	PushoverHTMLSupport  bool   `json:"pushover_html_support"`
	PushoverKeys         string `json:"pushover_keys"`
	PushoverInclURL      bool   `json:"pushover_incl_url"`
	PushoverOnPause      bool   `json:"pushover_on_pause"`
	PushoverOnPlay       bool   `json:"pushover_on_play"`
	PushoverOnStop       bool   `json:"pushover_on_stop"`
	PushoverApitoken     string `json:"pushover_apitoken"`
	PushoverInclPMSlink  bool   `json:"pushover_incl_pmslink"`
	PushoverOnPMSUpdate  bool   `json:"pushover_on_pmsupdate"`
	PushoverOnExtdown    bool   `json:"pushover_on_extdown"`
	PushoverPriority     bool   `json:"pushover_priority"`
	PushoverOnIntDown    bool   `json:"pushover_on_intdown"`
	PushoverOnNewDevice  bool   `json:"pushover_on_newdevice"`
	PushoverSound        string `json:"pushover_sound"`
	PushoverOnCreated    bool   `json:"pushover_on_created"`
	PushoverOnResume     bool   `json:"pushover_on_resume"`
	PushoverOnConcurrent bool   `json:"pushover_on_concurrent"`
	PushoverOnWatched    bool   `json:"pushover_on_watched"`
}
type Telegram struct {
	TelegramOnBuffer          bool   `json:"telegram_on_buffer"`
	TelegramBotToken          string `json:"telegram_bot_token"`
	TelegramInclPoster        bool   `json:"telegram_incl_poster"`
	TelegramOnResume          bool   `json:"telegram_on_resume"`
	TelegramOnPMSUpdate       bool   `json:"telegram_on_pmsupdate"`
	TelegramHTMLSupport       bool   `json:"telegram_html_support"`
	TelegramInclSubject       bool   `json:"telegram_incl_subject"`
	TelegramDisableWebPreview bool   `json:"telegram_disable_web_preview"`
	TelegramEnabled           bool   `json:"telegram_enabled"`
	TelegramOnConcurrent      bool   `json:"telegram_on_concurrent"`
	TelegramOnIntDown         bool   `json:"telegram_on_intdown"`
	TelegramOnExtdown         bool   `json:"telegram_on_extdown"`
	TelegramOnPause           bool   `json:"telegram_on_pause"`
	TelegramOnNewDevice       bool   `json:"telegram_on_newdevice"`
	TelegramOnWatched         bool   `json:"telegram_on_watched"`
	TelegramOnStop            bool   `json:"telegram_on_stop"`
	TelegramOnPlay            bool   `json:"telegram_on_play"`
	TelegramOnCreated         bool   `json:"telegram_on_created"`
	TelegramOnExtUp           bool   `json:"telegram_on_extup"`
	TelegramChatID            string `json:"telegram_chat_id"`
	TelegramOnIntUp           bool   `json:"telegram_on_intup"`
}
type Twitter struct {
	TwitterOnResume          bool   `json:"twitter_on_resume"`
	TwitterOnPMSUpdate       bool   `json:"twitter_on_pmsupdate"`
	TwitterOnBuffer          bool   `json:"twitter_on_buffer"`
	TwitterOnConcurrent      bool   `json:"twitter_on_concurrent"`
	TwitterEnabled           bool   `json:"twitter_enabled"`
	TwitterOnWatched         bool   `json:"twitter_on_watched"`
	TwitterAccessToken       string `json:"twitter_access_token"`
	TwitterConsumerSecret    string `json:"twitter_consumer_secret"`
	TwitterOnExtUp           bool   `json:"twitter_on_extup"`
	TwitterOnIntDown         bool   `json:"twitter_on_intdown"`
	TwitterOnIntUp           bool   `json:"twitter_on_intup"`
	TwitterAccessTokenSecret string `json:"twitter_access_token_secret"`
	TwitterOnNewDevice       bool   `json:"twitter_on_newdevice"`
	TwitterOnStop            bool   `json:"twitter_on_stop"`
	TwitterConsumerKey       string `json:"twitter_consumer_key"`
	TwitterInclSubject       bool   `json:"twitter_incl_subject"`
	TwitterOnExtdown         bool   `json:"twitter_on_extdown"`
	TwitterInclPoster        bool   `json:"twitter_incl_poster"`
	TwitterOnPause           bool   `json:"twitter_on_pause"`
	TwitterOnPlay            bool   `json:"twitter_on_play"`
	TwitterOnCreated         bool   `json:"twitter_on_created"`
}
type Plex struct {
	PlexOnCreated    bool   `json:"plex_on_created"`
	PlexOnConcurrent bool   `json:"plex_on_concurrent"`
	PlexClientHost   string `json:"plex_client_host"`
	PlexOnStop       bool   `json:"plex_on_stop"`
	PlexOnResume     bool   `json:"plex_on_resume"`
	PlexOnWatched    bool   `json:"plex_on_watched"`
	PlexOnPlay       bool   `json:"plex_on_play"`
	PlexOnExtdown    bool   `json:"plex_on_extdown"`
	PlexOnIntDown    bool   `json:"plex_on_intdown"`
	PlexOnNewDevice  bool   `json:"plex_on_newdevice"`
	PlexEnabled      bool   `json:"plex_enabled"`
	PlexOnBuffer     bool   `json:"plex_on_buffer"`
	PlexOnIntUp      bool   `json:"plex_on_intup"`
	PlexOnPMSUpdate  bool   `json:"plex_on_pmsupdate"`
	PlexUsername     string `json:"plex_username"`
	PlexOnPause      bool   `json:"plex_on_pause"`
	PlexPassword     string `json:"plex_password"`
	PlexOnExtUp      bool   `json:"plex_on_extup"`
}
type Join struct {
	JoinOnConcurrent bool   `json:"join_on_concurrent"`
	JoinOnNewDevice  bool   `json:"join_on_newdevice"`
	JoinDeviceID     string `json:"join_deviceid"`
	JoinOnPlay       bool   `json:"join_on_play"`
	JoinOnWatched    bool   `json:"join_on_watched"`
	JoinOnResume     bool   `json:"join_on_resume"`
	JoinOnCreated    bool   `json:"join_on_created"`
	JoinOnExtdown    bool   `json:"join_on_extdown"`
	JoinOnIntUp      bool   `json:"join_on_intup"`
	JoinOnBuffer     bool   `json:"join_on_buffer"`
	JoinInclSubject  bool   `json:"join_incl_subject"`
	JoinEnabled      bool   `json:"join_enabled"`
	JoinOnPause      bool   `json:"join_on_pause"`
	JoinAPIKey       string `json:"join_apikey"`
	JoinOnPMSUpdate  bool   `json:"join_on_pmsupdate"`
	JoinOnIntDown    bool   `json:"join_on_intdown"`
	JoinOnExtUp      bool   `json:"join_on_extup"`
	JoinOnStop       bool   `json:"join_on_stop"`
}
type XBMC struct {
	XBMCPassword     string `json:"xbmc_password"`
	XBMCOnPlay       bool   `json:"xbmc_on_play"`
	XBMCOnPause      bool   `json:"xbmc_on_pause"`
	XBMCUsername     string `json:"xbmc_username"`
	XBMCOnExtUp      bool   `json:"xbmc_on_extup"`
	XBMCOnWatched    bool   `json:"xbmc_on_watched"`
	XBMCOnBuffer     bool   `json:"xbmc_on_buffer"`
	XBMCEnabled      bool   `json:"xbmc_enabled"`
	XBMCHost         string `json:"xbmc_host"`
	XBMCOnCreated    bool   `json:"xbmc_on_created"`
	XBMCOnIntUp      bool   `json:"xbmc_on_intup"`
	XBMCOnExtdown    bool   `json:"xbmc_on_extdown"`
	XBMCOnPMSUpdate  bool   `json:"xbmc_on_pmsupdate"`
	XBMCOnStop       bool   `json:"xbmc_on_stop"`
	XBMCOnNewDevice  bool   `json:"xbmc_on_newdevice"`
	XBMCOnResume     bool   `json:"xbmc_on_resume"`
	XBMCOnIntDown    bool   `json:"xbmc_on_intdown"`
	XBMCOnConcurrent bool   `json:"xbmc_on_concurrent"`
}
type Newsletter struct {
	NewsletterDir          string `json:"newsletter_dir"`
	NewsletterStaticURL    bool   `json:"newsletter_static_url"`
	NewsletterTemplates    string `json:"newsletter_templates"`
	NewsletterSelfHosted   bool   `json:"newsletter_self_hosted"`
	NewsletterInlineStyles bool   `json:"newsletter_inline_styles"`
	NewsletterPassword     string `json:"newsletter_password"`
	NewsletterCustomDir    string `json:"newsletter_custom_dir"`
	NewsletterAuth         string `json:"newsletter_auth"`
}
type Browser struct {
	BrowserOnIntDown     bool   `json:"browser_on_intdown"`
	BrowserOnNewDevice   bool   `json:"browser_on_newdevice"`
	BrowserOnPMSUpdate   bool   `json:"browser_on_pmsupdate"`
	BrowserOnIntUp       bool   `json:"browser_on_intup"`
	BrowserOnBuffer      bool   `json:"browser_on_buffer"`
	BrowserAutoHideDelay string `json:"browser_auto_hide_delay"`
	BrowserOnCreated     bool   `json:"browser_on_created"`
	BrowserOnPause       bool   `json:"browser_on_pause"`
	BrowserOnStop        bool   `json:"browser_on_stop"`
	BrowserOnPlay        bool   `json:"browser_on_play"`
	BrowserOnWatched     bool   `json:"browser_on_watched"`
	BrowserOnExtdown     bool   `json:"browser_on_extdown"`
	BrowserOnExtUp       bool   `json:"browser_on_extup"`
	BrowserOnResume      bool   `json:"browser_on_resume"`
	BrowserOnConcurrent  bool   `json:"browser_on_concurrent"`
	BrowserEnabled       bool   `json:"browser_enabled"`
}
type Cloudinary struct {
	CloudinaryAPIKey    string `json:"cloudinary_api_key"`
	CloudinaryCloudName string `json:"cloudinary_cloud_name"`
	CloudinaryAPISecret string `json:"cloudinary_api_secret"`
}
