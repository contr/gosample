package main

import (
	"net"
	"runtime"
	"fmt"
	"os"
	
	"github.com/spf13/pflag"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"	
)

type ProxyServer struct {
	BindAddress        util.IP
	HealthzPort        int
	HealthzBindAddress util.IP
	OOMScoreAdj        int
	ResourceContainer  string
	Master             string
	Kubeconfig         string
}

// NewProxyServer creates a new ProxyServer object with default parameters
func NewProxyServer() *ProxyServer {
	return &ProxyServer{
		BindAddress:        util.IP(net.ParseIP("0.0.0.0")),
		HealthzPort:        10249,
		HealthzBindAddress: util.IP(net.ParseIP("127.0.0.1")),
		OOMScoreAdj:        -899,
		ResourceContainer:  "/kube-proxy",
	}
}

func (s *ProxyServer) AddFlags(fs *pflag.FlagSet) {
	fs.Var(&s.BindAddress, "bind-address", "The IP address for the proxy server to serve on (set to 0.0.0.0 for all interfaces)")
	fs.StringVar(&s.Master, "master", s.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.IntVar(&s.HealthzPort, "healthz-port", s.HealthzPort, "The port to bind the health check server. Use 0 to disable.")
	fs.Var(&s.HealthzBindAddress, "healthz-bind-address", "The IP address for the health check server to serve on, defaulting to 127.0.0.1 (set to 0.0.0.0 for all interfaces)")
	fs.IntVar(&s.OOMScoreAdj, "oom-score-adj", s.OOMScoreAdj, "The oom_score_adj value for kube-proxy process. Values must be within the range [-1000, 1000]")
	fs.StringVar(&s.ResourceContainer, "resource-container", s.ResourceContainer, "Absolute name of the resource-only container to create and run the Kube-proxy in (Default: /kube-proxy).")
	fs.StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
}

func (s *ProxyServer) Run(_ []string) error {
	// TODO(vmarmol): Use container config for this.


	return nil
}
	

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := NewProxyServer()
	s.AddFlags(pflag.CommandLine)
	pflag.Parse()

	fmt.Println("BindAddress: ", s.BindAddress)
	fmt.Println("Master: ", s.Master)
	fmt.Println("HealthzPorg: ", s.HealthzPort)

	if err := s.Run(pflag.CommandLine.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

