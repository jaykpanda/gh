package cmd

import (
	"gh/app"
	"strings"

	"github.com/spf13/cobra"
)

// GETCmd represents the GET command
var GETCmd = &cobra.Command{
	Use:   "GET",
	Short: "Perform a GET request",
	Run: func(cmd *cobra.Command, args []string) {
		performGet()
	},
}

func init() {
	rootCmd.AddCommand(GETCmd)

	GETCmd.Flags().StringVar(&urlString, "url", "", "URL endpoint to which the request is to be made")
	GETCmd.MarkFlagRequired("url")
	GETCmd.Flags().StringVar(&loginString, "login", "", "Login identifier to be used for the request")
	GETCmd.Flags().StringToStringVar(&headersMap, "headers", nil, "Headers to be used for the request")
	GETCmd.Flags().StringVar(&bodyString, "body", "", "Body for the request")
	GETCmd.Flags().StringVar(&outputFilter, "output-filter", "", "FIlter the outputs and presents the resulting entries in STDOUT.")
}

func performGet() {
	body := strings.NewReader(bodyString)
	requestParameters := app.RequestHandler{
		Method:  app.GET,
		URL:     urlString,
		Login:   loginString,
		Headers: headersMap,
		Body:    body,
	}
	requestParameters.PerformRequest()
}
