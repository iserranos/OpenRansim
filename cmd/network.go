//	This file is part of OpenRansim.
//
//	Foobar is free software: you can redistribute it and/or modify
//	it under the terms of the GNU General Public License as published by
//	the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//
//	Foobar is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"net/http"
	"strings"
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
