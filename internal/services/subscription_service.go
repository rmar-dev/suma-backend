package services

import (
	"math"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rmar-dev/suma-backend/internal/models"
	"github.com/rmar-dev/suma-backend/internal/repository"
)

type SubscriptionService struct {
	subscriptionRepo *repository.SubscriptionRepository
	transactionRepo  *repository.TransactionRepository
}

func NewSubscriptionService(subRepo *repository.SubscriptionRepository, transRepo *repository.TransactionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subRepo,
		transactionRepo:  transRepo,
	}
}

// DetectRecurringPayments analyzes transactions to find potential subscriptions
func (s *SubscriptionService) DetectRecurringPayments(userID string, accountID string) ([]*models.Subscription, error) {
	// Get last 90 days of transactions
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -90)
	
	transactions, err := s.transactionRepo.GetByDateRange(accountID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Group transactions by merchant and amount
	merchantGroups := s.groupTransactionsByMerchant(transactions)
	
	var detectedSubscriptions []*models.Subscription
	
	for merchant, txns := range merchantGroups {
		// Need at least 2 transactions to detect pattern
		if len(txns) < 2 {
			continue
		}
		
		// Check if amounts are consistent
		amounts := s.extractAmounts(txns)
		if !s.areAmountsConsistent(amounts) {
			continue
		}
		
		// Check if timing is regular
		dates := s.extractDates(txns)
		billingCycle, confidence := s.detectBillingCycle(dates)
		
		if confidence > 0.7 { // 70% confidence threshold
			subscription := &models.Subscription{
				ID:               uuid.New(),
				UserID:           uuid.MustParse(userID),
				AccountID:        uuid.MustParse(accountID),
				Name:             s.cleanMerchantName(merchant),
				MerchantName:     merchant,
				Amount:           amounts[0], // Use most recent amount
				Currency:         "EUR",
				BillingCycle:     models.BillingCycle(billingCycle),
				StartDate:        dates[0],
				NextBillingDate:  s.calculateNextBillingDate(dates[len(dates)-1], billingCycle),
				LastBillingDate:  &dates[len(dates)-1],
				Status:           "active",
				AutoDetected:     true,
				Confidence:       float32(confidence),
				NotifyDaysBefore: 3,
				NotifyEnabled:    true,
				Category:         s.categorizeSubscription(merchant),
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}
			
			detectedSubscriptions = append(detectedSubscriptions, subscription)
		}
	}
	
	return detectedSubscriptions, nil
}

// groupTransactionsByMerchant groups transactions by merchant name
func (s *SubscriptionService) groupTransactionsByMerchant(transactions []*models.Transaction) map[string][]*models.Transaction {
	groups := make(map[string][]*models.Transaction)
	
	for _, txn := range transactions {
		if txn.MerchantName != "" && txn.Amount < 0 { // Only consider debits
			normalizedName := s.normalizeMerchantName(txn.MerchantName)
			groups[normalizedName] = append(groups[normalizedName], txn)
		}
	}
	
	return groups
}

// normalizeMerchantName standardizes merchant names for grouping
func (s *SubscriptionService) normalizeMerchantName(name string) string {
	// Remove common suffixes and clean up
	name = strings.ToLower(name)
	name = strings.TrimSpace(name)
	
	// Remove common payment processor suffixes
	suffixes := []string{
		" recurring", " subscription", " monthly", " annual",
		" *", ".", ",", "-", "_",
	}
	
	for _, suffix := range suffixes {
		name = strings.TrimSuffix(name, suffix)
	}
	
	return name
}

// cleanMerchantName creates a display-friendly name
func (s *SubscriptionService) cleanMerchantName(name string) string {
	// Capitalize first letter of each word
	words := strings.Fields(name)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// extractAmounts gets all amounts from transactions
func (s *SubscriptionService) extractAmounts(transactions []*models.Transaction) []float64 {
	amounts := make([]float64, len(transactions))
	for i, txn := range transactions {
		amounts[i] = math.Abs(txn.Amount)
	}
	return amounts
}

// areAmountsConsistent checks if amounts are relatively consistent
func (s *SubscriptionService) areAmountsConsistent(amounts []float64) bool {
	if len(amounts) < 2 {
		return false
	}
	
	// Calculate average
	var sum float64
	for _, amt := range amounts {
		sum += amt
	}
	avg := sum / float64(len(amounts))
	
	// Check if all amounts are within 10% of average
	for _, amt := range amounts {
		deviation := math.Abs(amt-avg) / avg
		if deviation > 0.1 { // 10% tolerance
			return false
		}
	}
	
	return true
}

// extractDates gets all transaction dates sorted
func (s *SubscriptionService) extractDates(transactions []*models.Transaction) []time.Time {
	dates := make([]time.Time, len(transactions))
	for i, txn := range transactions {
		dates[i] = txn.TransactionDate
	}
	
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	
	return dates
}

// detectBillingCycle analyzes dates to find billing pattern
func (s *SubscriptionService) detectBillingCycle(dates []time.Time) (string, float64) {
	if len(dates) < 2 {
		return "", 0
	}
	
	// Calculate intervals between dates
	intervals := make([]int, len(dates)-1)
	for i := 1; i < len(dates); i++ {
		intervals[i-1] = int(dates[i].Sub(dates[i-1]).Hours() / 24) // Days between
	}
	
	// Calculate average interval
	var sum int
	for _, interval := range intervals {
		sum += interval
	}
	avgInterval := float64(sum) / float64(len(intervals))
	
	// Determine billing cycle based on average interval
	var billingCycle string
	var expectedInterval float64
	
	switch {
	case avgInterval >= 7-2 && avgInterval <= 7+2:
		billingCycle = "weekly"
		expectedInterval = 7
	case avgInterval >= 14-3 && avgInterval <= 14+3:
		billingCycle = "biweekly"
		expectedInterval = 14
	case avgInterval >= 30-5 && avgInterval <= 30+5:
		billingCycle = "monthly"
		expectedInterval = 30
	case avgInterval >= 90-10 && avgInterval <= 90+10:
		billingCycle = "quarterly"
		expectedInterval = 90
	case avgInterval >= 365-30 && avgInterval <= 365+30:
		billingCycle = "annual"
		expectedInterval = 365
	default:
		return "", 0
	}
	
	// Calculate confidence based on consistency
	var variance float64
	for _, interval := range intervals {
		variance += math.Pow(float64(interval)-expectedInterval, 2)
	}
	variance = variance / float64(len(intervals))
	stdDev := math.Sqrt(variance)
	
	// Confidence decreases with higher standard deviation
	confidence := math.Max(0, 1-stdDev/expectedInterval)
	
	return billingCycle, confidence
}

// calculateNextBillingDate calculates the next expected billing date
func (s *SubscriptionService) calculateNextBillingDate(lastDate time.Time, billingCycle string) time.Time {
	switch billingCycle {
	case "weekly":
		return lastDate.AddDate(0, 0, 7)
	case "biweekly":
		return lastDate.AddDate(0, 0, 14)
	case "monthly":
		return lastDate.AddDate(0, 1, 0)
	case "quarterly":
		return lastDate.AddDate(0, 3, 0)
	case "annual":
		return lastDate.AddDate(1, 0, 0)
	default:
		return lastDate.AddDate(0, 1, 0) // Default to monthly
	}
}

// categorizeSubscription attempts to categorize based on merchant name
func (s *SubscriptionService) categorizeSubscription(merchant string) string {
	merchant = strings.ToLower(merchant)
	
	categories := map[string][]string{
		"Entertainment": {"netflix", "spotify", "hbo", "disney", "prime video", "youtube", "twitch"},
		"Software":      {"adobe", "microsoft", "dropbox", "google", "slack", "zoom", "notion"},
		"Gaming":        {"xbox", "playstation", "steam", "epic", "nintendo"},
		"Fitness":       {"gym", "fitness", "strava", "peloton"},
		"News":          {"times", "post", "journal", "news"},
		"Food":          {"uber eats", "grubhub", "doordash", "hello fresh"},
		"Transportation": {"uber", "lyft", "lime", "bird"},
		"Utilities":     {"electric", "gas", "water", "internet", "phone", "mobile"},
	}
	
	for category, keywords := range categories {
		for _, keyword := range keywords {
			if strings.Contains(merchant, keyword) {
				return category
			}
		}
	}
	
	return "Other"
}

// CreateManualSubscription allows users to manually add subscriptions
func (s *SubscriptionService) CreateManualSubscription(subscription *models.Subscription) error {
	subscription.ID = uuid.New()
	subscription.AutoDetected = false
	subscription.Confidence = 1.0 // 100% confidence for manual entries
	subscription.Status = "active"
	subscription.CreatedAt = time.Now()
	subscription.UpdatedAt = time.Now()
	
	return s.subscriptionRepo.Create(subscription)
}

// GetUserSubscriptions retrieves all subscriptions for a user
func (s *SubscriptionService) GetUserSubscriptions(userID string) ([]*models.Subscription, error) {
	return s.subscriptionRepo.GetByUserID(userID)
}

// CancelSubscription marks a subscription as cancelled
func (s *SubscriptionService) CancelSubscription(subscriptionID string, reason string) error {
	subscription, err := s.subscriptionRepo.GetByID(subscriptionID)
	if err != nil {
		return err
	}
	
	subscription.Status = "cancelled"
	now := time.Now()
	subscription.CancelledAt = &now
	subscription.CancellationReason = reason
	subscription.UpdatedAt = now
	
	return s.subscriptionRepo.Update(subscription)
}

// GetUpcomingPayments gets subscriptions with upcoming billing dates
func (s *SubscriptionService) GetUpcomingPayments(userID string, days int) ([]*models.Subscription, error) {
	endDate := time.Now().AddDate(0, 0, days)
	return s.subscriptionRepo.GetUpcomingPayments(userID, endDate)
}

// CalculateTotalMonthlySpend calculates total monthly subscription spend
func (s *SubscriptionService) CalculateTotalMonthlySpend(userID string) (float64, error) {
	subscriptions, err := s.subscriptionRepo.GetActiveByUserID(userID)
	if err != nil {
		return 0, err
	}
	
	var total float64
	for _, sub := range subscriptions {
		monthlyAmount := s.convertToMonthly(sub.Amount, string(sub.BillingCycle))
		total += monthlyAmount
	}
	
	return total, nil
}

// convertToMonthly converts subscription amount to monthly equivalent
func (s *SubscriptionService) convertToMonthly(amount float64, billingCycle string) float64 {
	switch billingCycle {
	case "weekly":
		return amount * 4.33 // Average weeks per month
	case "biweekly":
		return amount * 2.17
	case "monthly":
		return amount
	case "quarterly":
		return amount / 3
	case "annual":
		return amount / 12
	default:
		return amount
	}
}
