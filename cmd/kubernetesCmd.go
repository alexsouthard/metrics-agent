package cmd

import (
	"github.com/cloudability/metrics-agent/kubernetes"
	"github.com/cloudability/metrics-agent/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config       kubernetes.KubeAgentConfig
	requiredArgs = []string{
		"api_key",
		"cluster_name",
	}
	kubernetesCmd = &cobra.Command{
		Use:   "kubernetes",
		Short: "Collect Kubernetes Metrics",
		Long:  "Command to collect Kubernetes Metrics",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return util.CheckRequiredSettings(requiredArgs)
		},
		Run: func(cmd *cobra.Command, args []string) {
			kubernetes.CollectKubeMetrics(config)
		},
	}
)

func init() {

	// add cobra and viper ENVs and flags
	kubernetesCmd.PersistentFlags().StringVar(
		&config.APIKey,
		"api_key",
		"",
		"Cloudability API Key - required",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.ClusterName,
		"cluster_name",
		"",
		"Kubernetes Cluster Name - required this must be unique to every cluster.",
	)
	kubernetesCmd.PersistentFlags().IntVar(
		&config.PollInterval,
		"poll_interval",
		180,
		"Time, in seconds, to poll the services infrastructure.",
	)
	kubernetesCmd.PersistentFlags().UintVar(
		&config.CollectionRetryLimit,
		"collection_retry_limit",
		kubernetes.DefaultCollectionRetry,
		"Number of times agent should attempt to gather metrics from each source upon a failure",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.Cert,
		"certificate_file",
		"",
		"The path to a certificate file. - Optional",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.Key,
		"key_file",
		"",
		"The path to a key file. - Optional",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.OutboundProxy,
		"outbound_proxy",
		"",
		"Outbound HTTP/HTTPS proxy eg: http://x.x.x.x:8080. Must have a scheme prefix (http:// or https://) - Optional",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.OutboundProxyAuth,
		"outbound_proxy_auth",
		"",
		"Outbound proxy basic authentication credentials. Must defined in the form username:password - Optional",
	)
	kubernetesCmd.PersistentFlags().BoolVar(
		&config.OutboundProxyInsecure,
		"outbound_proxy_insecure",
		false,
		"When true, does not verify TLS certificates when using the outbound proxy. Default: False",
	)
	kubernetesCmd.PersistentFlags().BoolVar(
		&config.Insecure,
		"insecure",
		false,
		"When true, does not verify certificates when making TLS connections. Default: False",
	)
	kubernetesCmd.PersistentFlags().BoolVar(
		&config.ForceKubeProxy,
		"force_kube_proxy",
		false,
		"When true, disables direct node connection and forces proxy use.",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.Namespace,
		"namespace",
		"cloudability",
		"Kubernetes Namespace that the Agent is Running In",
	)
	kubernetesCmd.PersistentFlags().StringVar(
		&config.ScratchDir,
		"scratch_dir",
		"/tmp",
		"Directory metrics will be written to",
	)
	kubernetesCmd.PersistentFlags().IntVar(
		&config.InformerResyncInterval,
		"informer_resync_interval",
		24,
		"Time (in hours) between informer resync",
	)
	kubernetesCmd.PersistentFlags().IntVar(
		&config.ConcurrentPollers,
		"number_of_concurrent_node_pollers",
		100,
		"Number of concurrent goroutines created when polling node data. Default 100",
	)
	kubernetesCmd.PersistentFlags().BoolVar(
		&config.ParseMetricData,
		"parse_metric_data",
		false,
		"When true, core files will be parsed and non-relevant data will be removed prior to upload. Default: False",
	)

	//nolint gas
	_ = viper.BindPFlag("api_key", kubernetesCmd.PersistentFlags().Lookup("api_key"))
	_ = viper.BindPFlag("cluster_name", kubernetesCmd.PersistentFlags().Lookup("cluster_name"))
	_ = viper.BindPFlag("heapster_override_url", kubernetesCmd.PersistentFlags().Lookup("heapster_override_url"))
	_ = viper.BindPFlag("poll_interval", kubernetesCmd.PersistentFlags().Lookup("poll_interval"))
	_ = viper.BindPFlag("collection_retry_limit", kubernetesCmd.PersistentFlags().Lookup("collection_retry_limit"))
	_ = viper.BindPFlag("certificate_file", kubernetesCmd.PersistentFlags().Lookup("certificate_file"))
	_ = viper.BindPFlag("key_file", kubernetesCmd.PersistentFlags().Lookup("key_file"))
	_ = viper.BindPFlag("outbound_proxy", kubernetesCmd.PersistentFlags().Lookup("outbound_proxy"))
	_ = viper.BindPFlag("outbound_proxy_auth", kubernetesCmd.PersistentFlags().Lookup("outbound_proxy_auth"))
	_ = viper.BindPFlag("outbound_proxy_insecure", kubernetesCmd.PersistentFlags().Lookup("outbound_proxy_insecure"))
	_ = viper.BindPFlag("insecure", kubernetesCmd.PersistentFlags().Lookup("insecure"))
	_ = viper.BindPFlag("retrieve_node_summaries", kubernetesCmd.PersistentFlags().Lookup("retrieve_node_summaries"))
	_ = viper.BindPFlag("get_all_container_stats", kubernetesCmd.PersistentFlags().Lookup("get_all_container_stats"))
	_ = viper.BindPFlag("force_kube_proxy", kubernetesCmd.PersistentFlags().Lookup("force_kube_proxy"))
	_ = viper.BindPFlag("namespace", kubernetesCmd.PersistentFlags().Lookup("namespace"))
	_ = viper.BindPFlag("collect_heapster_export", kubernetesCmd.PersistentFlags().Lookup("collect_heapster_export"))
	_ = viper.BindPFlag("scratch_dir", kubernetesCmd.PersistentFlags().Lookup("scratch_dir"))
	_ = viper.BindPFlag("informer_resync_interval", kubernetesCmd.PersistentFlags().Lookup("informer_resync_interval"))
	_ = viper.BindPFlag("number_of_concurrent_node_pollers",
		kubernetesCmd.PersistentFlags().Lookup("number_of_concurrent_node_pollers"))
	_ = viper.BindPFlag("parse_metric_data", kubernetesCmd.PersistentFlags().Lookup("parse_metric_data"))

	viper.SetEnvPrefix("cloudability")
	viper.AutomaticEnv()

	RootCmd.AddCommand(kubernetesCmd)

	config = kubernetes.KubeAgentConfig{
		APIKey:                viper.GetString("api_key"),
		ClusterName:           viper.GetString("cluster_name"),
		PollInterval:          viper.GetInt("poll_interval"),
		CollectionRetryLimit:  viper.GetUint("collection_retry_limit"),
		OutboundProxy:         viper.GetString("outbound_proxy"),
		OutboundProxyAuth:     viper.GetString("outbound_proxy_auth"),
		OutboundProxyInsecure: viper.GetBool("outbound_proxy_insecure"),
		Insecure:              viper.GetBool("insecure"),
		Cert:                  viper.GetString("certificate_file"),
		Key:                   viper.GetString("key_file"),
		ConcurrentPollers:     viper.GetInt("number_of_concurrent_node_pollers"),
		ForceKubeProxy:        viper.GetBool("force_kube_proxy"),
		Namespace:             viper.GetString("namespace"),
		ScratchDir:            viper.GetString("scratch_dir"),
		InformerResyncInterval: viper.GetInt("informer_resync_interval"),
		ParseMetricData:        viper.GetBool("parse_metric_data"),
	}

}
