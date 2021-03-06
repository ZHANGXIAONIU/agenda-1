// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


// create_meetingCmd represents the create_meeting command
var create_meetingCmd = &cobra.Command{
	Use:   "create_meeting -t [meeting title] -s [start time] -e [end time]",
	Short: "create a meeting",
	Long: `Input command model like : create_meeting -t Golang -s 2017-12-12/12:00 -e 2017-12-14/13:00

	in order to avoid longer command , we will call a fuction to input paticilator after you input command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		//this function need to be defined to input paticipator 
		tmp_t, _ := cmd.Flags().GetString("title")
		tmp_s, _ := cmd.Flags().GetString("start")
		tmp_e, _ := cmd.Flags().GetString("end")
		if service.GetFlag() == true {
			service.Create_meeting(tmp_t, tmp_s, tmp_e)
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(create_meetingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// create_meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// create_meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	create_meetingCmd.Flags().StringP("title", "t", "", "meeting title")
	create_meetingCmd.Flags().StringP("start",  "s", "", "start time")
	create_meetingCmd.Flags().StringP("end",  "e", "", "end time")

}
