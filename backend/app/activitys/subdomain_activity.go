package activitys

import (
	"backend/app"
	"bytes"
	"context"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
	"log"
	"strings"
)

// 用project discovery的项目来
// subfinder
func SearchSubDomain(ctx context.Context, sti app.ScanTaskItem) ([]*app.SubdomainS, error) {
	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            10,                       // Thread controls the number of threads to use for active enumerations
		Timeout:            300,                      // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10,                       // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Resolvers:          resolve.DefaultResolvers, // Use the default list of resolvers by marshaling it to the config
		ResultCallback: func(s *resolve.HostEntry) { // Callback function to execute for available host
			log.Println(s.Host, s.Source)
		},
	})
	ssr := make([]*app.SubdomainS, 0)
	for _, domain := range sti.Domains {
		s := new(app.SubdomainS)
		buf := bytes.Buffer{}
		err = runnerInstance.EnumerateSingleDomain(domain, []io.Writer{&buf})
		if err != nil {
			log.Fatal(err)
		}

		data, err := io.ReadAll(&buf)
		if err != nil {
			log.Fatal(err)
		}
		r := string(data)

		rs := strings.SplitN(r, "\n", -1)
		for _, v := range rs {
			if v != "" {
				s.SubdomainsSclice = append(s.SubdomainsSclice, v)
			}
		}
		if len(rs) != 0 {
			s.MainDomain = domain
			s.OrgName = sti.OrgName
			s.TaskId = sti.TaskId
			ssr = append(ssr, s)
		}
	}

	return ssr, nil
}
