package services

import (
	"time"
)

type ConfigService struct {
	// Could be extended to use database or Redis for dynamic configs
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// FeatureFlags defines available features
type FeatureFlags struct {
	SubscriptionTracking     bool `json:"subscription_tracking"`
	AutoDetection           bool `json:"auto_detection"`
	BankConnection          bool `json:"bank_connection"`
	BudgetManagement        bool `json:"budget_management"`
	Analytics               bool `json:"analytics"`
	DataExport              bool `json:"data_export"`
	MultiAccount            bool `json:"multi_account"`
	Notifications           bool `json:"notifications"`
	DarkMode                bool `json:"dark_mode"`
	MultiLanguage           bool `json:"multi_language"`
	TwoFactorAuth           bool `json:"two_factor_auth"`
	PaymentReminders        bool `json:"payment_reminders"`
	SpendingInsights        bool `json:"spending_insights"`
	GoalSetting             bool `json:"goal_setting"`
	RecurringTransactions   bool `json:"recurring_transactions"`
	CategoryCustomization   bool `json:"category_customization"`
	ReportsGeneration       bool `json:"reports_generation"`
}

// MenuItem represents a navigation menu item
type MenuItem struct {
	ID          string      `json:"id"`
	Label       string      `json:"label"`
	Icon        string      `json:"icon"`
	Route       string      `json:"route"`
	Permission  string      `json:"permission,omitempty"`
	Badge       *Badge      `json:"badge,omitempty"`
	Children    []MenuItem  `json:"children,omitempty"`
	IsVisible   bool        `json:"is_visible"`
	IsPremium   bool        `json:"is_premium"`
	Order       int         `json:"order"`
}

// Badge for menu items (e.g., notification count)
type Badge struct {
	Value string `json:"value"`
	Type  string `json:"type"` // info, warning, error, success
}

// AppConfig represents the complete app configuration
type AppConfig struct {
	Features        FeatureFlags          `json:"features"`
	Menus           MenuConfig           `json:"menus"`
	Settings        AppSettings          `json:"settings"`
	Themes          []Theme              `json:"themes"`
	Languages       []Language           `json:"languages"`
	Categories      []Category           `json:"categories"`
	SubscriptionTypes []SubscriptionType `json:"subscription_types"`
	Currencies      []Currency           `json:"currencies"`
}

// MenuConfig contains all menu configurations
type MenuConfig struct {
	MainMenu     []MenuItem `json:"main_menu"`
	UserMenu     []MenuItem `json:"user_menu"`
	SettingsMenu []MenuItem `json:"settings_menu"`
	QuickActions []MenuItem `json:"quick_actions"`
}

// AppSettings contains global app settings
type AppSettings struct {
	AppName             string   `json:"app_name"`
	AppVersion          string   `json:"app_version"`
	APIVersion          string   `json:"api_version"`
	MaintenanceMode     bool     `json:"maintenance_mode"`
	MaintenanceMessage  string   `json:"maintenance_message"`
	MinClientVersion    string   `json:"min_client_version"`
	TermsURL            string   `json:"terms_url"`
	PrivacyURL          string   `json:"privacy_url"`
	SupportEmail        string   `json:"support_email"`
	SupportURL          string   `json:"support_url"`
	DefaultCurrency     string   `json:"default_currency"`
	DefaultLanguage     string   `json:"default_language"`
	DefaultTheme        string   `json:"default_theme"`
	DateFormat          string   `json:"date_format"`
	TimeFormat          string   `json:"time_format"`
	WeekStartsOn        int      `json:"week_starts_on"` // 0=Sunday, 1=Monday
	SessionTimeout      int      `json:"session_timeout"` // minutes
	MaxLoginAttempts    int      `json:"max_login_attempts"`
	PasswordMinLength   int      `json:"password_min_length"`
	AllowedFileTypes    []string `json:"allowed_file_types"`
	MaxFileSize         int64    `json:"max_file_size"` // bytes
}

// Theme represents a UI theme
type Theme struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	IsDark  bool             `json:"is_dark"`
	Colors  map[string]string `json:"colors"`
}

// Language represents a supported language
type Language struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	NativeName string `json:"native_name"`
	Direction string `json:"direction"` // ltr or rtl
	IsDefault bool   `json:"is_default"`
}

// Category for transactions
type Category struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Icon     string   `json:"icon"`
	Color    string   `json:"color"`
	Type     string   `json:"type"` // income, expense, both
	ParentID *string  `json:"parent_id,omitempty"`
}

// SubscriptionType represents billing cycles
type SubscriptionType struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Days  int    `json:"days"`
}

// Currency represents supported currencies
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rate   float64 `json:"rate"` // Exchange rate to EUR
}

// GetFeatureFlags returns feature flags for the user
func (s *ConfigService) GetFeatureFlags(userType string, isPremium bool) FeatureFlags {
	// Base features for all users
	features := FeatureFlags{
		SubscriptionTracking:   true,
		AutoDetection:         isPremium,
		BankConnection:        true,
		BudgetManagement:      true,
		Analytics:             isPremium,
		DataExport:            isPremium,
		MultiAccount:          isPremium,
		Notifications:         true,
		DarkMode:              true,
		MultiLanguage:         true,
		TwoFactorAuth:         true,
		PaymentReminders:      isPremium,
		SpendingInsights:      isPremium,
		GoalSetting:           isPremium,
		RecurringTransactions: true,
		CategoryCustomization: isPremium,
		ReportsGeneration:     isPremium,
	}

	return features
}

// GetMainMenu returns the main navigation menu
func (s *ConfigService) GetMainMenu(features FeatureFlags, isPremium bool) []MenuItem {
	menu := []MenuItem{
		{
			ID:        "dashboard",
			Label:     "Dashboard",
			Icon:      "dashboard",
			Route:     "/dashboard",
			IsVisible: true,
			Order:     1,
		},
		{
			ID:        "accounts",
			Label:     "Accounts",
			Icon:      "account_balance",
			Route:     "/accounts",
			IsVisible: features.BankConnection,
			Order:     2,
		},
		{
			ID:        "transactions",
			Label:     "Transactions",
			Icon:      "receipt_long",
			Route:     "/transactions",
			IsVisible: true,
			Order:     3,
		},
		{
			ID:        "subscriptions",
			Label:     "Subscriptions",
			Icon:      "autorenew",
			Route:     "/subscriptions",
			IsVisible: features.SubscriptionTracking,
			Order:     4,
			Badge:     nil, // Could be populated with upcoming payment count
		},
		{
			ID:        "budgets",
			Label:     "Budgets",
			Icon:      "account_balance_wallet",
			Route:     "/budgets",
			IsVisible: features.BudgetManagement,
			IsPremium: !isPremium && features.BudgetManagement,
			Order:     5,
		},
		{
			ID:        "analytics",
			Label:     "Analytics",
			Icon:      "analytics",
			Route:     "/analytics",
			IsVisible: features.Analytics,
			IsPremium: !isPremium && features.Analytics,
			Order:     6,
		},
		{
			ID:        "goals",
			Label:     "Goals",
			Icon:      "flag",
			Route:     "/goals",
			IsVisible: features.GoalSetting,
			IsPremium: !isPremium && features.GoalSetting,
			Order:     7,
		},
		{
			ID:        "reports",
			Label:     "Reports",
			Icon:      "description",
			Route:     "/reports",
			IsVisible: features.ReportsGeneration,
			IsPremium: !isPremium && features.ReportsGeneration,
			Order:     8,
		},
	}

	// Filter out invisible items
	var visibleMenu []MenuItem
	for _, item := range menu {
		if item.IsVisible {
			visibleMenu = append(visibleMenu, item)
		}
	}

	return visibleMenu
}

// GetUserMenu returns the user profile menu
func (s *ConfigService) GetUserMenu(isPremium bool) []MenuItem {
	return []MenuItem{
		{
			ID:        "profile",
			Label:     "Profile",
			Icon:      "person",
			Route:     "/profile",
			IsVisible: true,
			Order:     1,
		},
		{
			ID:        "settings",
			Label:     "Settings",
			Icon:      "settings",
			Route:     "/settings",
			IsVisible: true,
			Order:     2,
		},
		{
			ID:        "subscription",
			Label:     "My Subscription",
			Icon:      "card_membership",
			Route:     "/subscription",
			IsVisible: true,
			Badge: func() *Badge {
				if !isPremium {
					return &Badge{
						Value: "Upgrade",
						Type:  "info",
					}
				}
				return nil
			}(),
			Order: 3,
		},
		{
			ID:        "help",
			Label:     "Help & Support",
			Icon:      "help",
			Route:     "/help",
			IsVisible: true,
			Order:     4,
		},
		{
			ID:        "logout",
			Label:     "Logout",
			Icon:      "logout",
			Route:     "/logout",
			IsVisible: true,
			Order:     5,
		},
	}
}

// GetQuickActions returns quick action buttons
func (s *ConfigService) GetQuickActions(features FeatureFlags) []MenuItem {
	return []MenuItem{
		{
			ID:        "add_transaction",
			Label:     "Add Transaction",
			Icon:      "add_circle",
			Route:     "/transactions/new",
			IsVisible: true,
			Order:     1,
		},
		{
			ID:        "add_subscription",
			Label:     "Add Subscription",
			Icon:      "add_card",
			Route:     "/subscriptions/new",
			IsVisible: features.SubscriptionTracking,
			Order:     2,
		},
		{
			ID:        "connect_bank",
			Label:     "Connect Bank",
			Icon:      "account_balance",
			Route:     "/accounts/connect",
			IsVisible: features.BankConnection,
			Order:     3,
		},
		{
			ID:        "export_data",
			Label:     "Export Data",
			Icon:      "download",
			Route:     "/export",
			IsVisible: features.DataExport,
			Order:     4,
		},
	}
}

// GetAppConfig returns the complete app configuration
func (s *ConfigService) GetAppConfig(userType string, isPremium bool, language string) AppConfig {
	features := s.GetFeatureFlags(userType, isPremium)
	
	config := AppConfig{
		Features: features,
		Menus: MenuConfig{
			MainMenu:     s.GetMainMenu(features, isPremium),
			UserMenu:     s.GetUserMenu(isPremium),
			SettingsMenu: s.GetSettingsMenu(features),
			QuickActions: s.GetQuickActions(features),
		},
		Settings:  s.GetAppSettings(),
		Themes:    s.GetThemes(),
		Languages: s.GetLanguages(),
		Categories: s.GetCategories(language),
		SubscriptionTypes: s.GetSubscriptionTypes(language),
		Currencies: s.GetCurrencies(),
	}

	return config
}

// GetSettingsMenu returns settings submenu
func (s *ConfigService) GetSettingsMenu(features FeatureFlags) []MenuItem {
	return []MenuItem{
		{
			ID:        "account",
			Label:     "Account",
			Icon:      "manage_accounts",
			Route:     "/settings/account",
			IsVisible: true,
			Order:     1,
		},
		{
			ID:        "security",
			Label:     "Security",
			Icon:      "security",
			Route:     "/settings/security",
			IsVisible: true,
			Order:     2,
		},
		{
			ID:        "notifications",
			Label:     "Notifications",
			Icon:      "notifications",
			Route:     "/settings/notifications",
			IsVisible: features.Notifications,
			Order:     3,
		},
		{
			ID:        "privacy",
			Label:     "Privacy",
			Icon:      "privacy_tip",
			Route:     "/settings/privacy",
			IsVisible: true,
			Order:     4,
		},
		{
			ID:        "appearance",
			Label:     "Appearance",
			Icon:      "palette",
			Route:     "/settings/appearance",
			IsVisible: true,
			Order:     5,
		},
		{
			ID:        "data",
			Label:     "Data Management",
			Icon:      "storage",
			Route:     "/settings/data",
			IsVisible: true,
			Order:     6,
		},
	}
}

// GetAppSettings returns global app settings
func (s *ConfigService) GetAppSettings() AppSettings {
	return AppSettings{
		AppName:            "SUMA Finance",
		AppVersion:         "1.0.0",
		APIVersion:         "v1",
		MaintenanceMode:    false,
		MaintenanceMessage: "",
		MinClientVersion:   "1.0.0",
		TermsURL:          "/terms",
		PrivacyURL:        "/privacy",
		SupportEmail:      "support@suma.pt",
		SupportURL:        "/support",
		DefaultCurrency:   "EUR",
		DefaultLanguage:   "pt",
		DefaultTheme:      "light",
		DateFormat:        "DD/MM/YYYY",
		TimeFormat:        "24h",
		WeekStartsOn:      1, // Monday
		SessionTimeout:    60, // minutes
		MaxLoginAttempts:  5,
		PasswordMinLength: 8,
		AllowedFileTypes:  []string{"pdf", "csv", "xlsx"},
		MaxFileSize:       10485760, // 10MB
	}
}

// GetThemes returns available themes
func (s *ConfigService) GetThemes() []Theme {
	return []Theme{
		{
			ID:     "light",
			Name:   "Light",
			IsDark: false,
			Colors: map[string]string{
				"primary":    "#4F46E5",
				"secondary":  "#7C3AED",
				"background": "#FFFFFF",
				"surface":    "#F9FAFB",
				"text":       "#111827",
			},
		},
		{
			ID:     "dark",
			Name:   "Dark",
			IsDark: true,
			Colors: map[string]string{
				"primary":    "#6366F1",
				"secondary":  "#8B5CF6",
				"background": "#111827",
				"surface":    "#1F2937",
				"text":       "#F9FAFB",
			},
		},
	}
}

// GetLanguages returns supported languages
func (s *ConfigService) GetLanguages() []Language {
	return []Language{
		{
			Code:       "pt",
			Name:       "Portuguese",
			NativeName: "Português",
			Direction:  "ltr",
			IsDefault:  true,
		},
		{
			Code:       "en",
			Name:       "English",
			NativeName: "English",
			Direction:  "ltr",
			IsDefault:  false,
		},
		{
			Code:       "es",
			Name:       "Spanish",
			NativeName: "Español",
			Direction:  "ltr",
			IsDefault:  false,
		},
	}
}

// GetCategories returns transaction categories
func (s *ConfigService) GetCategories(language string) []Category {
	// Categories would be localized based on language
	return []Category{
		{ID: "1", Name: "Food & Dining", Icon: "restaurant", Color: "#EF4444", Type: "expense"},
		{ID: "2", Name: "Transport", Icon: "directions_car", Color: "#F59E0B", Type: "expense"},
		{ID: "3", Name: "Shopping", Icon: "shopping_bag", Color: "#10B981", Type: "expense"},
		{ID: "4", Name: "Entertainment", Icon: "movie", Color: "#8B5CF6", Type: "expense"},
		{ID: "5", Name: "Bills & Utilities", Icon: "receipt", Color: "#6366F1", Type: "expense"},
		{ID: "6", Name: "Healthcare", Icon: "local_hospital", Color: "#EC4899", Type: "expense"},
		{ID: "7", Name: "Education", Icon: "school", Color: "#14B8A6", Type: "expense"},
		{ID: "8", Name: "Fitness", Icon: "fitness_center", Color: "#F97316", Type: "expense"},
		{ID: "9", Name: "Salary", Icon: "payments", Color: "#10B981", Type: "income"},
		{ID: "10", Name: "Investment", Icon: "trending_up", Color: "#6366F1", Type: "income"},
		{ID: "11", Name: "Other Income", Icon: "attach_money", Color: "#14B8A6", Type: "income"},
		{ID: "12", Name: "Other", Icon: "more_horiz", Color: "#6B7280", Type: "both"},
	}
}

// GetSubscriptionTypes returns billing cycle types
func (s *ConfigService) GetSubscriptionTypes(language string) []SubscriptionType {
	return []SubscriptionType{
		{ID: "weekly", Name: "Weekly", Days: 7},
		{ID: "biweekly", Name: "Bi-weekly", Days: 14},
		{ID: "monthly", Name: "Monthly", Days: 30},
		{ID: "quarterly", Name: "Quarterly", Days: 90},
		{ID: "semiannual", Name: "Semi-annual", Days: 180},
		{ID: "annual", Name: "Annual", Days: 365},
	}
}

// GetCurrencies returns supported currencies
func (s *ConfigService) GetCurrencies() []Currency {
	return []Currency{
		{Code: "EUR", Name: "Euro", Symbol: "€", Rate: 1.0},
		{Code: "USD", Name: "US Dollar", Symbol: "$", Rate: 1.08},
		{Code: "GBP", Name: "British Pound", Symbol: "£", Rate: 0.86},
		{Code: "BRL", Name: "Brazilian Real", Symbol: "R$", Rate: 5.40},
	}
}

// GetUserConfig returns user-specific configuration
func (s *ConfigService) GetUserConfig(userID string, isPremium bool) (map[string]interface{}, error) {
	// This could fetch user-specific settings from database
	config := map[string]interface{}{
		"user_id":     userID,
		"is_premium":  isPremium,
		"features":    s.GetFeatureFlags("user", isPremium),
		"preferences": map[string]interface{}{
			"theme":        "light",
			"language":     "pt",
			"currency":     "EUR",
			"date_format":  "DD/MM/YYYY",
			"time_format":  "24h",
			"week_starts":  1,
		},
		"limits": map[string]interface{}{
			"max_accounts":       func() int { if isPremium { return 10 } else { return 3 } }(),
			"max_subscriptions":  func() int { if isPremium { return -1 } else { return 20 } }(),
			"export_formats":     func() []string { if isPremium { return []string{"csv", "pdf", "xlsx"} } else { return []string{"csv"} } }(),
			"data_retention":     func() int { if isPremium { return 365 } else { return 90 } }(), // days
		},
		"subscription": map[string]interface{}{
			"plan":         func() string { if isPremium { return "premium" } else { return "free" } }(),
			"price":        func() float64 { if isPremium { return 4.99 } else { return 0 } }(),
			"renewal_date": func() *time.Time { if isPremium { t := time.Now().AddDate(0, 1, 0); return &t } else { return nil } }(),
		},
	}

	return config, nil
}