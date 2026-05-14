package service

import (
	"context"
	"testing"
)

func TestRechargePlanCRUDListsOnlyForSalePlans(t *testing.T) {
	ctx := context.Background()
	svc := newRechargePlanTestService(t)

	visible, err := svc.CreateRechargePlan(ctx, CreateRechargePlanRequest{
		Name:          "Starter top-up",
		Description:   "External checkout",
		Price:         100,
		CreditAmount:  110,
		OriginalPrice: floatPtr(128),
		Features:      "Fast delivery\nSupport included",
		PurchaseURL:   "https://shop.example.com/buy/starter",
		Badge:         "Recommended",
		ForSale:       true,
		SortOrder:     2,
	})
	if err != nil {
		t.Fatalf("CreateRechargePlan visible returned error: %v", err)
	}
	_, err = svc.CreateRechargePlan(ctx, CreateRechargePlanRequest{
		Name:         "Hidden top-up",
		Price:        200,
		CreditAmount: 230,
		PurchaseURL:  "https://shop.example.com/buy/hidden",
		ForSale:      false,
		SortOrder:    1,
	})
	if err != nil {
		t.Fatalf("CreateRechargePlan hidden returned error: %v", err)
	}

	all, err := svc.ListRechargePlans(ctx)
	if err != nil {
		t.Fatalf("ListRechargePlans returned error: %v", err)
	}
	if len(all) != 2 {
		t.Fatalf("ListRechargePlans len = %d, want 2", len(all))
	}

	forSale, err := svc.ListRechargePlansForSale(ctx)
	if err != nil {
		t.Fatalf("ListRechargePlansForSale returned error: %v", err)
	}
	if len(forSale) != 1 || forSale[0].ID != visible.ID {
		t.Fatalf("ListRechargePlansForSale = %#v, want only visible plan %d", forSale, visible.ID)
	}
	if forSale[0].PurchaseURL != "https://shop.example.com/buy/starter" {
		t.Fatalf("PurchaseURL = %q", forSale[0].PurchaseURL)
	}
}

func TestCreateRechargePlanRejectsUnsafePurchaseURL(t *testing.T) {
	ctx := context.Background()
	svc := newRechargePlanTestService(t)

	_, err := svc.CreateRechargePlan(ctx, CreateRechargePlanRequest{
		Name:         "Unsafe",
		Price:        100,
		CreditAmount: 100,
		PurchaseURL:  "javascript:alert(1)",
		ForSale:      true,
	})
	if err == nil {
		t.Fatal("CreateRechargePlan returned nil error for unsafe URL")
	}
}

func newRechargePlanTestService(t *testing.T) *PaymentConfigService {
	t.Helper()
	client := newPaymentConfigServiceTestClient(t)
	if err := ensureRechargePlansTable(context.Background(), client); err != nil {
		t.Fatalf("create recharge plans table: %v", err)
	}
	return &PaymentConfigService{entClient: client}
}

func floatPtr(v float64) *float64 {
	return &v
}
