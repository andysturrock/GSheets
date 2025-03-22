package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()

	// Explicitly set the token source.
	ts, err := google.DefaultTokenSource(ctx, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to get token source: %v", err)
	}

	quotaProject, found := os.LookupEnv("GOOGLE_CLOUD_QUOTA_PROJECT")
	if !found {
		log.Fatalf("GOOGLE_CLOUD_QUOTA_PROJECT environment variable not set")
		return
	}

	srv, err := sheets.NewService(ctx, option.WithTokenSource(ts), option.WithQuotaProject(quotaProject))
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
		return
	}

	// --- Example: Reading data (replace with your actual Spreadsheet ID) ---
	spreadsheetId, found := os.LookupEnv("GOOGLE_SHEET_ID")
	if !found {
		log.Fatalf("GOOGLE_SHEET_ID environment variable not set")
		return
	}

	readRange, found := os.LookupEnv("GOOGLE_SHEET_RANGE")
	if !found {
		log.Fatalf("GOOGLE_SHEET_RANGE environment variable not set")
		return
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		// *** CRITICAL ERROR HANDLING ***
		if googleapiErr, ok := err.(*googleapi.Error); ok {
			if googleapiErr.Code == 403 && strings.Contains(googleapiErr.Message, "ACCESS_TOKEN_SCOPE_INSUFFICIENT") {
				log.Fatalf("ERROR: ACCESS_TOKEN_SCOPE_INSUFFICIENT. You need to update your scopes or ensure your ADC has the correct permissions! %v", err)
			} else {
				log.Fatalf("API Error: %v", err) // Handle other API errors
			}
		} else {
			log.Fatalf("Unable to retrieve data from sheet: %v", err) // Handle non-API errors
		}
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range resp.Values {
			// Handle potential index out of bounds errors.
			if len(row) > 0 {
				fmt.Printf("%s", row[0])
			}
			if len(row) > 1 {
				fmt.Printf(", %s", row[1])
			}
			fmt.Println()
		}
	}
}
