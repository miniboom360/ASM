package activitys

import (
  "bytes"
  "context"
  "github.com/projectdiscovery/subfinder/v2/pkg/resolve"
  "github.com/projectdiscovery/subfinder/v2/pkg/runner"
  "io"
  "log"
)

// 用project discovery的项目来
// subfinder
func SearchSubDomain(ctx context.Context, domains []string) ([]string, error) {

  ss := make([]string, 0)

  runnerInstance, err := runner.NewRunner(&runner.Options{
    Threads:            10,                       // Thread controls the number of threads to use for active enumerations
    Timeout:            300,                      // Timeout is the seconds to wait for sources to respond
    MaxEnumerationTime: 10,                       // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
    Resolvers:          resolve.DefaultResolvers, // Use the default list of resolvers by marshaling it to the config
    ResultCallback: func(s *resolve.HostEntry) { // Callback function to execute for available host
      log.Println(s.Host, s.Source)
    },
  })

  buf := bytes.Buffer{}
  // runnerInstance.EnumerateMultipleDomains()

  for _, domain := range domains {
    err = runnerInstance.EnumerateSingleDomain(domain, []io.Writer{&buf})
    if err != nil {
      log.Fatal(err)
    }

    data, err := io.ReadAll(&buf)
    if err != nil {
      log.Fatal(err)
    }

    ss = append(ss, string(data))
    // fmt.Printf("%s", data)
  }

  return ss, nil
}
