package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"strings"
	"net/http"
)

var NetworkCmd = &cobra.Command{
	Use:   "crypt",
	Short: "Realize networking operations",
	Long:  `Realize networking operations`,
	RunE: func(cmd *cobra.Command, args []string) error {

		//if Order == 1 {
		//	fmt.Println("Print: " + strings.Join(args, " "))
		//	return nil
		//} else {
		//	return errors.New("pete")
		//}
		return nil
	},
}

func post(url string, data string) {
	payload := strings.NewReader(data)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cookie", rand_string(30))
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "text/plain")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res)
}
