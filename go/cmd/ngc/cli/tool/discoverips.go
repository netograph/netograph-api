package tool

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const greynoiseIP = "http://api.greynoise.io:8888/v1/query/ip"

type GNRecord struct {
	Name string `json:"name"`
}

type GNIPResponse struct {
	Status  string     `json:"status"`
	Count   int        `json:"returned_count"`
	Records []GNRecord `json:"records"`
}

func fetchCertNames(
	ctx context.Context,
	dc dsetapi.DsetClient,
	query string,
) ([]string, error) {
	names := map[string]bool{}
	r, err := dc.CertValidNamesSearch(ctx, &dsetapi.CertValidNamesSearchRequest{
		Dataset: viper.GetString("dset"),
		Query:   query,
	})
	if err != nil {
		return nil, err
	}
	for {
		if v, err := r.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		} else {
			for _, i := range v.Cert.Dnsnames {
				if i[0] == '*' {
					i = i[2:]
				}
				names[i] = true
			}
		}
	}

	ret := make([]string, len(names))
	i := 0
	for v := range names {
		ret[i] = v
		i++
	}
	return ret, nil
}

func fetchIPs(
	ctx context.Context,
	dc dsetapi.DsetClient,
	domain string,
) ([]string, error) {
	r, err := dc.IPsForDomain(ctx, &dsetapi.IPsForDomainRequest{
		Dataset: viper.GetString("dset"),
		Query:   domain,
	})
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for {
		if v, err := r.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		} else {
			ret = append(ret, v.Ip)
		}
	}
	return ret, nil
}

func discoverIPsCommand() *cobra.Command {
	var greynoise *bool
	cmd := &cobra.Command{
		Use:     "discoverips domain",
		Aliases: []string{"discips"},
		Short:   "Uses certificate data to discover IPs owned by an entity",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}

			fmt.Println("Loading certificates...")
			names, err := fetchCertNames(ctx, c, args[0])
			if err != nil {
				return err
			}
			fmt.Printf("Discovered %d names...\n", len(names))

			allips := map[string]string{}
			ipseq := []string{}
			for _, name := range names {
				ips, err := fetchIPs(ctx, c, name)
				if err != nil {
					return err
				}
				fmt.Printf("Discovering IPs for %s...\n", name)
				for _, ip := range ips {
					if _, ok := allips[ip]; !ok {
						ipseq = append(ipseq, ip)
					}
					allips[ip] = name
				}
			}
			sort.Strings(ipseq)
			if *greynoise {
				fmt.Println("Checking IPs on greynoise.io...")
				for _, ip := range ipseq {
					resp, err := http.PostForm(
						greynoiseIP,
						url.Values{
							"ip": {ip},
						},
					)
					if err != nil {
						continue
					}
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						continue
					}
					fmt.Printf("%s (%s) greynoise: ", ip, allips[ip])
					fmt.Println(string(body))
				}
			} else {
				fmt.Printf("Discovered %d associated IPs...\n", len(allips))
				for _, k := range ipseq {
					fmt.Printf("%s (%s)\n", k, allips[k])
				}
			}
			return nil
		},
	}
	greynoise = cmd.Flags().BoolP(
		"greynoise", "g", false,
		"Run resulting IPs through Greynoise",
	)
	return cmd
}
