package spegel

func CLIFlags() []string {
	var flags []string
	// flags for local setup - TODO handle http-bootstrap-peer properly
	flags = append(flags,
		"registry",
		"--log-level=DEBUG",
		"--mirror-resolve-retries=3",
		"--mirror-resolve-timeout=20ms",
		"--registry-addr=:5500",
		"--router-addr=:5501",
		"--metrics-addr=:9590",
		"--registries",
		"https://docker.io",
		"https://ghcr.io",
		"https://quay.io",
		"https://registry.k8s.io",
		"https://k8s.gcr.io",
		"https://europe-docker.pkg.dev",
		"--containerd-sock=/run/containerd/containerd.sock",
		"--containerd-namespace=k8s.io",
		"--containerd-registry-config-path=/etc/containerd/certs.d",
		"--bootstrap-kind=http",
		"--http-bootstrap-addr=:5601",
		"--http-bootstrap-peer=http://10.10.130.192:5601/id",
		"--resolve-latest-tag=true",
		"--local-addr=$(NODE_IP):5500",
		"--containerd-content-path=/var/lib/containerd/io.containerd.content.v1.content",
	)

	return flags
}
