package hahow

import "time"

type CoueseInfo struct {
	SuccessCriteria struct {
		NumSoldTickets int `json:"numSoldTickets"`
	} `json:"successCriteria"`
	IncubationSchedule struct {
		Owner         string    `json:"owner"`
		PreferredTime time.Time `json:"preferredTime"`
	} `json:"incubationSchedule"`
	Campaign struct {
		Types []any `json:"types"`
	} `json:"campaign"`
	ContentRating struct {
		Brief   string `json:"brief"`
		Content string `json:"content"`
	} `json:"contentRating"`
	NumSoldTickets                     int      `json:"numSoldTickets"`
	Status                             string   `json:"status"`
	Reviewing                          bool     `json:"reviewing"`
	Price                              int      `json:"price"`
	PreOrderedPrice                    int      `json:"preOrderedPrice"`
	EstimatedCourseStartIntervalInDays int      `json:"estimatedCourseStartIntervalInDays"`
	EstimatedCourseVideoLengthInMins   int      `json:"estimatedCourseVideoLengthInMins"`
	IsReject                           bool     `json:"isReject"`
	SkipIncubating                     bool     `json:"skipIncubating"`
	AverageRating                      float64  `json:"averageRating"`
	NumRating                          int      `json:"numRating"`
	BookmarkCount                      int      `json:"bookmarkCount"`
	MigratedToWistia                   bool     `json:"migratedToWistia"`
	WistiaStatus                       string   `json:"wistiaStatus"`
	VideoBackupStatus                  string   `json:"videoBackupStatus"`
	Tags                               []string `json:"tags"`
	ViewCount                          int      `json:"viewCount"`
	IsTop10Percent                     bool     `json:"isTop10Percent"`
	ID                                 string   `json:"_id"`
	Owner                              struct {
		Privacies struct {
			IsShowTaughtCourses              bool `json:"isShowTaughtCourses"`
			IsShowBoughtCourses              bool `json:"isShowBoughtCourses"`
			IsShowBoughtArticles             bool `json:"isShowBoughtArticles"`
			IsShowBoughtKnowledgeCollections bool `json:"isShowBoughtKnowledgeCollections"`
			IsShowBoughtLiveEvents           bool `json:"isShowBoughtLiveEvents"`
			IsShowIdeas                      bool `json:"isShowIdeas"`
			IsShowBookmarkedCourses          bool `json:"isShowBookmarkedCourses"`
			IsShowBookmarkedArticles         bool `json:"isShowBookmarkedArticles"`
			IsShowBookmarkedIdeas            bool `json:"isShowBookmarkedIdeas"`
			IsShowCreations                  bool `json:"isShowCreations"`
			IsShowCertificates               bool `json:"isShowCertificates"`
		} `json:"privacies"`
		Links struct {
			Facebook string `json:"facebook"`
			Youtube  string `json:"youtube"`
			Website  string `json:"website"`
		} `json:"links"`
		States struct {
			Club2020 struct {
				IsShared            bool `json:"isShared"`
				IsSharedAfterBought bool `json:"isSharedAfterBought"`
			} `json:"club2020"`
			AcknowledgedContractRenewal2019 string `json:"acknowledgedContractRenewal2019"`
			TeacherTypeLock                 bool   `json:"teacherTypeLock"`
			HaveNotifiedStudyPlanFeature    bool   `json:"haveNotifiedStudyPlanFeature"`
		} `json:"states"`
		ID                         string `json:"_id"`
		Username                   string `json:"username"`
		Name                       string `json:"name"`
		ProfileImageURL            string `json:"profileImageUrl"`
		CoverImageURL              string `json:"coverImageUrl"`
		MetaDescription            string `json:"metaDescription"`
		Skills                     string `json:"skills"`
		PublishedProductStatistics struct {
			CourseCount  int `json:"courseCount"`
			ArticleCount int `json:"articleCount"`
			StudentCount int `json:"studentCount"`
		} `json:"publishedProductStatistics"`
	} `json:"owner"`
	DraftRevisionUpdateTime     time.Time `json:"draftRevisionUpdateTime"`
	Uniquename                  string    `json:"uniquename"`
	RevisionReleaseTime         time.Time `json:"revisionReleaseTime"`
	EstimatedCourseStartTime    time.Time `json:"estimatedCourseStartTime"`
	IncubateTime                time.Time `json:"incubateTime"`
	LatestCourseStartTime       time.Time `json:"latestCourseStartTime"`
	ProposalDueTime             time.Time `json:"proposalDueTime"`
	PublishTime                 time.Time `json:"publishTime"`
	CreatedAt                   time.Time `json:"createdAt"`
	PriceInMoneyPoint           int       `json:"priceInMoneyPoint"`
	PreOrderedPriceInMoneyPoint int       `json:"preOrderedPriceInMoneyPoint"`
	I18N                        struct {
		Description           string `json:"description"`
		MetaDescription       string `json:"metaDescription"`
		RecommendedBackground string `json:"recommendedBackground"`
		RequiredTools         string `json:"requiredTools"`
		TargetGroup           string `json:"targetGroup"`
		Title                 string `json:"title"`
		WillLearn             string `json:"willLearn"`
	} `json:"i18n"`
	Modules               []string `json:"modules"`
	Description           string   `json:"description"`
	Title                 string   `json:"title"`
	MetaDescription       string   `json:"metaDescription"`
	RequiredTools         string   `json:"requiredTools"`
	RecommendedBackground string   `json:"recommendedBackground"`
	WillLearn             string   `json:"willLearn"`
	TargetGroup           string   `json:"targetGroup"`
	CoverImage            struct {
		ID     string `json:"_id"`
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"coverImage"`
	SquareCoverImage struct {
		ID     string `json:"_id"`
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"squareCoverImage"`
	Video struct {
		ID               string `json:"_id"`
		PreviewImageUrls struct {
			Source string `json:"source"`
			Images struct {
				DimensionW300  string `json:"DIMENSION_W300"`
				DimensionW600  string `json:"DIMENSION_W600"`
				DimensionW1000 string `json:"DIMENSION_W1000"`
			} `json:"images"`
		} `json:"previewImageUrls"`
		Videos []struct {
			Quality string `json:"quality"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Size    int    `json:"size"`
			Link    string `json:"link"`
		} `json:"videos"`
		VideoLengthInSeconds int   `json:"videoLengthInSeconds"`
		Subtitles            []any `json:"subtitles"`
	} `json:"video"`
	TotalVideoLengthInSeconds int `json:"totalVideoLengthInSeconds"`
	ChargePlan                struct {
		Type             string   `json:"type"`
		PaymentPlanTypes []string `json:"paymentPlanTypes"`
		Contracts        []struct {
			PaymentPlanTypes []string  `json:"paymentPlanTypes"`
			Hidden           bool      `json:"hidden"`
			ID               string    `json:"_id"`
			Title            string    `json:"title"`
			Type             string    `json:"type"`
			Status           string    `json:"status"`
			ChargePlanType   string    `json:"chargePlanType"`
			CreatedAt        time.Time `json:"createdAt"`
			URL              string    `json:"url"`
		} `json:"contracts"`
		ID        string    `json:"_id"`
		Course    string    `json:"course"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"chargePlan"`
	Group struct {
		ID         string `json:"_id"`
		Title      string `json:"title"`
		Uniquename string `json:"uniquename"`
		SubGroup   struct {
			ID         string `json:"_id"`
			Title      string `json:"title"`
			Uniquename string `json:"uniquename"`
		} `json:"subGroup"`
	} `json:"group"`
	IncludedKnowledgeCollections  []any   `json:"includedKnowledgeCollections"`
	GlobalAverageRating           float64 `json:"globalAverageRating"`
	HasSubscribedSaleNotification bool    `json:"hasSubscribedSaleNotification"`
	CurrencyCode                  string  `json:"currencyCode"`
	BasePricingInfo               struct {
		Price           int    `json:"price"`
		PreOrderedPrice int    `json:"preOrderedPrice"`
		CurrencyCode    string `json:"currencyCode"`
	} `json:"basePricingInfo"`
}

type Course []struct {
	ID   string `json:"_id"`
	I18N struct {
		Title string `json:"title"`
	} `json:"i18n"`
	Items []struct {
		ID            string `json:"_id"`
		Type          string `json:"type"`
		ChapterNumber string `json:"chapterNumber"`
		Content       struct {
			MigratedToWistia bool      `json:"migratedToWistia"`
			ID               string    `json:"_id"`
			Owner            string    `json:"owner"`
			Course           string    `json:"course"`
			ModuleItem       string    `json:"moduleItem"`
			UpdatedAt        time.Time `json:"updatedAt"`
			I18N             struct {
				Description string `json:"description"`
				Title       string `json:"title"`
			} `json:"i18n"`
			Assets []struct {
				ID          string `json:"_id"`
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
			} `json:"assets"`
			Permission  string `json:"permission"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Video       struct {
				ID                   string `json:"_id"`
				IsExisted            bool   `json:"isExisted"`
				VideoLengthInSeconds int    `json:"videoLengthInSeconds"`
				VideoThumbnailID     string `json:"videoThumbnailId"`
				PreviewImageUrls     struct {
					Source string `json:"source"`
					Images struct {
						DimensionW300  string `json:"DIMENSION_W300"`
						DimensionW600  string `json:"DIMENSION_W600"`
						DimensionW1000 string `json:"DIMENSION_W1000"`
					} `json:"images"`
				} `json:"previewImageUrls"`
			} `json:"video"`
		} `json:"content"`
	} `json:"items"`
	Title         string `json:"title"`
	ChapterNumber string `json:"chapterNumber"`
}

type LectureItems []struct {
	Items []struct {
		ID            string `json:"_id"`
		Type          string `json:"type"`
		ChapterNumber string `json:"chapterNumber"`
		Content       struct {
			MigratedToWistia bool      `json:"migratedToWistia"`
			ID               string    `json:"_id"`
			Owner            string    `json:"owner"`
			Course           string    `json:"course"`
			ModuleItem       string    `json:"moduleItem"`
			UpdatedAt        time.Time `json:"updatedAt"`
			I18N             struct {
				Description string `json:"description"`
				Title       string `json:"title"`
			} `json:"i18n"`
			Assets []struct {
				ID          string `json:"_id"`
				DisplayName string `json:"displayName"`
				URL         string `json:"url"`
			} `json:"assets"`
			Permission  string `json:"permission"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Video       struct {
				ID                   string `json:"_id"`
				IsExisted            bool   `json:"isExisted"`
				VideoLengthInSeconds int    `json:"videoLengthInSeconds"`
				VideoThumbnailID     string `json:"videoThumbnailId"`
				PreviewImageUrls     struct {
					Source string `json:"source"`
					Images struct {
						DimensionW300  string `json:"DIMENSION_W300"`
						DimensionW600  string `json:"DIMENSION_W600"`
						DimensionW1000 string `json:"DIMENSION_W1000"`
					} `json:"images"`
				} `json:"previewImageUrls"`
			} `json:"video"`
		} `json:"content"`
	} `json:"items"`
}
type Lecture struct {
	MigratedToWistia bool   `json:"migratedToWistia"`
	ID               string `json:"_id"`
	Owner            struct {
		Privacies struct {
			IsShowTaughtCourses              bool `json:"isShowTaughtCourses"`
			IsShowBoughtCourses              bool `json:"isShowBoughtCourses"`
			IsShowBoughtArticles             bool `json:"isShowBoughtArticles"`
			IsShowBoughtKnowledgeCollections bool `json:"isShowBoughtKnowledgeCollections"`
			IsShowBoughtLiveEvents           bool `json:"isShowBoughtLiveEvents"`
			IsShowIdeas                      bool `json:"isShowIdeas"`
			IsShowBookmarkedCourses          bool `json:"isShowBookmarkedCourses"`
			IsShowBookmarkedArticles         bool `json:"isShowBookmarkedArticles"`
			IsShowBookmarkedIdeas            bool `json:"isShowBookmarkedIdeas"`
			IsShowCreations                  bool `json:"isShowCreations"`
			IsShowCertificates               bool `json:"isShowCertificates"`
		} `json:"privacies"`
		Links struct {
			Facebook string `json:"facebook"`
			Youtube  string `json:"youtube"`
			Website  string `json:"website"`
		} `json:"links"`
		States struct {
			Club2020 struct {
				IsShared            bool `json:"isShared"`
				IsSharedAfterBought bool `json:"isSharedAfterBought"`
			} `json:"club2020"`
			AcknowledgedContractRenewal2019 string `json:"acknowledgedContractRenewal2019"`
			TeacherTypeLock                 bool   `json:"teacherTypeLock"`
			HaveNotifiedStudyPlanFeature    bool   `json:"haveNotifiedStudyPlanFeature"`
		} `json:"states"`
		ID              string `json:"_id"`
		Username        string `json:"username"`
		Name            string `json:"name"`
		ProfileImageURL string `json:"profileImageUrl"`
		CoverImageURL   string `json:"coverImageUrl"`
		MetaDescription string `json:"metaDescription"`
		Skills          string `json:"skills"`
	} `json:"owner"`
	Course     string    `json:"course"`
	ModuleItem string    `json:"moduleItem"`
	UpdatedAt  time.Time `json:"updatedAt"`
	I18N       struct {
		Description string `json:"description"`
		Title       string `json:"title"`
	} `json:"i18n"`
	Assets []struct {
		ID          string `json:"_id"`
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
	} `json:"assets"`
	Permission  string `json:"permission"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       struct {
		ID               string `json:"_id"`
		PreviewImageUrls struct {
			Source string `json:"source"`
			Images struct {
				DimensionW300  string `json:"DIMENSION_W300"`
				DimensionW600  string `json:"DIMENSION_W600"`
				DimensionW1000 string `json:"DIMENSION_W1000"`
			} `json:"images"`
		} `json:"previewImageUrls"`
		Videos []struct {
			Quality string    `json:"quality"`
			Width   int       `json:"width,omitempty"`
			Height  int       `json:"height,omitempty"`
			Size    int       `json:"size"`
			Link    string    `json:"link"`
			Expires time.Time `json:"expires"`
		} `json:"videos"`
		VideoLengthInSeconds int `json:"videoLengthInSeconds"`
		Subtitles            []struct {
			Link     string `json:"link"`
			Language string `json:"language"`
		} `json:"subtitles"`
		VideoThumbnailID string `json:"videoThumbnailId"`
	} `json:"video"`
}

type Config struct {
	Token            string `json:"TOKEN"`
	CourseID         string `json:"COURSE_ID"`
	ConcurrencyLimit int    `json:"ConcurrencyLimit"`
}
