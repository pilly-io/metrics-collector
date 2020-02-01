generate-mocks:
	mockgen -package mocks -destination internal/runner/mocks/executable_mock.go github.com/pilly-io/metrics-collector/internal/runner Executable
	mockgen -package mocks -destination internal/prometheus/client/mocks/prometheus_v1_mock.go github.com/prometheus/client_golang/api/prometheus/v1 API
	mockgen -package mocks -destination internal/prometheus/collector/mocks/client_mock.go github.com/pilly-io/metrics-collector/internal/prometheus/client Client
	mockgen -package mocks -mock_names IClient=MockKubernetesClient -destination internal/prometheus/collector/mocks/kubernetes_client_mock.go github.com/pilly-io/metrics-collector/internal/kubernetes IClient
	mockgen -package mocks -destination internal/kubernetes/mocks/mock_configurator.go github.com/pilly-io/metrics-collector/internal/kubernetes Configurator
	mockgen -package mocks -destination internal/cleaner/mocks/mock_database.go github.com/pilly-io/metrics-collector/internal/db Database