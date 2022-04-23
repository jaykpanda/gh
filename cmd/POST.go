package cmd

import (
	"gh/app"
	"strings"

	"github.com/spf13/cobra"
)

// POSTCmd represents the POST command
var POSTCmd = &cobra.Command{
	Use:   "POST",
	Short: "Perform a HTTP Post request",
	Run: func(cmd *cobra.Command, args []string) {
		performPost()
	},
}

func init() {
	rootCmd.AddCommand(POSTCmd)

	POSTCmd.Flags().StringVar(&urlString, "url", "", "URL endpoint to which the request is to be made")
	POSTCmd.MarkFlagRequired("url")
	POSTCmd.Flags().StringVar(&loginString, "login", "", "Login identifier to be used for the request")
	POSTCmd.Flags().StringToStringVar(&headersMap, "headers", nil, "Headers to be used for the request")
	POSTCmd.Flags().StringVar(&bodyString, "body", "", "Body for the request")
	POSTCmd.Flags().StringVar(&outputFilter, "output-filter", "", "FIlter the outputs and presents the resulting entries in STDOUT.")
}

func performPost() {
	body := strings.NewReader(bodyString)
	requestParameters := app.RequestHandler{
		Method:  app.POST,
		URL:     urlString,
		Login:   loginString,
		Headers: headersMap,
		Body:    body,
	}
	requestParameters.PerformRequest()
}
