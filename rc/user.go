package rc

// Generated by https://quicktype.io

type Me struct {
	ID               string       `json:"_id"`
	Name             string       `json:"name"`
	Emails           []Email      `json:"emails"`
	Status           string       `json:"status"`
	StatusConnection string       `json:"statusConnection"`
	Username         string       `json:"username"`
	UTCOffset        int64        `json:"utcOffset"`
	Active           bool         `json:"active"`
	Roles            []string     `json:"roles"`
	Settings         Settings     `json:"settings"`
	CustomFields     CustomFields `json:"customFields"`
	Success          bool         `json:"success"`
}

type CustomFields map[string]interface{}

type Email struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type Settings struct {
	Preferences Preferences `json:"preferences"`
}

type Preferences struct {
	EnableAutoAway              bool   `json:"enableAutoAway"`
	IdleTimeoutLimit            int64  `json:"idleTimeoutLimit"`
	DesktopNotificationDuration int64  `json:"desktopNotificationDuration"`
	AudioNotifications          string `json:"audioNotifications"`
	DesktopNotifications        string `json:"desktopNotifications"`
	MobileNotifications         string `json:"mobileNotifications"`
	UnreadAlert                 bool   `json:"unreadAlert"`
	UseEmojis                   bool   `json:"useEmojis"`
	ConvertASCIIEmoji           bool   `json:"convertAsciiEmoji"`
	AutoImageLoad               bool   `json:"autoImageLoad"`
	SaveMobileBandwidth         bool   `json:"saveMobileBandwidth"`
	CollapseMediaByDefault      bool   `json:"collapseMediaByDefault"`
	HideUsernames               bool   `json:"hideUsernames"`
	HideRoles                   bool   `json:"hideRoles"`
	HideFlexTab                 bool   `json:"hideFlexTab"`
	HideAvatars                 bool   `json:"hideAvatars"`
	RoomsListExhibitionMode     string `json:"roomsListExhibitionMode"`
	SidebarViewMode             string `json:"sidebarViewMode"`
	SidebarHideAvatar           bool   `json:"sidebarHideAvatar"`
	SidebarShowUnread           bool   `json:"sidebarShowUnread"`
	SidebarShowFavorites        bool   `json:"sidebarShowFavorites"`
	SendOnEnter                 string `json:"sendOnEnter"`
	MessageViewMode             int64  `json:"messageViewMode"`
	EmailNotificationMode       string `json:"emailNotificationMode"`
	RoomCounterSidebar          bool   `json:"roomCounterSidebar"`
	NewRoomNotification         string `json:"newRoomNotification"`
	NewMessageNotification      string `json:"newMessageNotification"`
	MuteFocusedConversations    bool   `json:"muteFocusedConversations"`
	NotificationsSoundVolume    int64  `json:"notificationsSoundVolume"`
}
