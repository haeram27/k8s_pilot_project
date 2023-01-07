package k8s

import (
	"bufio"
	"fmt"
	"os"
)

func getK8SAccount() (string, string) {
	/*
		export KUBERNETES_SERVICE_PORT_HTTPS=8443
		export KUBERNETES_SERVICE_HOST=192.168.49.2

		https://kubernetes.default.svc.cluster.local
	*/

	api_server := fmt.Sprintf("https://%s:%s",
		os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT_HTTPS"))
	api_token := ""

	rfile, _ := os.Open("/var/run/secrets/kubernetes.io/serviceaccount/token")
	scanner := bufio.NewScanner(rfile)
	for scanner.Scan() {
		api_token = scanner.Text()
	}
	defer rfile.Close()

	return api_server, api_token
}
