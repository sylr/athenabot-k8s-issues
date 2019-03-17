package main

type Sig struct {
	name          string
	strongMatches []string
	weakMatches   []string
}

var sigAws = Sig{
	name:          "aws",
	strongMatches: []string{"aws", "eks", "cloud-provider-aws", "aws-alb-ingress-controller", "aws-iam-authenticator", "aws-encryption-provider", "aws-ebs-csi-driver", "aws alb ingress controller", "aws iam authenticator", "aws encryption provider", "aws ebs csi driver"},
	weakMatches:   []string{"iam", "efs", "ebs", "alb ingress", "heptio authenticator"},
}

var sigCli = Sig{
	name:          "cli",
	strongMatches: []string{},
	weakMatches:   []string{"kubectl"},
}

var sigClusterLifeCycle = Sig{
	name:          "cluster-lifecycle",
	strongMatches: []string{"kubeadm"},
	weakMatches:   []string{},
}

var sigNetwork = Sig{
	name:          "network",
	strongMatches: []string{"ipv6", "ipvs", "ingress", "kube-dns", "kube dns", "kube-proxy", "kube proxy", "cni"},
	weakMatches:   []string{"envoy", "network", "service", "connection", "calico", "flannel", "istio", "linkerd"},
}

var sigNode = Sig{
	name:          "node",
	strongMatches: []string{},
	weakMatches:   []string{"node", "kubelet"},
}

var sigScheduling = Sig{
	name:          "scheduling",
	strongMatches: []string{"sheduler"},
	weakMatches:   []string{"schedule"},
}

var sigStorage = Sig{
	name:          "storage",
	strongMatches: []string{"persistentvolume"},
	weakMatches:   []string{"pv", "pvc", "efs", "ebs"},
}

var allSigs = []Sig{
	sigCli,
	sigClusterLifeCycle,
	sigNetwork,
	sigNode,
	sigScheduling,
	sigStorage,
	sigAws,
}
